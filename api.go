package goib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	l5g "github.com/neocortical/log5go"
)

const deliveryURL = "http://ibsys-api.ib-prod.com/v2.0/delivery/{channel}/json/{service}"

var log = l5g.Logger(l5g.LogAll)

type httpClient func(url string) ([]byte, error)

// API is the entrance point for interacting with the IB API
type API interface {
	Entry(channel string, entrytype string, params url.Values) (Item, error)
	Search(channel string, query string, params url.Values) (*Collection, error)
	Content(channel string, contentID int, params url.Values) (Item, error)
	ContentMedia(channel string, contentID int, params url.Values) ([]Item, error)
	ContentItems(channel string, contentID int, params url.Values) ([]Item, error)
	Closings(channel string, filter ClosingsFilter, providerID ...string) (ClosingsResponse, error)
	UnmarshalReceiver(r Receiver) (Item, error)
}

// NewAPI constructs an API object for the given channel
func NewAPI() API {
	return &api{
		deliveryURL: deliveryURL,
	}
}

func NewAPIWithDeliveryURL(url string) API {
	return &api{
		deliveryURL: url,
	}
}

type api struct {
	deliveryURL string
}

func (api *api) Entry(channel string, entrytype string, params url.Values) (entry Item, err error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "entry", 1)
	uri += "/" + entrytype

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := doGet(uri)
	if err != nil {
		return entry, err
	}

	return api.unmarshalResponse(bytes)
}

func (api *api) Search(channel string, query string, params url.Values) (s *Collection, err error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "search", 1)

	if params == nil {
		params = url.Values{}
	}
	params.Set("q", query)
	uri += "?" + params.Encode()

	bytes, err := doGet(uri)
	if err != nil {
		return s, err
	}

	r, err := api.unmarshalResponse(bytes)
	if err != nil {
		return s, err
	}

	if r.GetType() == CollectionType {
		return r.(*Collection), nil
	}

	return s, fmt.Errorf("invalid object type returned by search: %s", r.GetType())
}

func (api *api) Content(channel string, contentID int, params url.Values) (Item, error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID)

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := doGet(uri)
	if err != nil {
		return nil, err
	}

	return api.unmarshalResponse(bytes)
}

func (api *api) ContentMedia(channel string, contentID int, params url.Values) ([]Item, error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID) + "/media"

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := doGet(uri)
	if err != nil {
		return nil, err
	}

	return api.unmarshalArrayResponse(bytes)
}

func (api *api) ContentItems(channel string, contentID int, params url.Values) ([]Item, error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID) + "/items"

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := doGet(uri)
	if err != nil {
		return nil, err
	}

	return api.unmarshalArrayResponse(bytes)
}

func (api *api) Closings(channel string, filter ClosingsFilter, providerID ...string) (ClosingsResponse, error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "closings", 1)
	uri += "/" + string(filter)
	if filter == ClosingsInst && len(providerID) > 0 {
		uri += "/id/" + providerID[0]
	}

	bytes, err := doGet(uri)
	if err != nil {
		return ClosingsResponse{}, err
	}

	return unmarshalClosingsResponse(bytes)
}

func doGet(url string) (result []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Debug("got error response for URL %s: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	log.Debug("IB API: %s : HTTP Status: %s, Success: %t", url, resp.Status, (err == nil))
	return result, err
}

func (api *api) unmarshalResponse(bytes []byte) (Item, error) {
	var r Receiver

	err := json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	response, err := api.UnmarshalReceiver(r)

	return response, err
}

func (api *api) unmarshalArrayResponse(bytes []byte) (result []Item, err error) {
	var ra []Receiver

	err = json.Unmarshal(bytes, &ra)
	if err != nil {
		return nil, err
	}

	for _, r := range ra {
		item, err := api.UnmarshalReceiver(r)
		if err != nil {
			log.Warn("error unmarshalling item from array: %v", err)
		} else {
			result = append(result, item)
		}
	}

	return result, err
}

func unmarshalClosingsResponse(bytes []byte) (result ClosingsResponse, err error) {
	err = json.Unmarshal(bytes, &result)
	return result, err
}

func (api *api) unmarshalClsInstitution(bytes []byte) (result ClsInstitution, err error) {
	err = json.Unmarshal(bytes, &result)
	return result, err
}

func (api *api) UnmarshalReceiver(r Receiver) (Item, error) {
	switch r.Type {
	case ArticleType:
		return api.unmarshalArticle(r), nil
	case VideoType:
		return api.unmarshalVideo(r), nil
	case CollectionType:
		return api.unmarshalCollection(r), nil
	case SearchType:
		return api.unmarshalSearch(r), nil
	case ImageType:
		return unmarshalImage(r), nil
	case GalleryType:
		return api.unmarshalGallery(r), nil
	case MapType:
		return unmarshalMap(r), nil
	case ExternalContentType:
		return unmarshalExternalContent(r), nil
	case ExternalLinkType:
		return api.unmarshalExternalLink(r), nil
	case HTMLType:
		return unmarshalHTMLContent(r), nil
	case PersonType:
		return unmarshalPerson(r), nil
	case LivevideoType:
		return api.unmarshalLivevideo(r), nil
	case SettingsType:
		return unmarshalSettings(r), nil
	case AudioType:
		return api.unmarshalAudio(r), nil
	case TeaserType:
		return api.unmarshalTeaser(r)
	default:
		return nil, fmt.Errorf("unknonwn response type for obj %d: %s", r.ContentID, r.Type)
	}
}

func (api *api) unmarshalArticle(r Receiver) (a *Article) {
	a = &Article{}
	a.Type = r.Type
	a.ContentID = r.ContentID
	a.TeaserTitle = getTeaserTitle(&r)
	a.TeaserText = r.TeaserText
	a.TeaserImage = r.TeaserImage
	a.PublicationDate = r.PublicationDate
	a.Title = r.Title
	a.Subheadline = r.Subheadline
	a.Text = r.Text
	a.Authors = r.Authors
	a.CanonicalURL = r.CanonicalURL
	a.URL = r.URL
	a.NavContext = r.NavContext
	a.AnalyticsCategory = r.AnalyticsCategory
	a.AdvertisingCategory = r.AdvertisingCategory
	a.Dateline = r.Dateline
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			a.Media = append(a.Media, item)
		}
	}

	for _, rInner := range r.RelatedMedia {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling related media sub-object: %v", err)
		} else {
			a.RelatedMedia = append(a.RelatedMedia, item)
		}
	}

	return a
}

func (api *api) unmarshalVideo(r Receiver) (v *Video) {
	v = &Video{}
	v.Type = r.Type
	v.ContentID = r.ContentID
	v.TeaserTitle = getTeaserTitle(&r)
	v.TeaserText = r.TeaserText
	v.Title = r.Title
	v.Subheadline = r.Subheadline
	v.TeaserImage = r.TeaserImage
	v.PublicationDate = r.PublicationDate
	v.Authors = r.Authors
	v.Flavors = r.Flavors
	v.CanonicalURL = r.CanonicalURL
	v.URL = r.URL
	v.NavContext = r.NavContext
	v.AnalyticsCategory = r.AnalyticsCategory
	v.AdvertisingCategory = r.AdvertisingCategory
	v.ShowAds = r.ShowAds
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			v.Media = append(v.Media, item)
		}
	}

	return v
}

func (api *api) unmarshalLivevideo(r Receiver) (l *Livevideo) {
	l = &Livevideo{}
	l.Type = r.Type
	l.ContentID = r.ContentID
	l.TeaserTitle = getTeaserTitle(&r)
	l.TeaserText = r.TeaserText
	l.Title = r.Title
	l.Subheadline = r.Subheadline
	l.PublicationDate = r.PublicationDate
	l.Authors = r.Authors
	l.Title = r.Title
	l.CanonicalURL = r.CanonicalURL
	l.URL = r.URL
	l.Stream = r.Stream
	l.ExternalID = r.ExternalID
	l.NavContext = r.NavContext
	l.AnalyticsCategory = r.AnalyticsCategory
	l.AdvertisingCategory = r.AdvertisingCategory
	l.ShowAds = r.ShowAds
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			l.Media = append(l.Media, item)
		}
	}

	return l
}

func unmarshalImage(r Receiver) (i *Image) {
	i = &Image{}
	i.Type = r.Type
	i.ContentID = r.ContentID
	i.Title = r.Title
	i.Subheadline = r.Subheadline
	i.TeaserTitle = getTeaserTitle(&r)
	i.TeaserText = r.TeaserText
	i.TeaserImage = r.TeaserImage
	i.PublicationDate = r.PublicationDate
	i.Authors = r.Authors
	i.Caption = r.Caption
	i.Author = r.Author
	i.URLs = r.URLs
	i.Copyright = r.Copyright
	i.CopyrightObjects = r.CopyrightObjects
	i.CanonicalURL = r.CanonicalURL
	i.NavContext = r.NavContext
	i.AnalyticsCategory = r.AnalyticsCategory
	i.AdvertisingCategory = r.AdvertisingCategory

	return i
}

func (api *api) unmarshalGallery(r Receiver) (g *Gallery) {
	g = &Gallery{}
	g.Type = r.Type
	g.ContentID = r.ContentID
	g.TeaserTitle = getTeaserTitle(&r)
	g.TeaserText = r.TeaserText
	g.TeaserImage = r.TeaserImage
	g.PublicationDate = r.PublicationDate
	g.Authors = r.Authors
	g.Title = r.Title
	g.Subheadline = r.Subheadline
	g.CanonicalURL = r.CanonicalURL
	g.URL = r.URL
	g.NavContext = r.NavContext
	g.AnalyticsCategory = r.AnalyticsCategory
	g.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			g.Media = append(g.Media, item)
		}
	}
	for _, rInner := range r.Items {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			g.Items = append(g.Items, item)
		}
	}

	g.Captions = unmarshalGalleryCaptions(r)

	return g
}

func unmarshalMap(r Receiver) (m *Map) {
	m = &Map{}
	m.Type = r.Type
	m.ContentID = r.ContentID
	m.PublicationDate = r.PublicationDate
	m.Title = r.Title
	m.Subheadline = r.Subheadline
	m.TeaserTitle = getTeaserTitle(&r)
	m.TeaserText = r.TeaserText
	m.StaticMap = r.StaticMap
	m.InteractiveMap = r.InteractiveMap
	m.CanonicalURL = r.CanonicalURL
	m.URL = r.URL
	m.NavContext = r.NavContext
	m.AnalyticsCategory = r.AnalyticsCategory
	m.AdvertisingCategory = r.AdvertisingCategory

	return m
}

func (api *api) unmarshalCollection(r Receiver) (c *Collection) {
	c = &Collection{}
	c.Type = r.Type
	c.ContentID = r.ContentID
	c.TeaserTitle = getTeaserTitle(&r)
	c.CollectionName = r.CollectionName
	c.ContentName = r.ContentName
	c.TotalCount = r.TotalCount
	c.StartIndex = r.StartIndex
	for _, rInner := range r.Items {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			c.Items = append(c.Items, item)
		}
	}
	c.Settings = r.Settings

	c.NavContext = r.NavContext
	c.AnalyticsCategory = r.AnalyticsCategory
	c.AdvertisingCategory = r.AdvertisingCategory

	return c
}

func (api *api) unmarshalSearch(r Receiver) (s *Collection) {
	s = &Collection{}
	s.Type = s.GetType()
	s.Keywords = r.Keywords
	s.TotalCount = r.TotalCount
	s.StartIndex = r.StartIndex
	for _, rInner := range r.Items {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			s.Items = append(s.Items, item)
		}
	}

	return s
}

func unmarshalExternalContent(r Receiver) (e *ExternalContent) {
	e = &ExternalContent{}
	e.Type = r.Type
	e.ContentID = r.ContentID
	e.PublicationDate = r.PublicationDate
	e.TeaserTitle = getTeaserTitle(&r)
	e.ExternalContent = r.ExternalContent
	e.Struct = r.Struct
	return e
}

func (api *api) unmarshalExternalLink(r Receiver) (e *ExternalLink) {
	e = &ExternalLink{}
	e.Type = r.Type
	e.ContentID = r.ContentID
	e.PublicationDate = r.PublicationDate
	e.TeaserTitle = getTeaserTitle(&r)
	e.TeaserText = r.TeaserText
	e.CanonicalURL = r.CanonicalURL
	e.URL = r.URL
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			e.Media = append(e.Media, item)
		}
	}

	return e
}

func unmarshalHTMLContent(r Receiver) (h *HTMLContent) {
	h = &HTMLContent{}
	h.Type = r.Type
	h.ContentID = r.ContentID
	h.PublicationDate = r.PublicationDate
	h.Code = r.Code
	h.NavContext = r.NavContext
	h.AnalyticsCategory = r.AnalyticsCategory
	h.AdvertisingCategory = r.AdvertisingCategory
	return h
}

func getTeaserTitle(r *Receiver) (title string) {
	title = r.TeaserTitle
	if title == "" {
		title = r.Title
	}
	return title
}

func unmarshalSettings(r Receiver) (s *Settings) {
	s = &Settings{}
	s.ContentID = r.ContentID
	if len(r.Settings) > 0 {
		s.Settings = r.Settings[0]
	}

	return s
}

func unmarshalPerson(r Receiver) (p *Person) {
	p = &Person{}
	p.Type = r.Type
	p.ContentID = r.ContentID
	p.FullName = r.FullName
	p.Title = r.Title
	p.Blurb = r.TeaserText
	p.TeaserImage = r.TeaserImage
	p.PublicationDate = r.PublicationDate
	p.Bio = r.Bio
	p.Email = r.Email
	p.CanonicalURL = r.CanonicalURL
	p.URL = r.URL
	p.Photo = r.Photo
	p.NavContext = r.NavContext
	p.AnalyticsCategory = r.AnalyticsCategory
	p.AdvertisingCategory = r.AdvertisingCategory

	return p
}

func (api *api) unmarshalAudio(r Receiver) (a *Audio) {
	a = &Audio{}
	a.Type = r.Type
	a.ContentID = r.ContentID
	a.Title = r.Title
	a.Subheadline = r.Subheadline
	a.TeaserTitle = getTeaserTitle(&r)
	a.TeaserText = r.TeaserText
	a.PublicationDate = r.PublicationDate
	a.Authors = r.Authors
	a.CanonicalURL = r.CanonicalURL
	a.URL = r.URL
	a.Stream = r.Stream
	a.NavContext = r.NavContext
	a.AnalyticsCategory = r.AnalyticsCategory
	a.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			a.Media = append(a.Media, item)
		}
	}

	return a
}

func (api *api) unmarshalTeaser(r Receiver) (t *Teaser, err error) {
	t = &Teaser{}
	t.Type = r.Type
	t.ContentID = r.ContentID
	t.Title = r.Title
	t.TeaserTitle = getTeaserTitle(&r)
	t.TeaserText = r.TeaserText
	t.PublicationDate = r.PublicationDate
	t.Authors = r.Authors
	t.NavContext = r.NavContext
	t.AnalyticsCategory = r.AnalyticsCategory
	t.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := api.UnmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			t.Media = append(t.Media, item)
		}
	}

	if r.Target == nil {
		return t, fmt.Errorf("teaser object missing target %d", t.ContentID)
	}

	target, err := api.UnmarshalReceiver(*r.Target)
	if err != nil {
		return t, err
	}
	t.Target = target

	return t, nil
}

func unmarshalGalleryCaptions(r Receiver) map[int]string {
	result := make(map[int]string)

	strct := r.Struct
	if strct == nil || len(strct) < 1 {
		return result
	}
	dataMap, ok := strct[0].(map[string]interface{})
	if !ok {
		return result
	}
	items, ok := dataMap["items"]
	if !ok {
		return result
	}
	captionMap, ok := items.(map[string]interface{})
	if !ok {
		return result
	}
	for k, v := range captionMap {
		coid, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		capmap, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		caption, ok := capmap["caption"].(string)
		if !ok {
			continue
		}
		result[coid] = caption
	}

	return result
}

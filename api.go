package goib

import (
	"encoding/json"
	"errors"
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
	Closings(channel string, filter ClosingsFilter) (ClosingsResponse, error)
}

// NewAPI constructs an API object for the given channel
func NewAPI() API {
	return &api{
		deliveryURL: deliveryURL,
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

	return unmarshalResponse(bytes)
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

	r, err := unmarshalResponse(bytes)
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

	return unmarshalResponse(bytes)
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

	return unmarshalArrayResponse(bytes)
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

	return unmarshalArrayResponse(bytes)
}

func (api *api) Closings(channel string, filter ClosingsFilter) (ClosingsResponse, error) {
	if filter == ClosingsInst {
		return ClosingsResponse{}, errors.New("institution filter not yet supported")
	}

	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "closings", 1)
	uri += "/" + string(filter)
	if filter == ClosingsAll || filter == ClosingsClosed {
		uri += "/ungrouped"
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
		return nil, err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	log.Debug("IB API: %s : HTTP Status: %s, Success: %t", url, resp.Status, (err == nil))
	return result, err
}

func unmarshalResponse(bytes []byte) (Item, error) {
	var r Receiver

	err := json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	response, err := unmarshalReceiver(r)

	return response, err
}

func unmarshalArrayResponse(bytes []byte) (result []Item, err error) {
	var ra []Receiver

	err = json.Unmarshal(bytes, &ra)
	if err != nil {
		return nil, err
	}

	for _, r := range ra {
		item, err := unmarshalReceiver(r)
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

func unmarshalReceiver(r Receiver) (Item, error) {
	switch r.Type {
	case ArticleType:
		return unmarshalArticle(r), nil
	case VideoType:
		return unmarshalVideo(r), nil
	case CollectionType:
		return unmarshalCollection(r), nil
	case SearchType:
		return unmarshalSearch(r), nil
	case ImageType:
		return unmarshalImage(r), nil
	case GalleryType:
		return unmarshalGallery(r), nil
	case MapType:
		return unmarshalMap(r), nil
	case ExternalType:
		return unmarshalExternalContent(r), nil
	case HTMLType:
		return unmarshalHTMLContent(r), nil
	case PersonType:
		return unmarshalPerson(r), nil
	case LivevideoType:
		return unmarshalLivevideo(r), nil
	case SettingsType:
		return unmarshalSettings(r), nil
	case AudioType:
		return unmarshalAudio(r), nil
	default:
		return nil, fmt.Errorf("unknonwn response type: %s", r.Type)
	}
}

func unmarshalArticle(r Receiver) (a *Article) {
	a = &Article{}
	a.ContentID = r.ContentID
	a.TeaserTitle = getTeaserTitle(&r)
	a.TeaserText = r.TeaserText
	a.TeaserImage = r.TeaserImage
	a.PublicationDate = r.PublicationDate
	a.Title = r.Title
	a.Subheadline = r.Subheadline
	a.Text = r.Text
	a.Authors = r.AuthorObjects
	a.CanonicalURL = r.CanonicalURL
	a.URL = r.URL
	a.NavContext = r.NavContext
	a.AnalyticsCategory = r.AnalyticsCategory
	a.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			a.Media = append(a.Media, item)
		}
	}

	return a
}

func unmarshalVideo(r Receiver) (v *Video) {
	v = &Video{}
	v.ContentID = r.ContentID
	v.TeaserTitle = getTeaserTitle(&r)
	v.TeaserText = r.TeaserText
	v.TeaserImage = r.TeaserImage
	v.PublicationDate = r.PublicationDate
	v.Authors = r.AuthorObjects
	v.Title = r.Title
	v.Subheadline = r.Subheadline
	v.Flavors = r.Flavors
	v.CanonicalURL = r.CanonicalURL
	v.URL = r.URL
	v.NavContext = r.NavContext
	v.AnalyticsCategory = r.AnalyticsCategory
	v.AdvertisingCategory = r.AdvertisingCategory

	return v
}

func unmarshalLivevideo(r Receiver) (l *Livevideo) {
	l = &Livevideo{}
	l.ContentID = r.ContentID
	l.TeaserTitle = getTeaserTitle(&r)
	l.TeaserText = r.TeaserText
	l.Title = r.Title
	l.Subheadline = r.Subheadline
	l.PublicationDate = r.PublicationDate
	l.Authors = r.AuthorObjects
	l.Title = r.Title
	l.CanonicalURL = r.CanonicalURL
	l.URL = r.URL
	l.Stream = r.Stream
	l.NavContext = r.NavContext
	l.AnalyticsCategory = r.AnalyticsCategory
	l.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := unmarshalReceiver(rInner)
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
	i.ContentID = r.ContentID
	i.TeaserTitle = getTeaserTitle(&r)
	i.TeaserText = r.TeaserText
	i.TeaserImage = r.TeaserImage
	i.PublicationDate = r.PublicationDate
	i.Authors = r.AuthorObjects
	i.Title = r.Title
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

func unmarshalGallery(r Receiver) (g *Gallery) {
	g = &Gallery{}
	g.ContentID = r.ContentID
	g.TeaserTitle = getTeaserTitle(&r)
	g.TeaserText = r.TeaserText
	g.TeaserImage = r.TeaserImage
	g.PublicationDate = r.PublicationDate
	g.Authors = r.AuthorObjects
	g.Title = r.Title
	g.Subheadline = r.Subheadline
	g.CanonicalURL = r.CanonicalURL
	g.URL = r.URL
	g.NavContext = r.NavContext
	g.AnalyticsCategory = r.AnalyticsCategory
	g.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			g.Media = append(g.Media, item)
		}
	}
	for _, rInner := range r.Items {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			g.Items = append(g.Items, item)
		}
	}

	return g
}

func unmarshalMap(r Receiver) (m *Map) {
	m = &Map{}
	m.ContentID = r.ContentID
	m.TeaserTitle = getTeaserTitle(&r)
	m.StaticMap = r.StaticMap
	m.InteractiveMap = r.InteractiveMap
	m.NavContext = r.NavContext
	m.AnalyticsCategory = r.AnalyticsCategory
	m.AdvertisingCategory = r.AdvertisingCategory

	return m
}

func unmarshalCollection(r Receiver) (c *Collection) {
	c = &Collection{}
	c.ContentID = r.ContentID
	c.TeaserTitle = getTeaserTitle(&r)
	c.CollectionName = r.CollectionName
	c.ContentName = r.ContentName
	c.TotalCount = r.TotalCount
	c.StartIndex = r.StartIndex
	for _, rInner := range r.Items {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			c.Items = append(c.Items, item)
		}
	}
	if len(r.Settings) > 0 {
		c.Settings = r.Settings[0]
	}

	c.NavContext = r.NavContext
	c.AnalyticsCategory = r.AnalyticsCategory
	c.AdvertisingCategory = r.AdvertisingCategory

	return c
}

func unmarshalSearch(r Receiver) (s *Collection) {
	s = &Collection{}
	s.Keywords = r.Keywords
	s.TotalCount = r.TotalCount
	s.StartIndex = r.StartIndex
	for _, rInner := range r.Items {
		item, err := unmarshalReceiver(rInner)
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
	e.ContentID = r.ContentID
	e.TeaserTitle = getTeaserTitle(&r)
	e.ExternalContent = r.ExternalContent
	e.Struct = r.Struct
	return e
}

func unmarshalHTMLContent(r Receiver) (h *HTMLContent) {
	h = &HTMLContent{}
	h.ContentID = r.ContentID
	h.Code = r.Code
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
	if len(r.Photo) > 0 {
		photo := r.Photo[0]
		p.Photo = &photo
	}
	p.NavContext = r.NavContext
	p.AnalyticsCategory = r.AnalyticsCategory
	p.AdvertisingCategory = r.AdvertisingCategory

	return p
}

func unmarshalAudio(r Receiver) (a *Audio) {
	a = &Audio{}
	a.ContentID = r.ContentID
	a.TeaserTitle = getTeaserTitle(&r)
	a.TeaserText = r.TeaserText
	a.PublicationDate = r.PublicationDate
	a.Authors = r.AuthorObjects
	a.CanonicalURL = r.CanonicalURL
	a.URL = r.URL
	a.Stream = r.Stream
	a.NavContext = r.NavContext
	a.AnalyticsCategory = r.AnalyticsCategory
	a.AdvertisingCategory = r.AdvertisingCategory
	for _, rInner := range r.Media {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			a.Media = append(a.Media, item)
		}
	}

	return a
}

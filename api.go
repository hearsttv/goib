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

var Log = l5g.Logger(l5g.LogAll)

type httpClient func(url string) ([]byte, error)

// API is the entrance point for interacting with the IB API
type API interface {
	Entry(channel string, entrytype string, params url.Values) (Item, error)
	Search(channel string, query string, params url.Values) (*Collection, error)
	Content(channel string, contentID int, params url.Values) (Item, error)
	ContentMedia(channel string, contentID int, params url.Values) ([]Item, error)
	ContentItems(channel string, contentID int, params url.Values) ([]Item, error)
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

func doGet(url string) (result []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	Log.Debug("IB API: %s : HTTP Status: %s, Success: %t", url, resp.Status, (err == nil))
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
			Log.Warn("error unmarshalling item from array: %v", err)
		} else {
			result = append(result, item)
		}
	}

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
	default:
		return nil, fmt.Errorf("unknonwn response type: %s", r.Type)
	}
}

func unmarshalArticle(r Receiver) (a *Article) {
	a = &Article{}
	a.ContentID = r.ContentID
	a.TeaserTitle = r.TeaserTitle
	a.TeaserText = r.TeaserText
	a.TeaserImage = r.TeaserImage
	a.PublicationDate = r.PublicationDate
	a.Title = r.Title
	a.Text = r.Text
	a.Authors = r.AuthorObjects
	for _, rInner := range r.Media {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			Log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			a.Media = append(a.Media, item)
		}
	}

	return a
}

func unmarshalVideo(r Receiver) (v *Video) {
	v = &Video{}
	v.ContentID = r.ContentID
	v.TeaserTitle = r.TeaserTitle
	v.TeaserText = r.TeaserText
	v.TeaserImage = r.TeaserImage
	v.PublicationDate = r.PublicationDate
	v.Title = r.Title
	v.Flavors = r.Flavors

	return v
}

func unmarshalImage(r Receiver) (i *Image) {
	i = &Image{}
	i.ContentID = r.ContentID
	i.TeaserTitle = r.TeaserTitle
	i.TeaserText = r.TeaserText
	i.TeaserImage = r.TeaserImage
	i.PublicationDate = r.PublicationDate
	i.Title = r.Title
	i.Caption = r.Caption
	i.Author = r.Author
	i.URLs = r.URLs

	return i
}

func unmarshalGallery(r Receiver) (g *Gallery) {
	g = &Gallery{}
	g.ContentID = r.ContentID
	g.TeaserTitle = r.TeaserTitle
	g.TeaserText = r.TeaserText
	g.TeaserImage = r.TeaserImage
	g.PublicationDate = r.PublicationDate
	g.Title = r.Title
	for _, rInner := range r.Media {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			Log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			g.Media = append(g.Media, item)
		}
	}
	for _, rInner := range r.Items {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			Log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			g.Items = append(g.Items, item)
		}
	}

	return g
}

func unmarshalCollection(r Receiver) (c *Collection) {
	c = &Collection{}
	c.ContentID = r.ContentID
	c.TeaserTitle = r.TeaserTitle
	c.CollectionName = r.CollectionName
	c.ContentName = r.ContentName
	c.TotalCount = r.TotalCount
	c.StartIndex = r.StartIndex
	for _, rInner := range r.Items {
		item, err := unmarshalReceiver(rInner)
		if err != nil {
			Log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			c.Items = append(c.Items, item)
		}
	}
	if len(r.Settings) > 0 {
		c.Settings = r.Settings[0]
	}

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
			Log.Warn("error unmarshalling sub-object: %v", err)
		} else {
			s.Items = append(s.Items, item)
		}
	}

	return s
}

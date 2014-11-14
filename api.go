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
	Entry(channel string, entrytype string, params url.Values) (interface{}, error)
	Article(channel string, contentID int, params url.Values) (Article, error)
	Video(channel string, contentID int, params url.Values) (Video, error)
	Image(channel string, contentID int, params url.Values) (Image, error)
	Gallery(channel string, contentID int, params url.Values) (Gallery, error)
	Search(channel string, query string, params url.Values) (SearchResult, error)
	Content(channel string, contentID int, params url.Values) (interface{}, error)
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

func (api *api) Entry(channel string, entrytype string, params url.Values) (entry interface{}, err error) {
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

func (api *api) Article(channel string, contentID int, params url.Values) (article Article, err error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID)

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := doGet(uri)
	if err != nil {
		return article, err
	}

	iface, err := unmarshalResponse(bytes)
	if err != nil {
		return article, err
	}
	switch t := iface.(type) {
	case Article:
		return iface.(Article), nil
	default:
		return article, fmt.Errorf("invalid object type returned when getting video: %v", t)
	}
}

func (api *api) Video(channel string, contentID int, params url.Values) (video Video, err error) {
	uri := strings.Replace(api.deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID)

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := doGet(uri)
	if err != nil {
		return video, err
	}

	iface, err := unmarshalResponse(bytes)
	if err != nil {
		return video, err
	}
	switch t := iface.(type) {
	case Video:
		return iface.(Video), nil
	default:
		return video, fmt.Errorf("invalid object type returned when getting video: %v", t)
	}
}

func (api *api) Image(channel string, contentID int, params url.Values) (image Image, err error) {
	content, err := api.Content(channel, contentID, params)
	if err != nil {
		return image, err
	}

	switch t := content.(type) {
	case Image:
		return content.(Image), nil
	default:
		return image, fmt.Errorf("invalid object type returned when getting image: %v", t)
	}
}

func (api *api) Gallery(channel string, contentID int, params url.Values) (gallery Gallery, err error) {
	content, err := api.Content(channel, contentID, params)
	if err != nil {
		return gallery, err
	}

	switch t := content.(type) {
	case Gallery:
		return content.(Gallery), nil
	default:
		return gallery, fmt.Errorf("invalid object type returned when getting gallery: %v", t)
	}
}

func (api *api) Search(channel string, query string, params url.Values) (s SearchResult, err error) {
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

	iface, err := unmarshalResponse(bytes)
	if err != nil {
		return s, err
	}
	switch t := iface.(type) {
	case SearchResult:
		return iface.(SearchResult), nil
	default:
		return s, fmt.Errorf("invalid object type returned by search: %v", t)
	}
}

func (api *api) Content(channel string, contentID int, params url.Values) (interface{}, error) {
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

func unmarshalResponse(bytes []byte) (interface{}, error) {
	var r Receiver
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return unmarshalReceiver(r)
}

func unmarshalReceiver(r Receiver) (interface{}, error) {
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

func unmarshalArticle(r Receiver) (a Article) {
	a.ContentID = r.ContentID
	a.TeaserTitle = r.TeaserTitle
	a.TeaserText = r.TeaserText
	a.TeaserImage = r.TeaserImage
	a.Title = r.Title
	a.Text = r.Text
	a.Author = r.Author

	return a
}

func unmarshalVideo(r Receiver) (v Video) {
	v.ContentID = r.ContentID
	v.TeaserTitle = r.TeaserTitle
	v.TeaserText = r.TeaserText
	v.TeaserImage = r.TeaserImage
	v.Title = r.Title
	v.Flavors = r.Flavors

	return v
}

func unmarshalImage(r Receiver) (i Image) {
	i.ContentID = r.ContentID
	i.TeaserTitle = r.TeaserTitle
	i.TeaserText = r.TeaserText
	i.TeaserImage = r.TeaserImage
	i.Title = r.Title
	i.Caption = r.Caption
	i.Author = r.Author
	i.URLs = r.URLs

	return i
}

func unmarshalGallery(r Receiver) (g Gallery) {
	g.ContentID = r.ContentID
	g.TeaserTitle = r.TeaserTitle
	g.TeaserText = r.TeaserText
	g.TeaserImage = r.TeaserImage
	g.Title = r.Title
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

func unmarshalCollection(r Receiver) (c Collection) {
	c.ContentID = r.ContentID
	c.TeaserTitle = r.TeaserTitle
	c.CollectionName = r.CollectionName
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

	return c
}

func unmarshalSearch(r Receiver) (s SearchResult) {
	s.Type = SearchType // TODO: unnecessary, remove
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

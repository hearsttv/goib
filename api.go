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

var _httpClient = func(url string) (result []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	log.Debug("IB API: %s : HTTP Status: %s, Success: %t", url, resp.Status, (err == nil))
	return result, err
}

// API is the entrance point for interacting with the IB API
type API interface {
	Entry(channel string, entrytype string, params url.Values) (Collection, error)
	Article(channel string, contentID int, params url.Values) (Article, error)
	Video(channel string, contentID int, params url.Values) (Video, error)
	Search(channel string, query string, params url.Values) (SearchResult, error)
}

// NewAPI constructs an API object for the given channel
func NewAPI() API {
	return &api{
		httpClient: _httpClient,
	}
}

type api struct {
	httpClient httpClient
}

func (api *api) Entry(channel string, entrytype string, params url.Values) (coll Collection, err error) {
	uri := strings.Replace(deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "entry", 1)
	uri += "/" + entrytype

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := api.httpClient(uri)
	if err != nil {
		return coll, err
	}

	err = json.Unmarshal(bytes, &coll)
	return coll, err
}

func (api *api) Article(channel string, contentID int, params url.Values) (article Article, err error) {
	uri := strings.Replace(deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID)

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := api.httpClient(uri)
	if err != nil {
		return article, err
	}

	err = json.Unmarshal(bytes, &article)
	if err != nil {
		return article, err
	} else if article.Type != ArticleType {
		return article, fmt.Errorf("content ID %d not article: %s", contentID, article.Type)
	}

	return article, err
}

func (api *api) Video(channel string, contentID int, params url.Values) (video Video, err error) {
	uri := strings.Replace(deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "content", 1)
	uri += "/" + strconv.Itoa(contentID)

	if len(params) > 0 {
		uri += "?" + params.Encode()
	}

	bytes, err := api.httpClient(uri)
	if err != nil {
		return video, err
	}

	err = json.Unmarshal(bytes, &video)
	if err != nil {
		return video, err
	} else if video.Type != VideoType {
		return video, fmt.Errorf("content ID %d not article: %s", contentID, video.Type)
	}

	return video, err
}

func (api *api) Search(channel string, query string, params url.Values) (result SearchResult, err error) {
	uri := strings.Replace(deliveryURL, "{channel}", channel, 1)
	uri = strings.Replace(uri, "{service}", "search", 1)

	if params == nil {
		params = url.Values{}
	}
	params.Set("q", query)
	uri += "?" + params.Encode()

	bytes, err := api.httpClient(uri)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	return result, err
}

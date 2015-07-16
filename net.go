package goib

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

var netClient = &http.Client{
	Timeout:   time.Second * 30,
	Transport: netTransport,
}

// doGet is a method on the api object, but it's worth separating out here for clarity
func (api *api) doGet(url string) (result []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("this shouldn't happen: %v", err)
		return nil, err
	}
	resp, err := api.client.Do(req)
	if err != nil {
		log.Debug("got error response for URL %s: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	statusCode, status := resp.StatusCode, resp.Status
	if statusCode != 200 {
		return nil, fmt.Errorf("IB returned an error: %s: %s", status, url)
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v: %s", err, url)
	}

	log.Trace("%s : SUCCESS", url)
	return result, nil
}

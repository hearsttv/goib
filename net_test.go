package goib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_doGet_canary(t *testing.T) {
	testSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "I leik milk!!1")
	}))
	defer testSvr.Close()

	a := NewAPI().(*api)

	resp, err := a.doGet(testSvr.URL)
	assert.Nil(t, err)
	assert.Equal(t, []byte("I leik milk!!1\n"), resp)
}

func Test_doGet_errorResponse(t *testing.T) {
	testSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "I'm a teapot!", http.StatusTeapot)
	}))
	defer testSvr.Close()

	a := NewAPI().(*api)

	_, err := a.doGet(testSvr.URL)
	assert.NotNil(t, err)
	assert.Equal(t, "IB returned an error: 418 I'm a teapot: "+testSvr.URL, err.Error())
}

func Test_doGet_timeout(t *testing.T) {
	oldTimeout := netClient.Timeout
	netClient.Timeout = time.Millisecond

	testSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 2)
		fmt.Fprintln(w, "I leik milk!!1")
	}))
	defer testSvr.Close()

	a := NewAPI().(*api)

	_, err := a.doGet(testSvr.URL)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("Get %s: read tcp %s: use of closed network connection", testSvr.URL, testSvr.URL[7:]), err.Error())

	netClient.Timeout = oldTimeout
}

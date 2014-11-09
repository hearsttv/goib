package mocks

import "github.com/Hearst-DD/goib"
import "github.com/stretchr/testify/mock"

import "net/url"

type API struct {
	mock.Mock
}

func (m *API) Entry(channel string, entrytype string, params url.Values) (goib.Collection, error) {
	ret := m.Called(channel, entrytype, params)

	r0 := ret.Get(0).(goib.Collection)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) Article(channel string, contentID int, params url.Values) (goib.Article, error) {
	ret := m.Called(channel, contentID, params)

	r0 := ret.Get(0).(goib.Article)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) Video(channel string, contentID int, params url.Values) (goib.Video, error) {
	ret := m.Called(channel, contentID, params)

	r0 := ret.Get(0).(goib.Video)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) Search(channel string, query string, params url.Values) (goib.SearchResult, error) {
	ret := m.Called(channel, query, params)

	r0 := ret.Get(0).(goib.SearchResult)
	r1 := ret.Error(1)

	return r0, r1
}

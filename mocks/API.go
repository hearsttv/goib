package mocks

import "github.com/Hearst-DD/goib"
import "github.com/stretchr/testify/mock"

import "net/url"

type API struct {
	mock.Mock
}

func (m *API) Entry(channel string, entrytype string, params url.Values) (goib.Item, error) {
	ret := m.Called(channel, entrytype, params)

	r0 := ret.Get(0).(goib.Item)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) Search(channel string, query string, params url.Values) (*goib.Collection, error) {
	ret := m.Called(channel, query, params)

	var r0 *goib.Collection
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*goib.Collection)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) Content(channel string, contentID int, params url.Values) (goib.Item, error) {
	ret := m.Called(channel, contentID, params)

	r0 := ret.Get(0).(goib.Item)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) ContentMedia(channel string, contentID int, params url.Values) ([]goib.Item, error) {
	ret := m.Called(channel, contentID, params)

	var r0 []goib.Item
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]goib.Item)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) ContentItems(channel string, contentID int, params url.Values) ([]goib.Item, error) {
	ret := m.Called(channel, contentID, params)

	var r0 []goib.Item
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]goib.Item)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) Closings(channel string, filter goib.ClosingsFilter, providerID ...string) (goib.ClosingsResponse, error) {
	ret := m.Called(channel, filter, providerID)

	r0 := ret.Get(0).(goib.ClosingsResponse)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *API) UnmarshalReceiver(r goib.Receiver) (goib.Item, error) {
	ret := m.Called(r)

	r0 := ret.Get(0).(goib.Item)
	r1 := ret.Error(1)

	return r0, r1
}

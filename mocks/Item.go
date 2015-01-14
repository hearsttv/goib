package mocks

import "github.com/Hearst-DD/goib"
import "github.com/stretchr/testify/mock"

type Item struct {
	mock.Mock
}

func (m *Item) GetType() goib.ItemType {
	ret := m.Called()

	r0 := ret.Get(0).(goib.ItemType)

	return r0
}
func (m *Item) GetContentID() int {
	ret := m.Called()

	r0 := ret.Get(0).(int)

	return r0
}
func (m *Item) GetTeaserTitle() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *Item) GetTeaserText() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *Item) GetPublicationDate() int64 {
	ret := m.Called()

	r0 := ret.Get(0).(int64)

	return r0
}

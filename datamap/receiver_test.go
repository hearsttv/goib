package datamap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalBareBase(t *testing.T) {
	//arrange
	base := &bareBase{}
	receiver := Receiver{
		Type:                ArticleType,
		Hash:                "hash",
		ContentID:           1,
		ContentName:         "content",
		Title:               "Title",
		CreationDate:        "2015-01-01T00:00:00-03:00",
		PublicationDate:     "2016-01-01T00:00:00-03:00",
		NavContext:          []string{"home", "article"},
		ValidFrom:           "2017-01-01T00:00:00-03:00",
		ValidTo:             "2018-01-01T00:00:00-03:00",
		AnalyticsCategory:   "item",
		AdvertisingCategory: "item_adv",
	}

	//act
	base.unmarshalBareBase(&receiver)

	//assert
	assert.ObjectsAreEqual(receiver.Type, base.Type)
	assert.ObjectsAreEqual(receiver.Hash, base.Hash)
	assert.ObjectsAreEqual(receiver.ContentID, base.ContentID)
	assert.ObjectsAreEqual(receiver.ContentName, base.ContentName)
	assert.ObjectsAreEqual(receiver.Title, base.Title)
	assert.ObjectsAreEqual(receiver.CreationDate, base.CreationDate)
	assert.ObjectsAreEqual(receiver.PublicationDate, base.PublicationDate)
	assert.ObjectsAreEqual(receiver.NavContext, base.NavContext)
	assert.ObjectsAreEqual(receiver.ValidFrom, base.ValidFrom)
	assert.ObjectsAreEqual(receiver.ValidTo, base.ValidTo)
	assert.ObjectsAreEqual(receiver.AnalyticsCategory, base.AnalyticsCategory)
	assert.ObjectsAreEqual(receiver.AdvertisingCategory, base.AdvertisingCategory)
}

func Test_UnmarshalFullBase(t *testing.T) {
	//arrange
	base := &fullBase{}
	receiver := Receiver{
		Type:                ArticleType,
		Hash:                "hash",
		ContentID:           1,
		ContentName:         "content",
		Title:               "Title",
		CreationDate:        "2015-01-01T00:00:00-03:00",
		PublicationDate:     "2016-01-01T00:00:00-03:00",
		NavContext:          []string{"home", "article"},
		ValidFrom:           "2017-01-01T00:00:00-03:00",
		ValidTo:             "2018-01-01T00:00:00-03:00",
		AnalyticsCategory:   "item",
		AdvertisingCategory: "item_adv",
		TeaserTitle:         "teaser title",
		TeaserText:          "teaser text",
		TeaserImage:         "http://teaser.com/image",
		Author:              "The Author",
		Authors: []Person{
			{
				FullName:      "Full Name Smith",
				Email:         "full@name.smith",
				Bio:           "Smith was born, then lived and died.",
				Photo:         Image{},
				FacebookURL:   "https://facebook.com",
				GooglePlusURL: "https://plus.google.com",
				TwitterURL:    "https://twitter.com",
			},
		},
		EditorialComment: "I'm here to edit comments",
		Copyright:        "Everything was copied in the right way",
		Copyrights:       []Copyright{},
		Media: []Receiver{
			{
				Type: "VIDEO",
			},
			{
				Type: "IMAGE",
			},
		},
		CanonicalURL:      "http://url.com/index.html",
		URL:               "http://url.com/index.html?orderby=date&view=list",
		Categories:        nil,
		Struct:            nil,
		CommentingEnabled: true, //important to be true to test so it's not mistaken by the default value
		NoFollow:          true,
		NotSearchable:     true,
		Period:            "a month",
		Keywords:          "jail,locker,door,chest,keychain",
	}

	//act
	err := base.unmarshalFullBase(&receiver)

	//assert
	assert.Nil(t, err)
	assert.ObjectsAreEqual(receiver.Type, base.Type)
	assert.ObjectsAreEqual(receiver.Hash, base.Hash)
	assert.ObjectsAreEqual(receiver.ContentID, base.ContentID)
	assert.ObjectsAreEqual(receiver.ContentName, base.ContentName)
	assert.ObjectsAreEqual(receiver.Title, base.Title)
	assert.ObjectsAreEqual(receiver.CreationDate, base.CreationDate)
	assert.ObjectsAreEqual(receiver.PublicationDate, base.PublicationDate)
	assert.ObjectsAreEqual(receiver.NavContext, base.NavContext)
	assert.ObjectsAreEqual(receiver.ValidFrom, base.ValidFrom)
	assert.ObjectsAreEqual(receiver.ValidTo, base.ValidTo)
	assert.ObjectsAreEqual(receiver.AnalyticsCategory, base.AnalyticsCategory)
	assert.ObjectsAreEqual(receiver.AdvertisingCategory, base.AdvertisingCategory)
	assert.ObjectsAreEqualValues(receiver, base)
}

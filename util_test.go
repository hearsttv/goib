package goib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMediaShouldReturnInputIfItIsMedia(t *testing.T) {
	article := &Article{}
	result := ExtractMedia(article)
	assert.Equal(t, article, result[0], "article should be returned as single list item")

	video := &Video{}
	result = ExtractMedia(video)
	assert.Equal(t, video, result[0], "video should be returned as single list item")

	image := &Image{}
	result = ExtractMedia(image)
	assert.Equal(t, image, result[0], "image should be returned as single list item")

	gallery := &Gallery{}
	result = ExtractMedia(gallery)
	assert.Equal(t, gallery, result[0], "gallery should be returned as single list item")
}

func TestNilOrEmptyCollectionReturnsEmptySlice(t *testing.T) {
	result := ExtractMedia(nil)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))

	coll := Collection{}
	result = ExtractMedia(&coll)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))
}

func TestExtractMediaShouldScanCollectionsAndPullOutMedia(t *testing.T) {
	coll := &Collection{}
	article := &Article{}
	video := &Video{}
	image := &Image{}
	gallery := &Gallery{}
	coll.Items = []Item{article, video, image, gallery}

	result := ExtractMedia(coll)
	assert.Equal(t, article, result[0], "article should be first element")
	assert.Equal(t, video, result[1], "video should be second element")
	assert.Equal(t, image, result[2], "image should be thrid element")
	assert.Equal(t, gallery, result[3], "gallery should be fourth element")
}

func TestExtractMediaShouldScanSearchResultsForMedia(t *testing.T) {
	sr := &Collection{}
	article := &Article{}
	video := &Video{}
	image := &Image{}
	gallery := &Gallery{}
	sr.Items = []Item{article, video, image, gallery}

	result := ExtractMedia(sr)
	assert.Equal(t, article, result[0], "article should be first element")
	assert.Equal(t, video, result[1], "video should be second element")
	assert.Equal(t, image, result[2], "image should be thrid element")
	assert.Equal(t, gallery, result[3], "gallery should be fourth element")
}

func TestExtractMediaShouldRecurseIntoSubcollections(t *testing.T) {
	coll := &Collection{}
	sr := &Collection{}
	subcoll := &Collection{}
	article := &Article{}
	video := &Video{}
	image := &Image{}
	gallery := &Gallery{}
	subcoll.Items = []Item{article, video}
	sr.Items = []Item{image, gallery}
	coll.Items = []Item{subcoll, sr}

	result := ExtractMedia(coll)
	assert.Equal(t, article, result[0], "article should be first element")
	assert.Equal(t, video, result[1], "video should be second element")
	assert.Equal(t, image, result[2], "image should be thrid element")
	assert.Equal(t, gallery, result[3], "gallery should be fourth element")
}

func TestIteratorShouldReturnOnceIfRootIsMedia(t *testing.T) {
	article := &Article{}
	var i = 0
	for node := range MediaIterator(article) {
		assert.Equal(t, 0, i, "should only loop once")
		assert.Equal(t, article, node.Media, "should have gotten article back out")
		assert.Nil(t, node.ParentCollection, "parent should be nil")
		i++
	}

	video := &Video{}
	i = 0
	for node := range MediaIterator(video) {
		assert.Equal(t, 0, i, "should only loop once")
		assert.Equal(t, video, node.Media, "should have gotten video back out")
		assert.Nil(t, node.ParentCollection, "parent should be nil")
		i++
	}

	image := &Image{}
	i = 0
	for node := range MediaIterator(image) {
		assert.Equal(t, 0, i, "should only loop once")
		assert.Equal(t, image, node.Media, "should have gotten image back out")
		assert.Nil(t, node.ParentCollection, "parent should be nil")
		i++
	}

	gallery := &Gallery{}
	i = 0
	for node := range MediaIterator(gallery) {
		assert.Equal(t, 0, i, "should only loop once")
		assert.Equal(t, gallery, node.Media, "should have gotten gallery back out")
		assert.Nil(t, node.ParentCollection, "parent should be nil")
		i++
	}
}

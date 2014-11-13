package goib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMediaShouldReturnInputIfItIsMedia(t *testing.T) {
	article := Article{}
	result := ExtractMedia(article)
	assert.Equal(t, article, result[0], "article should be returned as single list item")

	video := Video{}
	result = ExtractMedia(video)
	assert.Equal(t, video, result[0], "video should be returned as single list item")

	image := Image{}
	result = ExtractMedia(image)
	assert.Equal(t, image, result[0], "image should be returned as single list item")

	gallery := Gallery{}
	result = ExtractMedia(gallery)
	assert.Equal(t, gallery, result[0], "gallery should be returned as single list item")
}

func TestNilOrEmptyCollectionReturnsEmptySlice(t *testing.T) {
	result := ExtractMedia(nil)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))

	coll := Collection{}
	result = ExtractMedia(coll)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))
}

func TestExtractMediaShouldScanCollectionsAndPullOutMedia(t *testing.T) {
	coll := Collection{}
	article := Article{}
	video := Video{}
	image := Image{}
	gallery := Gallery{}
	coll.Items = []interface{}{article, video, image, gallery}

	result := ExtractMedia(coll)
	assert.Equal(t, article, result[0], "article should be first element")
	assert.Equal(t, video, result[1], "video should be second element")
	assert.Equal(t, image, result[2], "image should be thrid element")
	assert.Equal(t, gallery, result[3], "gallery should be fourth element")
}

func TestExtractMediaShouldScanSearchResultsForMedia(t *testing.T) {
	sr := SearchResult{}
	article := Article{}
	video := Video{}
	image := Image{}
	gallery := Gallery{}
	sr.Items = []interface{}{article, video, image, gallery}

	result := ExtractMedia(sr)
	assert.Equal(t, article, result[0], "article should be first element")
	assert.Equal(t, video, result[1], "video should be second element")
	assert.Equal(t, image, result[2], "image should be thrid element")
	assert.Equal(t, gallery, result[3], "gallery should be fourth element")
}

func TestExtractMediaShouldRecurseIntoSubcollections(t *testing.T) {
	coll := Collection{}
	sr := SearchResult{}
	subcoll := Collection{}
	article := Article{}
	video := Video{}
	image := Image{}
	gallery := Gallery{}
	subcoll.Items = []interface{}{article, video}
	sr.Items = []interface{}{image, gallery}
	coll.Items = []interface{}{subcoll, sr}

	result := ExtractMedia(coll)
	assert.Equal(t, article, result[0], "article should be first element")
	assert.Equal(t, video, result[1], "video should be second element")
	assert.Equal(t, image, result[2], "image should be thrid element")
	assert.Equal(t, gallery, result[3], "gallery should be fourth element")
}

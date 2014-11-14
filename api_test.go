package goib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFailGracefullyOnEmptyResponse(t *testing.T) {
	svr, a := setupServerAndAPI(emptyJSON)
	defer svr.Close()

	assertGracefulFailOnAllAPIMethods(t, a)
}

func TestShouldFailGracefullyOnMalformedResponse(t *testing.T) {
	svr, a := setupServerAndAPI(missingCloseBracketJSON)
	defer svr.Close()

	assertGracefulFailOnAllAPIMethods(t, a)
}

func TestShouldFailGracefullyOnHTMLResponse(t *testing.T) {
	svr, a := setupServerAndAPI(badResponseHTML)
	defer svr.Close()

	assertGracefulFailOnAllAPIMethods(t, a)
}

func TestShouldFailGracefullyOn500Response(t *testing.T) {
	svr, a := setupServerAndAPIWithHTTPStatus(badResponseHTML, 500)
	defer svr.Close()

	assertGracefulFailOnAllAPIMethods(t, a)
}

func TestShouldFailGracefullyOn404Response(t *testing.T) {
	svr, a := setupServerAndAPIWithHTTPStatus(badResponseHTML, 404)
	defer svr.Close()

	assertGracefulFailOnAllAPIMethods(t, a)
}

func TestEntryApi(t *testing.T) {
	svr, a := setupServerAndAPI(entryJSON)
	defer svr.Close()

	response, err := a.Entry("someKrazyChannel", "someKookyKollection", nil)
	if err != nil {
		t.Errorf("error getting entry: %v", err)
	}

	collection := response.(Collection)
	if len(collection.Items) == 0 {
		t.Errorf("zero items returned from entry")
	}
}

func TestArticleApi(t *testing.T) {
	svr, a := setupServerAndAPI(articleJSON)
	defer svr.Close()

	article, err := a.Article("someNuttyChannel", 9137538, nil)
	if err != nil {
		t.Errorf("error getting content: %v", err)
	}
	if article.ContentID != 9137538 {
		t.Errorf("expected 9137538 type but got %d", article.ContentID)
	}
	if article.Title != "George Washington chicken nugget fetches $8,100" {
		t.Errorf("expected 'George Washington chicken nugget fetches $8,100' but got %s", article.Title)
	}
}

func TestVideoApi(t *testing.T) {
	svr, a := setupServerAndAPI(videoJSON)
	defer svr.Close()

	video, err := a.Video("wkrp", 1402356, nil)
	if err != nil {
		t.Errorf("error getting content: %v", err)
	}
	if video.ContentID != 1402356 {
		t.Errorf("expected 1402356 type but got %d", video.ContentID)
	}
	if video.Title != "Advertisers Ready For NFL Kickoff" {
		t.Errorf("expected 'Advertisers Ready For NFL Kickoff' but got %s", video.Title)
	}
}

func TestSearchApi(t *testing.T) {
	svr, a := setupServerAndAPI(searchJSON)
	defer svr.Close()

	search, err := a.Search("wkrp", "nfl", nil)
	if err != nil {
		t.Errorf("error getting content: %v", err)
	}
	if search.TotalCount == 0 {
		t.Errorf("zero results returned from search")
	}
	if search.Keywords != "nfl" {
		t.Errorf("expected keywords 'nfl' but got %s", search.Keywords)
	}
}

func TestContentApiShouldParseCollectionsOfCollections(t *testing.T) {
	svr, a := setupServerAndAPI(multitieredCollectionJSON)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	if err != nil {
		t.Errorf("error getting collection: %v", err)
	}

	collection := response.(Collection)
	if len(collection.Items) == 0 {
		t.Errorf("zero items returned from entry")
	}
	innerCollection := collection.Items[0].(Collection)
	assert.Equal(t, 8, len(innerCollection.Items), "first inner collection should have 8 items but only has %d", len(innerCollection.Items))
}

func TestContentAPIShouldParseImageType(t *testing.T) {
	svr, a := setupServerAndAPI(imageJSON)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	if err != nil {
		t.Errorf("error getting image: %v", err)
	}

	img := response.(Image)
	assert.Equal(t, 29283344, img.ContentID, "expected COID 29283344 but got %d", img.ContentID)
}

func TestContentAPIShouldParseGalleryType(t *testing.T) {
	svr, a := setupServerAndAPI(galleryJSON)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	if err != nil {
		t.Errorf("error getting gallery: %v", err)
	}

	g := response.(Gallery)
	assert.Equal(t, 29283428, g.ContentID, "expected COID 29283428 but got %d", g.ContentID)

	assert.Equal(t, 53, len(g.Items), "expected 50 image items in gallery but got %d", len(g.Items))
}

func setupServerAndAPI(cannedResponse string) (*httptest.Server, API) {
	testSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, cannedResponse)
	}))

	a := NewAPI().(*api)
	a.deliveryURL = testSvr.URL

	return testSvr, a
}

func setupServerAndAPIWithHTTPStatus(cannedResponse string, status int) (*httptest.Server, API) {
	testSvr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "floopy/yowza")
		w.WriteHeader(500)
		fmt.Fprintln(w, cannedResponse)
	}))

	a := NewAPI().(*api)
	a.deliveryURL = testSvr.URL

	return testSvr, a
}

func assertGracefulFailOnAllAPIMethods(t *testing.T, a API) {
	_, err := a.Entry("someKrazyChannel", "someKookyKollection", nil)
	assert.NotNil(t, err, "error should not be nil")

	_, err = a.Content("someKrazyChannel", 12345, nil)
	assert.NotNil(t, err, "error should not be nil")

	_, err = a.Article("someKrazyChannel", 12345, nil)
	assert.NotNil(t, err, "error should not be nil")

	_, err = a.Video("someKrazyChannel", 12345, nil)
	assert.NotNil(t, err, "error should not be nil")

	_, err = a.Image("someKrazyChannel", 12345, nil)
	assert.NotNil(t, err, "error should not be nil")

	_, err = a.Gallery("someKrazyChannel", 12345, nil)
	assert.NotNil(t, err, "error should not be nil")
}

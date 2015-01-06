package goib

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

	collection := response.(*Collection)
	if len(collection.Items) == 0 {
		t.Errorf("zero items returned from entry")
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

	collection := response.(*Collection)
	if len(collection.Items) == 0 {
		t.Errorf("zero items returned from entry")
	}
	innerCollection := collection.Items[0].(*Collection)
	assert.Equal(t, 8, len(innerCollection.Items), "first inner collection should have 8 items but only has %d", len(innerCollection.Items))
}

func TestContentAPIShouldParseImageType(t *testing.T) {
	svr, a := setupServerAndAPI(imageJSON)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	if err != nil {
		t.Errorf("error getting image: %v", err)
	}

	img := response.(*Image)
	assert.Equal(t, 29283344, img.ContentID, "expected COID 29283344 but got %d", img.ContentID)
}

func TestContentAPIShouldParseGalleryType(t *testing.T) {
	svr, a := setupServerAndAPI(galleryJSON)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	if err != nil {
		t.Errorf("error getting gallery: %v", err)
	}

	g := response.(*Gallery)
	assert.Equal(t, 29283428, g.ContentID, "expected COID 29283428 but got %d", g.ContentID)

	assert.Equal(t, 53, len(g.Items), "expected 50 image items in gallery but got %d", len(g.Items))
}

func TestContentAPIShouldParseHTMLContentType(t *testing.T) {
	svr, a := setupServerAndAPI(htmlContent)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	if err != nil {
		t.Errorf("error getting HTML: %v", err)
	}

	h := response.(*HTMLContent)
	assert.Equal(t, 14277264, h.ContentID, "expected COID 14277264 but got %d", h.ContentID)
	assert.True(t, strings.HasPrefix(h.Code, "\n\n<!-- Slideshow Widget"), "code field should be populated")
}

func TestShouldUnmarshallArticleMedia(t *testing.T) {
	r := Receiver{
		Type: ArticleType,
		Media: []Receiver{
			Receiver{Type: VideoType, ContentID: 123},
			Receiver{Type: ImageType, ContentID: 456},
		},
	}

	a := unmarshalArticle(r)
	assert.Equal(t, 2, len(a.Media), "media should have two elements")
}

func TestShouldUnmarshallSettingsForCollection(t *testing.T) {
	svr, a := setupServerAndAPI(collectionWithSettingsJSON)
	defer svr.Close()

	response, err := a.Content("someKrazyChannel", 12345, nil)
	c := response.(*Collection)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 3, len(c.Settings), "should have found 3 setings")
	assert.Equal(t, "hourly", c.Settings["collection.WeatherIndicatorType"], "should have 'hourly' for specific setting")
}

func TestClosingsCount(t *testing.T) {
	svr, a := setupServerAndAPI(closingsCount)
	defer svr.Close()

	closings, err := a.Closings("someKrazyChannel", ClosingsCount)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 3, closings.Count.Total, "should have found 3 closings")
}

func TestClosingsAll(t *testing.T) {
	svr, a := setupServerAndAPI(closingsAll)
	defer svr.Close()

	closings, err := a.Closings("someKrazyChannel", ClosingsAll)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 4, len(closings.Institutions), "should have found 4 institutions")
}

func TestClosingsClosed(t *testing.T) {
	svr, a := setupServerAndAPI(closingsClosed)
	defer svr.Close()

	closings, err := a.Closings("someKrazyChannel", ClosingsClosed)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 4, len(closings.ClosedInstitutions), "should have found 4 institutions")
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

	_, err = a.Search("someKrazyChannel", "kweery", nil)
	assert.NotNil(t, err, "error should not be nil")
}

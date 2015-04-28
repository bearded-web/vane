package checks

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

func TestRSS(t *testing.T) {
	fakeBody, err := ioutil.ReadFile(rssfixtures + "/wordpress-3.5.htm")
	assert.NoError(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fakeBody)
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)
	link, err := RSSURL(s)
	assert.NoError(t, err)
	assert.Equal(t, "http://lamp-wp/wordpress-3.5/?feed=rss2", link)
}

func TestNoRSS(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)
	link, err := RSSURL(s)
	assert.NoError(t, err)
	assert.Empty(t, link)

	s, _ = site.NewSite("fakeHTTPaddress/")
	_, err = RSSURL(s)
	assert.Error(t, err)
}

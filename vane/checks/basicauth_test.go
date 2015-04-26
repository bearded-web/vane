package checks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

func TestHasBacicAuth(t *testing.T) {
	tsTrue := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusUnauthorized
		http.Error(w, http.StatusText(code), code)
	}))
	defer tsTrue.Close()

	s, _ := site.NewSite(tsTrue.URL)
	has, err := HasBasicAuth(s)
	assert.NoError(t, err)
	assert.True(t, has)

	tsFalse := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusOK
		http.Error(w, http.StatusText(code), code)
	}))
	defer tsFalse.Close()

	s, _ = site.NewSite(tsFalse.URL)
	has, err = HasBasicAuth(s)
	assert.NoError(t, err)
	assert.False(t, has)

	s, _ = site.NewSite("http://127.0.0.1:999/")
	_, err = HasBasicAuth(s)
	assert.Error(t, err)
}

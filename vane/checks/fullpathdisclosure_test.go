package checks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

var (
	trueBody  = []byte("blablabla faTaL error blabla")
	falseBody = []byte("A")
)

func TestFullPasthDisclosure(t *testing.T) {
	assert.True(t, fullPathDisclosure(trueBody))
	assert.False(t, fullPathDisclosure(falseBody))
}

func TestFullPasthDisclosureTrue(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(trueBody)
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)

	value, err := FullPathDisclosure(s)
	assert.NoError(t, err)
	assert.True(t, value)
}

func TestFullPasthDisclosureFalse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(falseBody)
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)

	value, err := FullPathDisclosure(s)
	assert.NoError(t, err)
	assert.False(t, value)
}

func TestFullPasthDisclosureError(t *testing.T) {
	s, _ := site.NewSite(fakeHTTPaddress)
	_, err := FullPathDisclosure(s)
	assert.Error(t, err)
}

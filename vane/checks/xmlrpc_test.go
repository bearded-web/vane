package checks

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

func TestXMLRPCTrue(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/xmlrpc.php" {
			t.Fatalf("requested URL %s is wrong", r.URL.Path)
			return
		}
		// The body is lower cased as we match in a case insensitive way
		fmt.Fprint(w, "xml-rpc server accepts post requests only")
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)
	has, err := HasXMLrpc(s)
	assert.NoError(t, err)
	assert.True(t, has)
}

func TestXMLRPCFalse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/xmlrpc.php" {
			t.Fatalf("requested URL %s is wrong", r.URL.Path)
			return
		}
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)
	has, err := HasXMLrpc(s)
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestXMLRPCError(t *testing.T) {
	s, _ := site.NewSite("fakeHTTPaddress9")
	_, err := HasXMLrpc(s)
	assert.Error(t, err)
}

package checks

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

func TestPageHashHomePage(t *testing.T) {
	body := []byte("Hello\n\n\nworld!")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello\n\n<!--I should <script>no longer be</script> there -->\nworld!"))
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)

	hash, err := HomepageHash(s)
	assert.NoError(t, err)
	assert.Equal(t, md5.Sum(body), hash)

	hash, err = Error404Hash(s)
	assert.NoError(t, err)
	assert.Equal(t, md5.Sum(body), hash)
}

func TestClean(t *testing.T) {
	input := []byte("Hello\n\n<!--I should <script>no longer be</script> there -->\nworld!")
	expected := []byte("Hello\n\n\nworld!")
	actual := clean(input)
	assert.Equal(t, expected, actual, fmt.Sprintf("Actual: %s", actual))
}

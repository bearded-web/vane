package checks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

func TestHasLog(t *testing.T) {
	root := http.Dir(has_logfixtures)
	ts := httptest.NewServer(http.FileServer(root))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)

	value, err := HasLog(s, "matches.txt", "PHP Fatal error")
	assert.NoError(t, err)
	assert.True(t, value)

	value, err = HasLog(s, "no_matches.txt", "AAA")
	assert.NoError(t, err)
	assert.False(t, value)

	value, err = HasLog(s, "match_out_of_range.txt", "PHP Fatal error")
	assert.NoError(t, err)
	assert.False(t, value)
}

package site

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedUrl(t *testing.T) {
	var (
		siteURIWithoutSlash = "http://example.com"
		siteURIWithSlash    = siteURIWithoutSlash + "/"

		expectedURL = siteURIWithSlash + "robots.txt"
	)

	s, err := NewSite(siteURIWithoutSlash)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, s.URLFor("robots.txt"))

	s, err = NewSite(siteURIWithSlash)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, s.URLFor("robots.txt"))
}

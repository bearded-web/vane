package checks

import (
	"regexp"

	"github.com/bearded-web/vane/vane/site"
)

var (
	rssPattern = regexp.MustCompile(`<link .* type="application/rss\+xml" .* href="([^"]+)" />`)
)

//ToDO: TESTS
func RSSURL(s site.Site) (string, error) {
	// Get the body to scan
	body, err := s.GetBody("/")
	if err != nil {
		return "", err
	}
	link := rssPattern.Find(body)
	return string(link), nil
}

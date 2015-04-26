package checks

import (
	"github.com/bearded-web/vane/vane/site"
)

// Online checks if the website is up
func Online(s site.Site) bool {
	resp, err := s.Get("/")
	if err != nil {
		return false
	}

	if resp.StatusCode >= 500 {
		return false
	}

	return true
}

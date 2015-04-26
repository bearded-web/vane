package checks

import (
	"github.com/bearded-web/vane/vane/site"
)

// Online checks if the website is up
func Online(s site.Site) bool {
	_, err := s.Get("/")
	if err != nil {
		return false
	}

	return true
}

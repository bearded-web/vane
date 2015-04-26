package checks

import (
	"net/http"

	"github.com/bearded-web/vane/vane/site"
)

// HasBasicAuth checks if the website supports basic authentification
func HasBasicAuth(s site.Site) (bool, error) {
	req, err := s.Get("/")
	if err != nil {
		return false, err
	}

	return req.StatusCode == http.StatusUnauthorized, nil
}

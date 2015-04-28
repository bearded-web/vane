package checks

import (
	"regexp"

	"github.com/bearded-web/vane/vane/site"
)

var fullPathDisclosureRegexp = regexp.MustCompile(`(?i)Fatal error`)

func fullPathDisclosure(body []byte) bool {
	return fullPathDisclosureRegexp.Match(body)
}

func FullPathDisclosure(s site.Site) (bool, error) {
	body, err := s.GetBody("wp-includes/rss-functions.php")
	if err != nil {
		return false, err
	}

	return fullPathDisclosure(body), nil
}

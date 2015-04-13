package checks

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/bearded-web/vane/vane/site"
	"github.com/bearded-web/vane/vane/utils"
)

var (
	robotsKnownDirs = utils.NewStringSet("/", "/wp-admin/", "/wp-includes/", "/wp-content/")

	// (?mi) are options: multiline mode & case-insensitive
	// ^(?:dis)?allow: 0 or 1 non captured `dis` group at the begining and allow
	// \s* - skip all spaces before a matching group
	robotsTxtPattern = regexp.MustCompile(`(?mi)^(?:dis)?allow:\s*(.*)$`)

	errInvalidOrEmptyRobotsTxt = errors.New("invalid or empty robots.txt")
	errNoRobotsTxt             = errors.New("no robots.txt")
)

// RobotsTxt is an interface which wraps
// operations with robots.txt
type RobotsTxt interface {
	HasRobotsTxt() (bool, error)
	ParseRobotsTxt() ([]string, error)
}

type robotsTxt struct {
	site.Site
}

// NewRobotsTxt returns new RobotsTxt to work with
// a site referenced by the given url
func NewRobotsTxt(s site.Site) (RobotsTxt, error) {
	return newRobotsTxt(s)
}

func newRobotsTxt(s site.Site) (*robotsTxt, error) {
	r := &robotsTxt{
		Site: s,
	}

	return r, nil
}

func (r *robotsTxt) HasRobotsTxt() (bool, error) {
	resp, err := r.Site.Head("robots.txt")
	if err != nil {
		return false, err
	}
	resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

func (r *robotsTxt) ParseRobotsTxt() ([]string, error) {
	has, err := r.HasRobotsTxt()
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errNoRobotsTxt
	}

	body, err := r.GetBody("robots.txt")
	if err != nil {
		return nil, err
	}

	return r.getUrlsFromRobotsTxt(body)
}

func (r *robotsTxt) getUrlsFromRobotsTxt(body []byte) ([]string, error) {
	urls, err := parseRobotsTxt(body)
	if err != nil {
		return nil, err
	}

	// nothing to filter
	// quit here
	if len(urls) == 0 {
		return urls, nil
	}

	filteredUrls := urls[:0]
	for _, item := range urls {
		filteredUrls = append(filteredUrls, r.Site.URLFor(item))
	}

	return filteredUrls, nil
}

func parseRobotsTxt(body []byte) ([]string, error) {
	// -1 means all
	entries := robotsTxtPattern.FindAllSubmatch(body, -1)
	if entries == nil {
		// it doens't match at all
		// return an empty slice
		return nil, errInvalidOrEmptyRobotsTxt
	}

	matches := make([]string, 0, len(entries))
	for _, entry := range entries {
		candidate := string(entry[robotsTxtPattern.NumSubexp()])
		if !robotsKnownDirs.Contains(candidate) {
			matches = append(matches, candidate)
		}
		//ToDo: clear subdirs too
		//ToDo: replace the set with a binary search and delete
	}
	return matches, nil
}

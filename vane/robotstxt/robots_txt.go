package robotstxt

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/bearded-web/vane/vane/utils"
)

var (
	robotsKnownDirs = utils.NewStringSet("/", "/wp-admin/", "/wp-includes/", "/wp-content/")

	// (?mi) are options: multiline mode & case-insensitive
	// ^(?:dis)?allow: 0 or 1 non captured `dis` group at the begining and allow
	// \s* - skip all spaces before a matching group
	robotsTxtPattern = regexp.MustCompile(`(?mi)^(?:dis)?allow:\s*(.*)$`)
)

// RobotsTxt is an interface which wraps
// operations with robots.txt
type RobotsTxt interface {
	HasRobotsTxt() (bool, error)
	ParseRobotsTxt() ([]string, error)
}

type robotsTxt struct {
	uri *url.URL
}

// NewRobotsTxt returns new RobotsTxt to work with
// a site referenced by the given url
func NewRobotsTxt(rawurl string) (RobotsTxt, error) {
	return newRobotsTxt(rawurl)
}

func newRobotsTxt(rawurl string) (*robotsTxt, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	r := &robotsTxt{
		uri: u,
	}

	return r, nil
}

func (r *robotsTxt) robotsURL() string {
	if r.uri.Path != "/robots.txt" {
		r.uri.Path = "/robots.txt"
	}
	//ToDo: cache this value
	return r.uri.String()
}

func (r *robotsTxt) HasRobotsTxt() (bool, error) {
	resp, err := http.Head(r.robotsURL())
	if err != nil {
		return false, err
	}
	resp.Body.Close()

	return resp.StatusCode == 200, nil
}

func (r *robotsTxt) ParseRobotsTxt() ([]string, error) {
	if has, err := r.HasRobotsTxt(); !has || err != nil {
		return nil, err
	}

	body, err := getRobotsTxt(r.robotsURL())
	if err != nil {
		return nil, err
	}

	return r.parseRobotsTxt(body)
}

func (r *robotsTxt) parseRobotsTxt(body []byte) ([]string, error) {
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
		r.uri.Path = item
		filteredUrls = append(filteredUrls, r.uri.String())
	}

	return filteredUrls, nil
}

func getRobotsTxt(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func parseRobotsTxt(body []byte) ([]string, error) {
	// -1 means all
	entries := robotsTxtPattern.FindAllSubmatch(body, -1)
	if entries == nil {
		// it doens't match at all
		// return an empty slice
		return []string{}, nil
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

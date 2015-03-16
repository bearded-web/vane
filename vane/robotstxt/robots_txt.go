package robotstxt

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/bearded-web/vane/vane/utils"
)

const robotsTxtFileName = "robots.txt"

var (
	robotsKnownDirs       = utils.NewStringSet("/", "/wp-admin/", "/wp-includes/", "/wp-content/")
	robotsTxtEntryPattern = regexp.MustCompile(`/^(?:dis)?allow:\s*(.*)$/`)
)

type RobotsTxt interface {
	HasRobotsTxt() (bool, error)
	ParseRobotsTxt() ([]string, error)
}

type robotsTxt struct {
	uri *url.URL
}

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

func (r *robotsTxt) robotsUrl() string {
	if r.uri.Path != "/robots.txt" {
		r.uri.Path = "/robots.txt"
	}
	//ToDo: cache this value
	return r.uri.String()
}

func (r *robotsTxt) HasRobotsTxt() (bool, error) {
	resp, err := http.Head(r.robotsUrl())
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

	body, err := getRobotsTxt(r.robotsUrl())
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
		filteredUrls = append(filteredUrls, r.uri.Path)
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
	entries := robotsTxtEntryPattern.FindAll(body, -1)
	if entries == nil {
		// it doens't match at all
		// return empty slice
		return []string{}, nil
	}

	matches := make([]string, 0, len(entries))
	for _, entry := range entries {
		if candidate := string(entry); !robotsKnownDirs.Contains(candidate) {
			matches = append(matches)
		}
		//ToDo: clear subdirs too
		//ToDo: replace set with binary search and delete
	}

	return matches, nil
}

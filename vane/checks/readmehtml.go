package checks

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/bearded-web/vane/vane/site"
)

var (
	wordPressRegexp = regexp.MustCompile("(?mi)wordpress")
)

const (
	readMePath = "readme.html"
)

func HasReadMe(s site.Site) (bool, error) {
	resp, err := s.Get(readMePath)
	if err != nil {
		return false, err
	}
	return hasReadMeCheckResponse(resp)
}

func hasReadMeCheckResponse(resp *http.Response) (bool, error) {
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return wordPressRegexp.Match(body), nil
}

func ReadMeURL(s site.Site) string {
	return s.URLFor(readMePath)
}

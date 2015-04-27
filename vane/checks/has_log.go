package checks

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/bearded-web/vane/vane/site"
)

func HasLog(s site.Site, log_url string, pattern string) (bool, error) {
	regexpPattern, err := regexp.Compile(pattern)
	if err != nil {
		return false, nil
	}

	req, _ := http.NewRequest("GET", s.URLFor(log_url), nil)
	req.Header.Add("Range", "bytes=0-700")

	resp, err := s.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// ToDo: RuneReader???
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return regexpPattern.Match(body), nil
}

package checks

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/bearded-web/vane/vane/site"
)

const (
	multiSignup    = `wp-signup.php`
	nonMultiSignup = `wp-login.php?action=register`
	disabledReg    = `wp-login\.php\?registration=disabled`
)

var (
	// disabledRegExp  = regexp.MustCompile(`(?i)wp-login\.php\?registration=disabled`)
	multisiteRegExp = regexp.MustCompile(`(?mi)<form id="setupform" method="post" action="[^"]*wp-signup\.php[^"]*">`)
	singleRegExp    = regexp.MustCompile(`(?mi)<form name="registerform" id="registerform" action="[^"]*wp-login\.php[^"]*"`)
)

func RegistrationEnabled(s site.Site) (bool, error) {
	regURL, err := registrationURL(s)
	if err != nil {
		return false, err
	}

	req, err := s.Get(regURL)
	if err != nil {
		return false, err
	}

	defer req.Body.Close()

	switch req.StatusCode {
	case http.StatusFound: // 302
		if strings.Contains(strings.ToUpper(req.Header.Get("Location")), disabledReg) {
			return false, nil
		}
	case http.StatusOK:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return false, err
		}

		if multisiteRegExp.Match(body) || singleRegExp.Match(body) {
			return true, nil
		}
	}

	return false, nil
}

func isMultisite(s site.Site) (bool, error) {
	req, err := s.Get("wp-signup.php")
	if err != nil {
		return false, err
	}

	switch req.StatusCode {
	case http.StatusOK:
		return true, nil

	case http.StatusFound: // 302
		// ToDo: think about req.Location
		locationHeader := req.Header.Get("Location")
		if strings.Contains(locationHeader, nonMultiSignup) {
			return false, nil
		}

		if strings.Contains(locationHeader, multiSignup) {
			return true, nil
		}
	}
	return false, nil
}

func registrationURL(s site.Site) (string, error) {
	multi, err := isMultisite(s)
	if err != nil {
		return "", err
	}

	if multi {
		return multiSignup, nil
	}

	return nonMultiSignup, nil
}

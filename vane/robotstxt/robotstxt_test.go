package robotstxt

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

var robotsfixtures = os.Getenv("FIXTURESPATH") + "/robotstxt/"

func getExpectedUrl(base string) []string {
	return []string{
		base + "/wordpress/admin/",
		base + "/wordpress/wp-admin/",
		base + "/wordpress/secret/",
		base + "/Wordpress/wp-admin/",
		base + "/randomurl/",
	}
}

func TestParseEmptyRobotsTxt(t *testing.T) {
	emptyRes, err := parseRobotsTxt([]byte{})
	assert.Error(t, err, "an error is expected after parsing an empty robots.txt")
	assert.Empty(t, emptyRes, "a result from empty_robots.txt should be empty")
}

func TestParseInvalidRobotsTxt(t *testing.T) {
	var (
		invalidRobotsPath = robotsfixtures + "invalid_robots.txt"
	)

	invalid, err := ioutil.ReadFile(invalidRobotsPath)
	assert.NoError(t, err, "unable to read invalid_robots.txt")

	invalidRes, err := parseRobotsTxt(invalid)
	assert.Error(t, err, "an error is expected after parsing invalid_robots.txt")
	assert.Empty(t, invalidRes, "a result from invalid_robots.txt should be empty")
}

func TestParseValidRobotsTxt(t *testing.T) {
	var (
		validRobotsPath     = robotsfixtures + "valid_robots.txt"
		expectedValidResult = getExpectedUrl("http://example.localhost")
	)

	s, _ := site.NewSite("http://example.localhost")

	r, err := newRobotsTxt(s)
	if !assert.NoError(t, err, "unable to create client") {
		t.FailNow()
	}

	valid, err := ioutil.ReadFile(validRobotsPath)
	if !assert.NoError(t, err, "unable to read valid_robots.txt") {
		t.FailNow()
	}

	result, err := r.getUrlsFromRobotsTxt(valid)
	assert.NoError(t, err, "unexpected error during parsing valid_robots.txt")
	assert.Equal(t, expectedValidResult, result)
}

func TestNoRobotsTxt(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "HEAD" && r.URL.Path == "/robots.txt" {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	s, _ := site.NewSite(ts.URL)
	r, err := NewRobotsTxt(s)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	has, _ := r.HasRobotsTxt()
	assert.False(t, has)
}

func TestRobotsTxt(t *testing.T) {
	var (
		validRobotsPath = robotsfixtures + "valid_robots.txt"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/robots.txt" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {
		case "HEAD":
			w.WriteHeader(http.StatusOK)
		case "GET":
			w.WriteHeader(http.StatusOK)
			valid, err := ioutil.ReadFile(validRobotsPath)
			if !assert.NoError(t, err, "unable to read valid_robots.txt") {
				t.FailNow()
			}
			w.Write(valid)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer ts.Close()

	var (
		expectedValidResult = getExpectedUrl(ts.URL)
	)

	s, _ := site.NewSite(ts.URL)

	r, err := NewRobotsTxt(s)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	actual, err := r.ParseRobotsTxt()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, expectedValidResult, actual)
}

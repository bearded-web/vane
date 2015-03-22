package robotstxt

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var robotsfixtures = os.Getenv("FIXTURESPATH") + "/robotstxt/"

func TestGeneratedUrl(t *testing.T) {
	var (
		siteURIWithoutSlash = "http://example.com"
		siteURIWithSlash    = siteURIWithoutSlash + "/"

		expectedURL = siteURIWithSlash + "robots.txt"
	)

	r, err := newRobotsTxt(siteURIWithoutSlash)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, r.robotsURL())

	r, err = newRobotsTxt(siteURIWithSlash)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, r.robotsURL())
}

func TestParseEmptyRobotsTxt(t *testing.T) {
	emptyRes, err := parseRobotsTxt([]byte{})
	assert.NoError(t, err, "unexpected error during parsing empty robots.txt")
	assert.Empty(t, emptyRes, "a result from empty_robots.txt should be empty")
}

func TestParseInvalidRobotsTxt(t *testing.T) {
	var (
		invalidRobotsPath = robotsfixtures + "invalid_robots.txt"
	)

	invalid, err := ioutil.ReadFile(invalidRobotsPath)
	assert.NoError(t, err, "unable to read invalid_robots.txt")

	invalidRes, err := parseRobotsTxt(invalid)
	assert.NoError(t, err, "unexpected error during parsing invalid_robots.txt")
	assert.Empty(t, invalidRes, "a result from invalid_robots.txt should be empty")
}

func TestParseValidRobotsTxt(t *testing.T) {
	var (
		validRobotsPath     = robotsfixtures + "valid_robots.txt"
		expectedValidResult = []string{
			"http://example.localhost/wordpress/admin/",
			"http://example.localhost/wordpress/wp-admin/",
			"http://example.localhost/wordpress/secret/",
			"http://example.localhost/Wordpress/wp-admin/",
			"http://example.localhost/randomurl/",
		}
	)

	r, err := newRobotsTxt("http://example.localhost")
	if !assert.NoError(t, err, "unable to create client") {
		t.FailNow()
	}

	valid, err := ioutil.ReadFile(validRobotsPath)
	if !assert.NoError(t, err, "unable to read valid_robots.txt") {
		t.FailNow()
	}

	result, err := r.parseRobotsTxt(valid)
	assert.NoError(t, err, "unexpected error during parsing valid_robots.txt")
	assert.Equal(t, expectedValidResult, result)
}

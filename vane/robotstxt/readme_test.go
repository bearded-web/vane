package robotstxt

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var readmefixtures = os.Getenv("FIXTURESPATH") + "/readme/"

func TestReadMeBadStatus(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	testResp := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(buff),
	}

	has, err := hasReadMeCheckResponse(testResp)
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestReadMeCheckBody(t *testing.T) {
	fakeBody, err := os.Open(readmefixtures + "readme-3.2.1.html")
	defer fakeBody.Close()

	if !assert.NoError(t, err) {
		t.FailNow()
	}

	testResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       fakeBody,
	}

	has, err := hasReadMeCheckResponse(testResp)
	assert.NoError(t, err)
	assert.True(t, has)
}

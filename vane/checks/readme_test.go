package checks

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearded-web/vane/vane/site"
)

//ToDO: add assert messages

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

type ErroredReader struct{}

func (e *ErroredReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}

func TestReadMeBodyError(t *testing.T) {
	testResp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(&ErroredReader{}),
	}

	has, err := hasReadMeCheckResponse(testResp)
	assert.Error(t, err)
	assert.False(t, has)
}

func TestReadMeURL(t *testing.T) {
	s, _ := site.NewSite("http://testhost.net/")
	assert.Equal(t, "http://testhost.net/"+readMePath, ReadMeURL(s))
}

func TestHasReadme(t *testing.T) {
	tsReadme := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer tsReadme.Close()

	s, err := site.NewSite(tsReadme.URL)
	assert.NoError(t, err)

	has, err := HasReadMe(s)
	assert.NoError(t, err)
	assert.False(t, has)

	s, _ = site.NewSite("http://127.0.0.1:9999")
	has, err = HasReadMe(s)
	assert.Error(t, err)
}

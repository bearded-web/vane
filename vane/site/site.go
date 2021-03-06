package site

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Site interface {
	String() string
	URLFor(path string) string
	Get(path string) (*http.Response, error)
	GetBody(path string) ([]byte, error)
	Head(path string) (*http.Response, error)
	Do(*http.Request) (*http.Response, error)
}

type site struct {
	*url.URL
	client *http.Client
}

func NewSite(rawurl string) (Site, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	s := &site{
		URL:    u,
		client: &http.Client{},
	}
	return s, nil
}

func (s *site) String() string {
	s.URL.Path = "/"
	return s.URL.String()
}

func (s *site) URLFor(path string) string {
	s.URL.Path = path
	return s.URL.String()
}

func (s *site) Get(path string) (*http.Response, error) {
	return s.client.Get(s.URLFor(path))
}

func (s *site) GetBody(path string) ([]byte, error) {
	resp, err := s.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (s *site) Head(path string) (*http.Response, error) {
	return s.client.Head(s.URLFor(path))
}

func (s *site) Do(req *http.Request) (*http.Response, error) {
	return s.client.Do(req)
}

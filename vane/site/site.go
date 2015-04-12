package site

import (
	"bytes"
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

	// Various state tests
	Online() bool
	HasBasicAuth() (bool, error)
	HasXMLrpc() (bool, error)
}

type site struct {
	*url.URL
	client *http.Client
}

var testXMLrpc = []byte("XML-RPC server accepts POST requests only")

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

func (s *site) Online() bool {
	_, err := s.Get("/")
	if err != nil {
		return false
	}

	return true
}

func (s *site) HasXMLrpc() (bool, error) {
	req, err := s.Get("xmlrpc.php")
	if err != nil {
		return false, err
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return false, err
	}

	return bytes.Contains(body, testXMLrpc), nil
}

func (s *site) HasBasicAuth() (bool, error) {
	req, err := s.Get("/")
	if err != nil {
		return false, err
	}

	return req.StatusCode == http.StatusUnauthorized, nil
}

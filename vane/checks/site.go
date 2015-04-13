package checks

import (
	"bytes"
	hash "crypto/md5"
	"io/ioutil"
	"net/http"

	"github.com/bearded-web/vane/vane/site"
	"github.com/bearded-web/vane/vane/utils"
)

var testXMLrpc = []byte("XML-RPC server accepts POST requests only")

// Online checks if the website is up
func Online(s site.Site) bool {
	_, err := s.Get("/")
	if err != nil {
		return false
	}

	return true
}

// HasXMLrpc checks if the website supports XML RPC
func HasXMLrpc(s site.Site) (bool, error) {
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

// HasBasicAuth checks if the website supports basic authentification
func HasBasicAuth(s site.Site) (bool, error) {
	req, err := s.Get("/")
	if err != nil {
		return false, err
	}

	return req.StatusCode == http.StatusUnauthorized, nil
}

func pageHash(s site.Site, path string) (h [hash.Size]byte, err error) {
	// ToDo: check redirections here
	req, err := s.Get(path)
	if err != nil {
		return
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	return hash.Sum(body), err
}

func HomepageHash(s site.Site) ([hash.Size]byte, error) {
	return pageHash(s, "/")
}

func Error404Hash(s site.Site) ([hash.Size]byte, error) {
	randompage := make([]byte, 64)
	utils.ReadRandomBytes(randompage)
	return pageHash(s, string(randompage))
}

package checks

import (
	"bytes"

	"github.com/bearded-web/vane/vane/site"
)

var testXMLrpc = []byte("XML-RPC server accepts POST requests only")

// HasXMLrpc checks if the website supports XML RPC
//ToDO: TESTS
func HasXMLrpc(s site.Site) (bool, error) {
	body, err := s.GetBody("xmlrpc.php")
	if err != nil {
		return false, err
	}

	return bytes.Contains(body, testXMLrpc), nil
}

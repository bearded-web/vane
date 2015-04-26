package checks

import (
	"regexp"

	"github.com/bearded-web/vane/vane/site"
)

var patternXMLRPC = regexp.MustCompile(`(?i)XML-RPC server accepts POST requests only`)

// HasXMLrpc checks if the website supports XML RPC
func HasXMLrpc(s site.Site) (bool, error) {
	body, err := s.GetBody("xmlrpc.php")
	if err != nil {
		return false, err
	}

	return patternXMLRPC.Match(body), nil
}

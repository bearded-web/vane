package checks

import (
	hash "crypto/md5"
	"regexp"

	"github.com/bearded-web/vane/vane/site"
	"github.com/bearded-web/vane/vane/utils"
)

var commentPattern = regexp.MustCompile(`(?m)<!--.*?-->`)

func pageHash(s site.Site, path string) (h [hash.Size]byte, err error) {
	body, err := s.GetBody(path)
	if err != nil {
		return
	}

	return calcHash(body), nil
}

func calcHash(body []byte) (h [hash.Size]byte) {
	return hash.Sum(clean(body))
}

func clean(body []byte) []byte {
	return commentPattern.ReplaceAll(body, []byte(""))
}

func HomepageHash(s site.Site) ([hash.Size]byte, error) {
	return pageHash(s, "/")
}

func Error404Hash(s site.Site) ([hash.Size]byte, error) {
	randompage := make([]byte, 64)
	utils.ReadRandomBytes(randompage)
	return pageHash(s, string(randompage))
}

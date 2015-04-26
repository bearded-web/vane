package checks

import (
	hash "crypto/md5"

	"github.com/bearded-web/vane/vane/site"
	"github.com/bearded-web/vane/vane/utils"
)

func pageHash(s site.Site, path string) (h [hash.Size]byte, err error) {
	// ToDo: check redirections here
	body, err := s.GetBody(path)
	if err != nil {
		return
	}

	return hash.Sum(body), nil
}

func HomepageHash(s site.Site) ([hash.Size]byte, error) {
	return pageHash(s, "/")
}

func Error404Hash(s site.Site) ([hash.Size]byte, error) {
	randompage := make([]byte, 64)
	utils.ReadRandomBytes(randompage)
	return pageHash(s, string(randompage))
}

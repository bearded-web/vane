package checks

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFilterHeaders(t *testing.T) {
	h1 := http.Header{
		"Location":   []string{"a"},
		"Connection": []string{"Close"},
	}

	result := FilterInterestingHeaders(h1)
	assert.Empty(t, result)
}

func TestFilterHeaders(t *testing.T) {
	h1 := http.Header{
		"CustomHeader": []string{"a"},
	}

	result := FilterInterestingHeaders(h1)
	assert.Equal(t, h1, result)
}

package utils

import (
	"math/rand"
)

// Inspired by https://github.com/dustin/randbo/blob/master/randbo.go
func ReadRandomBytes(output []byte) {
	todo, offset := len(output), 0
	for {
		val := int64(rand.Int63())
		for i := 0; i < 8; i++ {
			output[offset] = byte(val)
			todo--
			if todo == 0 {
				return
			}

			offset++
			val >>= 8
		}
	}
}

package utils

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

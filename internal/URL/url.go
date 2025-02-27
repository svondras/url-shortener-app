package URL

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Shorten(originalURL string) string {
	h := sha256.New()
	h.Write([]byte(originalURL))
	fmt.Println(h.Sum(nil))
	hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println(hash)
	shortURL := hash[:8]
	return shortURL
}

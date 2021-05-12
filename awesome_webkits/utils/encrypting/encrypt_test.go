package encrypting

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash := GetHashedPassword("password")
	assert.NotEmpty(t, hash)
}

func TestVerifyPassword(t *testing.T) {
	hash := GetHashedPassword("password")
	assert.True(t, CheckPassword(hash, "password"))
	assert.False(t, CheckPassword(hash, "password"))
}

func BenchmarkCalculateHash(b *testing.B) {
	s, _ := ioutil.ReadFile("/Users/iranmanesh/go/src/awesome_webkits/swagger.iso")
	b.Run("sha256", func(t *testing.B) {
		for i := 0; i < 100; i++ {
			hasher := sha256.New()
			hasher.Write(s)
			hex.EncodeToString(hasher.Sum(nil))
		}
	})
	b.Run("sha512", func(t *testing.B) {
		for i := 0; i < 100; i++ {
			hasher := sha512.New()
			hasher.Write(s)
			hex.EncodeToString(hasher.Sum(nil))
		}
	})

}

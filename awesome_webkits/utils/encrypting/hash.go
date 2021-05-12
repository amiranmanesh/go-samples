package encrypting

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"mime/multipart"
)

func GetHashedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package": "encrypting",
		}).Error("Error in encrypting...", errors.Trace(err))
		panic(errors.Trace(err))
	}

	return string(hash)

}

func CheckPassword(hashed string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}

func FileSha512(src multipart.File) (string, error) {
	hasher := sha512.New()
	if _, err := io.Copy(hasher, src); err != nil {
		return "", err
	}
	sha256Hash := hex.EncodeToString(hasher.Sum(nil))
	return sha256Hash, nil
}

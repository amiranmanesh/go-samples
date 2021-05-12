package files

import (
	"awesome_webkits/utils/encrypting"
	"github.com/sirupsen/logrus"
	"log"
	"mime/multipart"
)

type iFiles interface {
	Save(multipart.File, *multipart.FileHeader) (string, error)
}

type filesHandler struct{}

var FilesHandler iFiles = &filesHandler{}

func (filesHandler) Save(file multipart.File, multipartFileHeader *multipart.FileHeader) (string, error) {

	mime, err := GetFileMIME(file)
	if err != nil {
		return "", err
	}

	src, err := multipartFileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	sha512Hash, err := encrypting.FileSha512(src)
	if err != nil {
		log.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"Filename": multipartFileHeader.Filename,
		"MIME":     mime,
		"Size":     multipartFileHeader.Size,
		"sha512":   sha512Hash,
	}).Info("upload file log")

	return "", nil
}

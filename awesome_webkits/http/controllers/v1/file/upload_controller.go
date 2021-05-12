package file

import (
	files "awesome_webkits/file"
	"awesome_webkits/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	MB       = 1024 * 1024
	MAX_SIZE = 5 * MB // 2 Mb
)

type iUploadController interface {
	Save(ctx *gin.Context)
}

type uploadController struct{}

var UploadController iUploadController = &uploadController{}

func (uploadController) Save(ctx *gin.Context) {

	r := ctx.Request
	w := ctx.Writer

	// Limit upload size
	r.Body = http.MaxBytesReader(w, r.Body, MAX_SIZE)
	file, multipartFileHeader, err := r.FormFile("file")
	if err != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "[ERROR] 1", err.Error())
		return
	}

	_, err = files.FilesHandler.Save(file, multipartFileHeader)
	if err != nil {
		response.Json.Error(ctx, http.StatusBadRequest, "[ERROR] 2", err.Error())
		return
	}

	response.Json.Success(ctx, nil)

}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

var config = loadConfigFile()

type UploadRequest struct {
	Token string `form:"token"`
}

func bootstrap() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.Static("/files", "files")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(context *gin.Context) {
		render(200, nil, context)
	})
	router.POST("/", func(context *gin.Context) {
		var req UploadRequest
		if err := context.ShouldBind(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Token != config.Token {
			render(401, "Token is invalid", context)
			return
		}
		file, err := context.FormFile("file")
		if err != nil {
			render(400, "Please select file", context)
			return
		}
		newFileName := uuid.New().String() + filepath.Ext(file.Filename)
		if err := context.SaveUploadedFile(file, "files/"+newFileName); err != nil {
			abort(500, "Unable to save the file", context)
			return
		}
		abort(201, "File has been uploaded and available at "+getPath(newFileName, context), context)
	})
	return router
}

func main() {
	if err := bootstrap().Run(config.Bind); err != nil {
		return
	}
}

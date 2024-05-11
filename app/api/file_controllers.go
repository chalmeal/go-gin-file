package api

import (
	"errors"
	"fmt"
	"go-gin-file/app/models"
	"go-gin-file/common"
	"go-gin-file/config"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	v1 "github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

func fileUpload(c *gin.Context) {
	p := new(models.UploadFile)
	if err := c.Bind(&p); err != nil {
		log.Printf("Validation error: %s", err)
		common.Res(c, http.StatusBadRequest, err)
		return
	}

	user := getUser(p.UserId)
	if user.UserId == "" {
		common.Res(c, http.StatusBadRequest, errors.New("not found user"))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		common.Res(c, http.StatusBadRequest, err)
		return
	}

	r, err := file.Open()
	if err != nil {
		common.Res(c, http.StatusInternalServerError, err)
		return
	}

	bucket := config.GetInitKey("aws", "AWS_BUCKET_NAME")
	key := common.UuId()
	param := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key + "." + strings.Split(file.Filename, ".")[1]),
		Body:   v1.ReadSeekCloser(r),
	}
	if err := common.FileUpload(param); err != nil {
		common.Res(c, http.StatusInternalServerError, err)
		return
	}

	path := fmt.Sprintf("%s/%s/%s.%s", config.GetInitKey("aws", "AWS_ENDPOINT"), bucket, key, strings.Split(file.Filename, ".")[1])
	regParam := &models.RegFile{
		FileId: key,
		UserId: p.UserId,
		Path:   path,
	}
	if err := registerFile(regParam); err != nil {
		common.Res(c, http.StatusInternalServerError, err)
		return
	}

	common.Res(c, http.StatusOK, "Upload File Success")
}

// get file path
// Assume Blob on the front-end side when downloading files.
func getFilePath(c *gin.Context) {
	p := new(models.DownloadFile)
	if err := c.Bind(&p); err != nil {
		log.Printf("Validation error: %s", err)
		common.Res(c, http.StatusBadRequest, err)
		return
	}

	files := models.Files{}
	result := db.Where("file_id=?", p.FileId).Find(&files)
	if result.Error != nil {
		common.Res(c, http.StatusInternalServerError, result.Error)
		return
	}

	common.Res(c, http.StatusOK, &files)
}

func registerFile(p *models.RegFile) error {
	files := models.Files{
		FileId: p.FileId,
		UserId: p.UserId,
		Path:   p.Path,
	}

	result := db.Create(&files)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

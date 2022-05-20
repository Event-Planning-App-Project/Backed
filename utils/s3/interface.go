package s3

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type AwsS3 interface {
	UploadToS3(c echo.Context, filename string, src multipart.File) (string, error)
}

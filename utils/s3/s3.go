package s3

import (
	"event/config"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

func Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := c.FormValue("data")
		file, err := c.FormFile("myFile")
		if err != nil {
			fmt.Println(err, "file")
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"Code":    http.StatusForbidden,
				"Message": "Access Photo Denied",
				"data":    data,
			})
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"Code":    http.StatusForbidden,
				"Message": "Open Photo Denied",
				"data":    data,
			})
		}
		defer src.Close()
		result, err := UploadToS3(c, file.Filename, src)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"Code":    http.StatusForbidden,
				"Message": "Upload Photo Denied",
				"data":    data,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Code":    http.StatusOK,
			"Message": "Upload Photo Success",
			"Data":    result,
			"data":    data,
		})
	}
}

func ConnectAws() *session.Session {
	configuration := config.InitConfig()
	AccessKeyID := configuration.KeyIDs3
	SecretAccessKey := configuration.AccessKeyS3
	MyRegion := configuration.MyRegion

	sess, err := session.NewSession(
		&aws.Config{
			Region: &MyRegion,
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func UploadToS3(c echo.Context, filename string, src multipart.File) (string, error) {
	logger := c.Logger()
	sess := ConnectAws()

	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("eventapps"),
		Key:    aws.String(filename),
		Body:   src,
	})
	if err != nil {
		logger.Fatal(err)
		return "", err
	}

	return result.Location, nil
}

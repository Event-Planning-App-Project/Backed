package s3

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

func upload(c echo.Context) (string, error) {
	file, err := c.FormFile("myFile")
	if err != nil {
		fmt.Println(err, "file")
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer src.Close()
	result, err := UploadToS3(c, file.Filename, src)
	if err != nil {
		return "", err
	}

	return result, nil
}

func ConnectAws() *session.Session {
	// configuration := config.InitConfig()
	AccessKeyID := "AKIA3RNXISAGIKS7KVZK"
	SecretAccessKey := "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"
	MyRegion := "ap-southeast-1"

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

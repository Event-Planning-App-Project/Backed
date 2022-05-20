package s3

import (
	"event/config"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

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
		log.Warn("Error Access S3")
	}
	return sess
}

func UploadToS3(c echo.Context, filename string, src multipart.File) (string, error) {
	sess := ConnectAws()

	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("eventapps"),
		Key:    aws.String(filename),
		Body:   src,
	})
	if err != nil {
		log.Warn(err)
		return "", err
	}

	return result.Location, nil
}

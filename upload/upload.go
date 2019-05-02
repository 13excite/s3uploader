package upload

import (
	"fmt"
	"io"
	"log"
	"time"

	. "s3uploader/config"
	//	. "s3uploader/generator"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// fileNameLenght is lenght for HBKey's generator
//const fileNameLenght = 10

// TestUploadData  uploading body data to s3 storage
func TestUploadData(conf *Config, data io.ReadCloser, filename string, expireTime int) (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ru-msk"),
		Endpoint:    aws.String(conf.URL),
		Credentials: credentials.NewStaticCredentials(conf.Credentinal.AccessKey, conf.Credentinal.SecretKey, ""),
	}))
	download := s3.New(sess)
	upload := s3manager.NewUploader(sess)

	fmt.Printf("Uploading file to S3...\n")
	result, err := upload.Upload(&s3manager.UploadInput{
		Bucket: aws.String(conf.BucketName),
		Key:    aws.String(filename),
		Body:   data,
	})
	if err != nil {
		fmt.Println("Uploading failed: #", err)
		log.Fatal(err)
	}
	fmt.Printf("Successfully uploaded fileanme to %s\n", result.Location)

	//hbKey is name, which we will see in hotbox
	hbKey := filename
	req, _ := download.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(conf.BucketName),
		Key:    aws.String(hbKey),
	})
	urlStr, err := req.Presign(time.Duration(expireTime) * time.Minute)

	if err != nil {
		fmt.Println("Failed to sign request", err)
		log.Fatal(err)
	}
	return urlStr, err
}

package utils

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToS3(client *s3.Client, fileContent io.Reader, fileName string) error {
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   fileContent,
	})
	return err
}

func GetFromS3(client *s3.Client, fileName string) (*s3.GetObjectOutput, error) {
	return client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(fileName),
	})
}

func DeleteFromS3(client *s3.Client, fileName string) error {
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(fileName),
	})
	return err
}

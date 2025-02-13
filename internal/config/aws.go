package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	S3Client *s3.Client
)

func InitAWS() {
	// load aws conf.
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatal("unable to load AWS SDK config:", err)
	}

	// create s3 client
	S3Client = s3.NewFromConfig(cfg)
}

func GetS3Client() *s3.Client {
	return S3Client
}

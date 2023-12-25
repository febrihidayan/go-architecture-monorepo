package config

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/config"
)

type StorageConfig struct {
	HttpPort  string
	RpcPort   string
	Timeout   time.Duration
	MaxUpload int
	Aws       AwsConfig
}

type AwsConfig struct {
	AccessKey      string
	AccessSecret   string
	Region         types.BucketLocationConstraint
	Bucket         string
	RequestTimeout time.Duration
	Directory      string
}

func Storage() *StorageConfig {
	return &StorageConfig{
		HttpPort:  os.Getenv("HTTP_PORT"),
		RpcPort:   os.Getenv("RPC_PORT"),
		Timeout:   time.Duration(config.ConvertInt("APP_TIMEOUT")) * time.Second,
		MaxUpload: config.ConvertInt("MAX_UPLOAD_FILE"),
		Aws: AwsConfig{
			AccessKey:    os.Getenv("AWS_ACCESS_KEY"),
			AccessSecret: os.Getenv("AWS_ACCESS_SECRET"),
			Bucket:       os.Getenv("AWS_BUCKET"),
			Region:       types.BucketLocationConstraint(os.Getenv("AWS_REGION")),
			Directory:    os.Getenv("AWS_DIRECTORY"),
		},
	}
}

package main

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
)

const (
	AccessKey  = "access-key"
	SecretKey  = "secret-key"
	Endpoint   = "endpoint"
	BucketName = "bucket-name"
)

func AccessKeyFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     AccessKey,
		EnvVars:  []string{"CT_OOS_ACCESS_KEY"},
		Usage:    "天翼云 AccessKey",
		Required: true,
	}
}

func SecretKeyFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     SecretKey,
		EnvVars:  []string{"CT_OOS_SECRET_KEY"},
		Usage:    "天翼云 SecretKey",
		Required: true,
	}
}

func EndpointFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     Endpoint,
		EnvVars:  []string{"CT_OOS_ENDPOINT"},
		Usage:    "天翼云 Endpoint",
		Required: true,
	}
}

func BucketNameFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     BucketName,
		EnvVars:  []string{"CT_OOS_BUCKET_NAME"},
		Usage:    "天翼云 bucketName",
		Required: true,
	}
}

func CommonFlag() []cli.Flag {
	return []cli.Flag{
		AccessKeyFlag(),
		SecretKeyFlag(),
		EndpointFlag(),
		BucketNameFlag(),
	}
}

// NewClient create client
func NewClient(accessKey string, secretKey, endpoint string) (*oos.Client, error) {
	clientOptionV4 := oos.V4Signature(true)
	isEnableSha256 := oos.EnableSha256ForPayload(false)
	client, err := oos.New(endpoint, accessKey, secretKey, clientOptionV4, isEnableSha256)
	if err != nil {
		return nil, err
	}
	return client, nil
}

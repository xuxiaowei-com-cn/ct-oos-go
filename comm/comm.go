package comm

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
)

const (
	AccessKey  = "access-key"
	SecretKey  = "secret-key"
	Endpoint   = "endpoint"
	BucketName = "bucket-name"
	ObjectName = "object-name"
	Object     = "object"
	File       = "file"
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

func ObjectNameFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     ObjectName,
		Usage:    "上传对象路径-URI",
		Required: true,
	}
}

func ObjectFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     Object,
		Usage:    "上传对象-字符串",
		Required: true,
	}
}

func FileFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     File,
		Usage:    "上传对象-文件",
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
func NewClient(accessKey string, secretKey string, endpoint string) (*oos.Client, error) {
	clientOptionV4 := oos.V4Signature(true)
	isEnableSha256 := oos.EnableSha256ForPayload(false)
	client, err := oos.New(endpoint, accessKey, secretKey, clientOptionV4, isEnableSha256)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetBucket get the bucket
func GetBucket(accessKey string, secretKey string, endpoint string, bucketName string) (*oos.Object, error) {
	// New client
	client, err := NewClient(accessKey, secretKey, endpoint)
	if err != nil {
		return nil, err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return nil, err
	}

	return bucket, nil
}

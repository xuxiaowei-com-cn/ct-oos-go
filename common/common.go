package common

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
)

const (
	AccessKey  = "access-key"
	SecretKey  = "secret-key"
	Endpoint   = "endpoint"
	BucketName = "bucket-name"
	Uri        = "uri"
	String     = "string"
	File       = "file"
	PartSize   = "part-size"
	Routine    = "routine"
)

func AccessKeyFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     AccessKey,
		EnvVars:  []string{"CT_OOS_ACCESS_KEY"},
		Usage:    "天翼云 AccessKey",
		Required: required,
	}
}

func SecretKeyFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     SecretKey,
		EnvVars:  []string{"CT_OOS_SECRET_KEY"},
		Usage:    "天翼云 SecretKey",
		Required: required,
	}
}

func EndpointFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     Endpoint,
		EnvVars:  []string{"CT_OOS_ENDPOINT"},
		Usage:    "天翼云 Endpoint",
		Required: required,
	}
}

func BucketNameFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     BucketName,
		Aliases:  []string{"bucket"},
		EnvVars:  []string{"CT_OOS_BUCKET_NAME"},
		Usage:    "天翼云 bucketName",
		Required: required,
	}
}

func UriFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     Uri,
		Usage:    "上传路径-URI",
		Required: required,
	}
}

func StringFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     String,
		Usage:    "上传字符串",
		Required: required,
	}
}

func FileFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     File,
		Usage:    "上传文件",
		Required: required,
	}
}

func PartSizeFlag(required bool) cli.Flag {
	return &cli.IntFlag{
		Name:     PartSize,
		Usage:    "文件分片大小，单位 KB，分片数量不能超过 10000",
		Value:    100,
		Required: required,
	}
}

func RoutineFlag(required bool) cli.Flag {
	return &cli.IntFlag{
		Name:     Routine,
		Usage:    "线程",
		Value:    3,
		Required: required,
	}
}

func CommonFlag() []cli.Flag {
	return []cli.Flag{
		AccessKeyFlag(false),
		SecretKeyFlag(false),
		EndpointFlag(false),
		BucketNameFlag(false),
	}
}

func CommonFlagRequired() []cli.Flag {
	return []cli.Flag{
		AccessKeyFlag(true),
		SecretKeyFlag(true),
		EndpointFlag(true),
		BucketNameFlag(true),
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

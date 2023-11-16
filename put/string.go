package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
	"strings"
)

// PutStringCommand 上传 字符串
func PutStringCommand() *cli.Command {
	return &cli.Command{
		Name:  "string",
		Usage: "上传 字符串",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.StringFlag(true)),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var objectName = context.String(common.Uri)
			var str = context.String(common.String)

			return PutObject(accessKey, secretKey, endpoint, bucketName, objectName, str)
		},
	}
}

func PutObject(accessKey string, secretKey string, endpoint string, bucketName string, objectName string, str string) error {

	bucket, err := common.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	// Upload an object from a string
	err = bucket.PutObject(objectName, strings.NewReader(str))
	if err != nil {
		return err
	}

	log.Printf("上传字符串完成")

	return nil
}

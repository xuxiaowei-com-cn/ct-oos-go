package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
)

// PutObjectFromFileCommand 上传 文件
func PutObjectFromFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "file",
		Usage: "上传 文件",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FileFlag(true)),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var objectName = context.String(common.Uri)
			var file = context.String(common.File)

			return PutObjectFromFile(accessKey, secretKey, endpoint, bucketName, objectName, file)
		},
	}
}

func PutObjectFromFile(accessKey string, secretKey string, endpoint string, bucketName string, objectName string, file string) error {

	bucket, err := common.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	// Upload an object with local file name, user need not open the file.
	err = bucket.PutObjectFromFile(objectName, file)

	if err != nil {
		return err
	}

	log.Printf("上传文件完成")

	return nil
}

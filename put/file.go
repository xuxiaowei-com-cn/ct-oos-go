package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
	"time"
)

// PutFileCommand 上传 文件
func PutFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "file",
		Usage: "上传 文件",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FileFlag(true)),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var file = context.String(common.File)

			return PutObjectFromFile(accessKey, secretKey, endpoint, bucketName, uri, file)
		},
	}
}

func PutObjectFromFile(accessKey string, secretKey string, endpoint string, bucketName string, uri string, file string) error {

	start := time.Now()
	log.Printf("上传 文件 开始")

	bucket, err := common.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	// Upload an object with local file name, user need not open the file.
	err = bucket.PutObjectFromFile(uri, file)

	if err != nil {
		return err
	}

	log.Printf("上传 文件 结束（%s）", time.Since(start))

	return nil
}

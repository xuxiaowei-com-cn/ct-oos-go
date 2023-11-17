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
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FileFlag(true),
			common.ForceFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var file = context.String(common.File)
			var force = context.Bool(common.Force)

			log.Printf("是否开启强制上传：%t", force)

			return PutObjectFromFile(accessKey, secretKey, endpoint, bucketName, uri, file, force)
		},
	}
}

func PutObjectFromFile(accessKey string, secretKey string, endpoint string, bucketName string, uri string, file string, force bool) error {

	start := time.Now()
	log.Printf("上传 文件 开始")

	bucket, err := common.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	if !force {
		exist, err := bucket.IsObjectExist(uri)
		if err != nil {
			return err
		}

		if exist {
			log.Printf("文件 %s 已存在，跳过上传", uri)
			return nil
		}
	}

	// Upload an object with local file name, user need not open the file.
	err = bucket.PutObjectFromFile(uri, file)

	if err != nil {
		return err
	}

	log.Printf("上传 文件 结束（%s）", time.Since(start))

	return nil
}

package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
	"strings"
	"time"
)

// PutStringCommand 上传 字符串
func PutStringCommand() *cli.Command {
	return &cli.Command{
		Name:  "string",
		Usage: "上传 字符串",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.StringFlag(true),
			common.ForceFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var str = context.String(common.String)
			var force = context.Bool(common.Force)

			log.Printf("是否开启强制上传：%t", force)

			return PutObject(accessKey, secretKey, endpoint, bucketName, uri, str, force)
		},
	}
}

func PutObject(accessKey string, secretKey string, endpoint string, bucketName string, uri string, str string, force bool) error {

	start := time.Now()
	log.Printf("上传 字符串 开始")

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

	// Upload an object from a string
	err = bucket.PutObject(uri, strings.NewReader(str))
	if err != nil {
		return err
	}

	log.Printf("上传 字符串 结束（%s）", time.Since(start))

	return nil
}

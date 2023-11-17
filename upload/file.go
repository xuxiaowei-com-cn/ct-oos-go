package upload

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
	"time"
)

func UploadFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "file",
		Usage: "上传 文件-分片",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FileFlag(true),
			common.PartSizeFlag(), common.RoutineFlag(), common.ForceFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var file = context.String(common.File)
			var force = context.Bool(common.Force)
			var partSize = context.Int64(common.PartSize)
			var routine = context.Int(common.Routine)

			return UploadFile(accessKey, secretKey, endpoint, bucketName, uri, file, partSize, routine, force)
		},
	}
}

func UploadFile(accessKey string, secretKey string, endpoint string, bucketName string, uri string, file string,
	partSize int64, routine int, force bool) error {

	start := time.Now()
	log.Printf("分片上传 开始")

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

	err = bucket.UploadFile(uri, file, partSize*1024, oos.Routines(routine))
	if err != nil {
		return err
	}

	log.Printf("分片上传 结束（%s）", time.Since(start))

	return nil
}

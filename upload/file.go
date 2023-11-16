package upload

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
)

func UploadFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "file",
		Usage: "上传 文件-分片",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FileFlag(true),
			common.PartSizeFlag(false), common.RoutineFlag(false)),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var file = context.String(common.File)
			var partSize = context.Int64(common.PartSize)
			var routine = context.Int(common.Routine)

			return UploadFile(accessKey, secretKey, endpoint, bucketName, uri, file, partSize, routine)
		},
	}
}

func UploadFile(accessKey string, secretKey string, endpoint string, bucketName string, uri string, file string,
	partSize int64, routine int) error {

	bucket, err := common.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	err = bucket.UploadFile(uri, file, partSize*1024, oos.Routines(routine))
	if err != nil {
		return err
	}

	log.Printf("分片上传文件完成")

	return nil
}

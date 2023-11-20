package upload

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
	"os"
	"path/filepath"
	"time"
)

// UploadFolderCommand 上传 文件夹
func UploadFolderCommand() *cli.Command {
	return &cli.Command{
		Name:  "folder",
		Usage: "上传 文件夹-分片",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FolderFlag(true),
			common.ForceFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var folder = context.String(common.Folder)
			var force = context.Bool(common.Force)
			var partSize = context.Int64(common.PartSize)
			var routine = context.Int(common.Routine)

			log.Printf("是否开启强制上传：%t", force)

			fileInfo, err := os.Stat(folder)
			if err != nil {
				return err
			}

			if fileInfo.IsDir() {
				return UploadFolder(accessKey, secretKey, endpoint, bucketName, uri, folder, partSize, routine, force)
			} else {
				return errors.New(fmt.Sprintf("路径 %s 不是一个文件夹", folder))
			}
		},
	}
}

func UploadFolder(accessKey string, secretKey string, endpoint string, bucketName string, uri string, folder string, partSize int64, routine int, force bool) error {

	start := time.Now()
	log.Printf("上传 文件夹 开始")

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			// 跳过文件夹
			return nil
		}

		fileName := filepath.Base(path)

		if fileName == "desktop.ini" {
			return nil
		}

		file := path[len(folder)+1:]

		log.Printf("上传 文件 %s 开始", path)

		bucket, err := common.GetBucket(accessKey, secretKey, endpoint, bucketName)
		if err != nil {
			return err
		}

		objectKey := uri + "/" + file

		if !force {
			exist, err := bucket.IsObjectExist(objectKey)
			if err != nil {
				return err
			}

			if exist {
				log.Printf("文件 %s 已存在，跳过上传", objectKey)
				return nil
			}
		}

		err = bucket.UploadFile(uri, file, partSize*1024, oos.Routines(routine))

		if err != nil {
			return err
		}

		log.Printf("上传 文件 %s 结束", path)

		return nil
	})

	if err != nil {
		return err
	}

	log.Printf("分片上传 结束（%s）", time.Since(start))

	return nil
}
package put

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// PutFolderCommand 上传 文件夹
func PutFolderCommand() *cli.Command {
	return &cli.Command{
		Name:  "folder",
		Usage: "上传 文件夹",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FolderFlag(true),
			common.ForceFlag(), common.EnableLogFlag(), common.LogNameFlag(), common.LogFolderFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var folder = context.String(common.Folder)
			var force = context.Bool(common.Force)
			var enableLog = context.Bool(common.EnableLog)
			var logName = context.String(common.LogName)
			var logFolder = context.String(common.LogFolder)

			if enableLog {

				file, err := common.LogConfig(logName, logFolder)
				if err != nil {
					return err
				}

				defer func(file *os.File) {
					err := file.Close()
					if err != nil {
						log.Fatal(err)
					}
				}(file)

				// 设置日志输出到控制台和日志文件
				multi := io.MultiWriter(os.Stdout, file)
				log.SetOutput(multi)

				// 设置日志输出位置为文件
				// log.SetOutput(file)
			}

			log.Printf("是否开启强制上传：%t", force)

			fileInfo, err := os.Stat(folder)
			if err != nil {
				return err
			}

			if fileInfo.IsDir() {
				return PutObjectFromFolder(accessKey, secretKey, endpoint, bucketName, uri, folder, force)
			} else {
				return errors.New(fmt.Sprintf("路径 %s 不是一个文件夹", folder))
			}
		},
	}
}

func PutObjectFromFolder(accessKey string, secretKey string, endpoint string, bucketName string, uri string, folder string, force bool) error {

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

		// Upload an object with local file name, user need not open the file.
		err = bucket.PutObjectFromFile(objectKey, path)

		if err != nil {
			return err
		}

		log.Printf("上传 文件 %s 结束", path)

		return nil
	})

	if err != nil {
		return err
	}

	log.Printf("上传 文件夹 结束（%s）", time.Since(start))

	return nil
}

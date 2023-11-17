package put

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	accessKey  string
	secretKey  string
	endpoint   string
	bucketName string
	uri        string
	folder     string
)

// PutFolderCommand 上传 文件夹
func PutFolderCommand() *cli.Command {
	return &cli.Command{
		Name:  "folder",
		Usage: "上传 文件夹",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FolderFlag(true)),
		Action: func(context *cli.Context) error {
			accessKey = context.String(common.AccessKey)
			secretKey = context.String(common.SecretKey)
			endpoint = context.String(common.Endpoint)
			bucketName = context.String(common.BucketName)
			uri = context.String(common.Uri)
			folder = context.String(common.Folder)

			fileInfo, err := os.Stat(folder)
			if err != nil {
				return err
			}

			if fileInfo.IsDir() {
				return PutObjectFromFolder(folder)
			} else {
				return errors.New(fmt.Sprintf("路径 %s 不是一个文件夹", folder))
			}
		},
	}
}

func PutObjectFromFolder(folder string) error {

	start := time.Now()
	log.Printf("上传 文件夹 开始")

	err := filepath.WalkDir(folder, VisitFile)
	if err != nil {
		return err
	}

	log.Printf("上传 文件夹 结束（%s）", time.Since(start))

	return nil
}

func VisitFile(path string, info os.DirEntry, err error) error {
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

	// Upload an object with local file name, user need not open the file.
	err = bucket.PutObjectFromFile(objectKey, path)

	if err != nil {
		return err
	}

	log.Printf("上传 文件 %s 结束", path)

	return nil
}

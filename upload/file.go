package upload

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"io"
	"log"
	"os"
	"time"
)

func UploadFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "file",
		Usage: "上传 文件-分片",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.FileFlag(true),
			common.PartSizeFlag(), common.RoutineFlag(), common.ForceFlag(), common.EnableLogFlag(),
			common.LogNameFlag(), common.LogFolderFlag(), common.ConnectTimeoutSecFlag(),
			common.ReadWriteTimeoutSecFlag(), common.MicrosecondsFlag(), common.LongFileFlag()),
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
			var enableLog = context.Bool(common.EnableLog)
			var logName = context.String(common.LogName)
			var logFolder = context.String(common.LogFolder)
			var connectTimeoutSec = context.Int64(common.ConnectTimeoutSec)
			var readWriteTimeoutSec = context.Int64(common.ReadWriteTimeoutSec)
			var microseconds = context.Bool(common.Microseconds)
			var longFile = context.Bool(common.LongFile)

			flag := log.Ldate | log.Ltime

			if microseconds {
				flag = flag | log.Lmicroseconds
			}

			if longFile {
				flag = flag | log.Llongfile
			}

			log.SetFlags(flag)

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

			err := CheckPartSize(partSize)
			if err != nil {
				return err
			}

			log.Printf("是否开启强制上传：%t", force)

			return UploadFile(accessKey, secretKey, endpoint, bucketName, uri, file, partSize, routine, force,
				connectTimeoutSec, readWriteTimeoutSec)
		},
	}
}

func UploadFile(accessKey string, secretKey string, endpoint string, bucketName string, uri string, file string,
	partSize int64, routine int, force bool, connectTimeoutSec int64, readWriteTimeoutSec int64) error {

	start := time.Now()
	log.Printf("分片上传 开始")

	bucket, err := common.GetBucketWithTimeOut(accessKey, secretKey, endpoint, bucketName, connectTimeoutSec, readWriteTimeoutSec)
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

	err = bucket.UploadFile(uri, file, partSize*1024*1024, oos.Routines(routine))
	if err != nil {
		return err
	}

	log.Printf("分片上传 结束（%s）", time.Since(start))

	return nil
}

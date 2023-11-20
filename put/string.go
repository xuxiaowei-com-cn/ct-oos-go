package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// PutStringCommand 上传 字符串
func PutStringCommand() *cli.Command {
	return &cli.Command{
		Name:  "string",
		Usage: "上传 字符串",
		Flags: append(common.CommonFlagRequired(), common.UriFlag(true), common.StringFlag(true),
			common.ForceFlag(), common.EnableLogFlag(), common.LogNameFlag(), common.LogFolderFlag(),
			common.ConnectTimeoutSecFlag(), common.ReadWriteTimeoutSecFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)
			var uri = context.String(common.Uri)
			var str = context.String(common.String)
			var force = context.Bool(common.Force)
			var enableLog = context.Bool(common.EnableLog)
			var logName = context.String(common.LogName)
			var logFolder = context.String(common.LogFolder)
			var connectTimeoutSec = context.Int64(common.ConnectTimeoutSec)
			var readWriteTimeoutSec = context.Int64(common.ReadWriteTimeoutSec)

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

			return PutObject(accessKey, secretKey, endpoint, bucketName, uri, str, force,
				connectTimeoutSec, readWriteTimeoutSec)
		},
	}
}

func PutObject(accessKey string, secretKey string, endpoint string, bucketName string, uri string, str string, force bool,
	connectTimeoutSec int64, readWriteTimeoutSec int64) error {

	start := time.Now()
	log.Printf("上传 字符串 开始")

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

	// Upload an object from a string
	err = bucket.PutObject(uri, strings.NewReader(str))
	if err != nil {
		return err
	}

	log.Printf("上传 字符串 结束（%s）", time.Since(start))

	return nil
}

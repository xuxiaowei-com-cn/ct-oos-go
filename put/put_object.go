package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/comm"
	"log"
	"strings"
)

func PutObjectCommand() *cli.Command {
	return &cli.Command{
		Name:  "put-object",
		Usage: "上传 对象（字符串）",
		Flags: append(comm.CommonFlag(), comm.ObjectNameFlag(), comm.ObjectFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(comm.AccessKey)
			var secretKey = context.String(comm.SecretKey)
			var endpoint = context.String(comm.Endpoint)
			var bucketName = context.String(comm.BucketName)
			var objectName = context.String(comm.ObjectName)
			var object = context.String(comm.Object)

			return PutObject(accessKey, secretKey, endpoint, bucketName, objectName, object)
		},
	}
}

func PutObjectFromFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "put-object-from-file",
		Usage: "上传 对象（文件）",
		Flags: append(comm.CommonFlag(), comm.ObjectNameFlag(), comm.FileFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(comm.AccessKey)
			var secretKey = context.String(comm.SecretKey)
			var endpoint = context.String(comm.Endpoint)
			var bucketName = context.String(comm.BucketName)
			var objectName = context.String(comm.ObjectName)
			var file = context.String(comm.File)

			return PutObjectFromFile(accessKey, secretKey, endpoint, bucketName, objectName, file)
		},
	}
}

func PutObject(accessKey string, secretKey string, endpoint string, bucketName string, objectName string, object string) error {

	bucket, err := comm.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	// Upload an object from a string
	err = bucket.PutObject(objectName, strings.NewReader(object))
	if err != nil {
		return err
	}

	log.Printf("上传字符串完成")

	return nil
}

func PutObjectFromFile(accessKey string, secretKey string, endpoint string, bucketName string, objectName string, file string) error {

	bucket, err := comm.GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	// Upload an object with local file name, user need not open the file.
	err = bucket.PutObjectFromFile(objectName, file)

	if err != nil {
		return err
	}

	log.Printf("上传文件完成")

	return nil
}

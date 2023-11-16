package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

func PutObjectCommand() *cli.Command {
	return &cli.Command{
		Name:  "put-object",
		Usage: "上传 对象（字符串）",
		Flags: append(CommonFlag(), ObjectNameFlag(), ObjectFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(AccessKey)
			var secretKey = context.String(SecretKey)
			var endpoint = context.String(Endpoint)
			var bucketName = context.String(BucketName)
			var objectName = context.String(ObjectName)
			var object = context.String(Object)

			return PutObject(accessKey, secretKey, endpoint, bucketName, objectName, object)
		},
	}
}

func PutObjectFromFileCommand() *cli.Command {
	return &cli.Command{
		Name:  "put-object-from-file",
		Usage: "上传 对象（文件）",
		Flags: append(CommonFlag(), ObjectNameFlag(), FileFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(AccessKey)
			var secretKey = context.String(SecretKey)
			var endpoint = context.String(Endpoint)
			var bucketName = context.String(BucketName)
			var objectName = context.String(ObjectName)
			var file = context.String(File)

			return PutObjectFromFile(accessKey, secretKey, endpoint, bucketName, objectName, file)
		},
	}
}

func PutObject(accessKey string, secretKey string, endpoint string, bucketName string, objectName string, object string) error {

	bucket, err := GetBucket(accessKey, secretKey, endpoint, bucketName)
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

	bucket, err := GetBucket(accessKey, secretKey, endpoint, bucketName)
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

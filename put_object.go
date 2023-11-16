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
			var object = context.String(ObjectName)

			return PutObject(accessKey, secretKey, endpoint, bucketName, objectName, object)
		},
	}
}

func PutObject(accessKey string, secretKey string, endpoint string, bucketName string, objectName string, object string) error {

	bucket, err := GetBucket(accessKey, secretKey, endpoint, bucketName)
	if err != nil {
		return err
	}

	// Case 1: Upload an object from a string
	err = bucket.PutObject(objectName, strings.NewReader(object))
	if err != nil {
		return err
	}

	log.Printf("上传完成")

	return nil
}

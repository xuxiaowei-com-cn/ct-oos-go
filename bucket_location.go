package main

import (
	"github.com/urfave/cli/v2"
	"log"
)

func GetBucketLocationCommand() *cli.Command {
	return &cli.Command{
		Name:  "get-bucket-location",
		Usage: "获取 Bucket Location",
		Flags: append(CommonFlag()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(AccessKey)
			var secretKey = context.String(SecretKey)
			var endpoint = context.String(Endpoint)
			var bucketName = context.String(BucketName)

			return GetBucketLocation(accessKey, secretKey, endpoint, bucketName)
		},
	}
}

func GetBucketLocation(accessKey string, secretKey string, endpoint string, bucketName string) error {
	// New client
	client, err := NewClient(accessKey, secretKey, endpoint)
	if err != nil {
		return err
	}

	ret, err := client.GetBucketLocation(bucketName)
	if err != nil {
		return err
	}

	log.Printf("XMLName.Space: %s\n", ret.XMLName.Space)
	log.Printf("XMLName.Local: %s\n", ret.XMLName.Local)
	log.Printf("MetaLocation: %s\n", ret.MetaLocation)
	log.Printf("DataLocationType: %s\n", ret.DataLocationType)
	log.Printf("DataLocationList: %s\n", ret.DataLocationList)
	log.Printf("ScheduleStrategy: %s\n", ret.ScheduleStrategy)

	return nil
}

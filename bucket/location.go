package bucket

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
	"log"
)

func GetLocationCommand() *cli.Command {
	return &cli.Command{
		Name:  "get-location",
		Usage: "获取 location bucket",
		Flags: append(common.CommonFlagRequired()),
		Action: func(context *cli.Context) error {
			var accessKey = context.String(common.AccessKey)
			var secretKey = context.String(common.SecretKey)
			var endpoint = context.String(common.Endpoint)
			var bucketName = context.String(common.BucketName)

			return GetLocation(accessKey, secretKey, endpoint, bucketName)
		},
	}
}

func GetLocation(accessKey string, secretKey string, endpoint string, bucketName string) error {
	// New client
	client, err := common.NewClient(accessKey, secretKey, endpoint)
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
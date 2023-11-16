package bucket

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
)

func GetBucketCommand() *cli.Command {
	return &cli.Command{
		Name:  "bucket",
		Usage: "桶",
		Flags: common.CommonFlag(),
		Subcommands: []*cli.Command{
			GetLocationCommand(),
		},
	}
}

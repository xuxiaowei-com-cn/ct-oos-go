package put

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
)

// PutCommand 上传对象
func PutCommand() *cli.Command {
	return &cli.Command{
		Name:  "put",
		Usage: "上传对象",
		Flags: append(common.CommonFlag(), common.UriFlag(false), common.StringFlag(false), common.FileFlag(false)),
		Subcommands: []*cli.Command{
			PutStringCommand(),
			PutObjectFromFileCommand(),
		},
	}
}
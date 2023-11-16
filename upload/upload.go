package upload

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
)

// UploadCommand 上传文件-分片
func UploadCommand() *cli.Command {
	return &cli.Command{
		Name:  "upload",
		Usage: "上传文件-分片",
		Flags: append(common.CommonFlag(), common.UriFlag(false), common.FileFlag(false),
			common.PartSizeFlag(false), common.RoutineFlag(false)),
		Subcommands: []*cli.Command{
			UploadFileCommand(),
		},
	}
}

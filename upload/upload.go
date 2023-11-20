package upload

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go/common"
)

// UploadCommand 上传文件-分片
func UploadCommand() *cli.Command {
	return &cli.Command{
		Name:  "upload",
		Usage: "上传文件-分片",
		Flags: append(common.CommonFlag(), common.UriFlag(false), common.FileFlag(false),
			common.FolderFlag(false), common.PartSizeFlag(), common.RoutineFlag(), common.ForceFlag(),
			common.EnableLogFlag(), common.LogNameFlag(), common.LogFolderFlag()),
		Subcommands: []*cli.Command{
			UploadFileCommand(),
			UploadFolderCommand(),
		},
	}
}

func CheckPartSize(partSize int64) error {
	if partSize < common.MinPartSize {
		return errors.New(fmt.Sprintf("分片最小值是 %d M", common.MinPartSize))
	}
	return nil
}

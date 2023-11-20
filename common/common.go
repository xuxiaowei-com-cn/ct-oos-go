package common

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/ct-oos-go-sdk/oos"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	AccessKey                  = "access-key"
	SecretKey                  = "secret-key"
	Endpoint                   = "endpoint"
	BucketName                 = "bucket-name"
	Uri                        = "uri"
	String                     = "string"
	File                       = "file"
	Folder                     = "folder"
	Force                      = "force"
	PartSize                   = "part-size"
	DefaultPartSize            = 5 // 默认分片大小，单位 M
	MinPartSize                = 5 // 最小分片大小，单位 M
	Routine                    = "routine"
	EnableLog                  = "enable-log"
	LogName                    = "log-name"
	LogFolder                  = "log-folder"
	DefaultLogFolder           = ".ct-oos-go"
	DefaultLogName             = "ct-oos-go"           // 日志名称-前缀
	ConnectTimeoutSec          = "connect-timeout-sec" // 连接超时时间
	DefaultConnectTimeoutSec   = 3
	ReadWriteTimeoutSec        = "read-write-timeout-sec" // 读写超时时间
	DefaultReadWriteTimeoutSec = 3
	Microseconds               = "microseconds"
	LongFile                   = "long-file"
)

func AccessKeyFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     AccessKey,
		EnvVars:  []string{"CT_OOS_ACCESS_KEY"},
		Usage:    "天翼云 AccessKey",
		Required: required,
	}
}

func SecretKeyFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     SecretKey,
		EnvVars:  []string{"CT_OOS_SECRET_KEY"},
		Usage:    "天翼云 SecretKey",
		Required: required,
	}
}

func EndpointFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     Endpoint,
		EnvVars:  []string{"CT_OOS_ENDPOINT"},
		Usage:    "天翼云 Endpoint",
		Required: required,
	}
}

func BucketNameFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     BucketName,
		Aliases:  []string{"bucket"},
		EnvVars:  []string{"CT_OOS_BUCKET_NAME"},
		Usage:    "天翼云 bucketName",
		Required: required,
	}
}

func UriFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     Uri,
		Usage:    "上传路径-URI",
		Required: required,
	}
}

func StringFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     String,
		Usage:    "上传字符串",
		Required: required,
	}
}

func FileFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     File,
		Usage:    "上传文件",
		Required: required,
	}
}

func FolderFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     Folder,
		Usage:    "上传文件夹",
		Required: required,
	}
}

func ForceFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  Force,
		Usage: "是否强制上传，会覆盖文件",
		Value: false,
	}
}

func PartSizeFlag() cli.Flag {
	return &cli.IntFlag{
		Name:  PartSize,
		Usage: fmt.Sprintf("文件分片大小，单位 M，最小值 %d M，分片数量不能超过 10000", MinPartSize),
		Value: DefaultPartSize,
	}
}

func RoutineFlag() cli.Flag {
	return &cli.IntFlag{
		Name:  Routine,
		Usage: "线程",
		Value: 3,
	}
}

func EnableLogFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  EnableLog,
		Usage: "开启日志",
		Value: false,
	}
}

func LogNameFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  LogName,
		Usage: "日志名称-前缀",
		Value: DefaultLogName,
	}
}

func LogFolderFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  LogFolder,
		Usage: fmt.Sprintf("日志文件夹，默认是当前用户主目录下的 %s 文件夹", DefaultLogFolder),
	}
}

func MicrosecondsFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  Microseconds,
		Usage: "日志打印时间精确到微秒",
		Value: false,
	}
}

func LongFileFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:  LongFile,
		Usage: "日志打印使用长包名",
		Value: false,
	}
}

func ConnectTimeoutSecFlag() cli.Flag {
	return &cli.Int64Flag{
		Name:  ConnectTimeoutSec,
		Usage: "连接超时时间，单位是 s",
		Value: DefaultConnectTimeoutSec,
	}
}

func ReadWriteTimeoutSecFlag() cli.Flag {
	return &cli.Int64Flag{
		Name:  ReadWriteTimeoutSec,
		Usage: "读写超时时间，单位是 s",
		Value: DefaultReadWriteTimeoutSec,
	}
}

func LogConfig(name string, logFolder string) (*os.File, error) {

	// 获取当前用户信息
	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("获取当前用于异常：%s\n", err)
	}

	// 获取当前用户的主目录
	homeDir := currentUser.HomeDir

	if logFolder == "" {
		logFolder = filepath.Join(homeDir, DefaultLogFolder)
	}

	// 使用os.Stat判断文件夹是否存在
	_, err = os.Stat(logFolder)

	// 如果文件夹不存在，则创建
	if os.IsNotExist(err) {
		err := os.MkdirAll(logFolder, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("创建日志文件夹异常：%s\n", err)
		}
	} else if err != nil {
		// 如果发生其他错误，返回错误信息
		return nil, fmt.Errorf("检查日志文件夹异常：%s\n", err)
	}

	// 获取当前时间
	currentTime := time.Now()

	// 格式化为日期字符串
	dateString := currentTime.Format("2006-01-02_15-04-05")

	logFile := filepath.Join(logFolder, name+"-"+dateString+".log")

	// 打开或创建一个日志文件
	file, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err)
	}

	return file, nil
}

func CommonFlag() []cli.Flag {
	return []cli.Flag{
		AccessKeyFlag(false),
		SecretKeyFlag(false),
		EndpointFlag(false),
		BucketNameFlag(false),
	}
}

func CommonFlagRequired() []cli.Flag {
	return []cli.Flag{
		AccessKeyFlag(true),
		SecretKeyFlag(true),
		EndpointFlag(true),
		BucketNameFlag(true),
	}
}

// NewClient create client
func NewClient(accessKey string, secretKey string, endpoint string) (*oos.Client, error) {
	clientOptionV4 := oos.V4Signature(true)
	isEnableSha256 := oos.EnableSha256ForPayload(false)
	client, err := oos.New(endpoint, accessKey, secretKey, clientOptionV4, isEnableSha256)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewClientWithTimeOut create client
func NewClientWithTimeOut(accessKey string, secretKey string, endpoint string, connectTimeoutSec int64, readWriteTimeoutSec int64) (*oos.Client, error) {
	clientOptionV4 := oos.V4Signature(true)
	isEnableSha256 := oos.EnableSha256ForPayload(false)
	timeOut := oos.Timeout(connectTimeoutSec, readWriteTimeoutSec)
	client, err := oos.New(endpoint, accessKey, secretKey, clientOptionV4, isEnableSha256, timeOut)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetBucket get the bucket
func GetBucket(accessKey string, secretKey string, endpoint string, bucketName string) (*oos.Object, error) {
	// New client
	client, err := NewClient(accessKey, secretKey, endpoint)
	if err != nil {
		return nil, err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// GetBucketWithTimeOut get the bucket
func GetBucketWithTimeOut(accessKey string, secretKey string, endpoint string, bucketName string, connectTimeoutSec int64, readWriteTimeoutSec int64) (*oos.Object, error) {
	// New client
	client, err := NewClientWithTimeOut(accessKey, secretKey, endpoint, connectTimeoutSec, readWriteTimeoutSec)
	if err != nil {
		return nil, err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return nil, err
	}

	return bucket, nil
}

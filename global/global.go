package global

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lw000/gocommon/ip2region"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
	"path"
	"reportGameErr/config"
	SourceMap "reportGameErr/sourcemap"
	"time"
)

var (
	// ProjectConfig 工程配置
	ProjectConfig *config.IniConfig
	// IpServer IP地址转换
	IpServer *tyip2region.IpRegionServer
	// SourceMapServer SourceMap文件解析对象
	SourceMapServer *SourceMap.SourceMapManager
)

func init() {
	IpServer = tyip2region.NewIpRegionServer()
	SourceMapServer = SourceMap.NewSourceMapManager()
}

// config logrus log to local filesystem, with file rotation
func configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d_%H%M",
		// rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)

	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	// 打印函数，和函数所在行号
	// log.SetReportCaller(true)

	log.AddHook(lfHook)

	// 配置mongodb
	//mgoHooker, err := mgorus.NewHooker("root:root@192.168.110.110:27017", "h5_error_log", "collection")
	//if err == nil {
	//	log.AddHook(mgoHooker)
	//} else {
	//	log.Error(err)
	//}

	mgoHooker, err := mgorus.NewHookerWithAuth("192.168.110.110:27017", "h5_error_log", "collection", "root", "root")
	if err == nil {
		log.AddHook(mgoHooker)
	} else {
		log.Error(err)
	}
}

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig() error {
	var err error
	ProjectConfig, err = config.LoadIniConfig("conf/conf.ini")
	if err != nil {
		log.Error(err)
		return err
	}

	// 日志分割 1按天分割，2按周分割, 3 按月分割，4按年分割
	var logName = "h5_error"
	switch ProjectConfig.SplitLog {
	case 1:
		configLocalFilesystemLogger("log", logName, time.Hour*24*365, time.Hour*24)
	case 2:
		configLocalFilesystemLogger("log", logName, time.Hour*24*365, time.Hour*24*7)
	case 3:
		configLocalFilesystemLogger("log", logName, time.Hour*24*365, time.Hour*24*30)
	case 4:
		configLocalFilesystemLogger("log", logName, time.Hour*24*365, time.Hour*24*365)
	default:
		configLocalFilesystemLogger("log", logName, time.Hour*24*365, time.Hour*24)
	}

	return nil
}

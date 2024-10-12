package pkg

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = &logrus.Logger{
	Out: os.Stdout,
	//ReportCaller: true, //打印日志代码所在行
	Formatter: &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	},
	Hooks: make(logrus.LevelHooks),
	Level: logLevel,
}

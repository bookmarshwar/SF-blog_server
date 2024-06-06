package core

import (
	"blog_server/global"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", global.CONFIG.Logger.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s \n", global.CONFIG.Logger.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
func InitLogger() *logrus.Logger {
	mlog := logrus.New()                                 //新建实例
	mlog.SetOutput(os.Stdout)                            //设置输出类型
	mlog.SetReportCaller(global.CONFIG.Logger.Show_line) //开启返回函数名和行号
	mlog.SetFormatter(&LogFormatter{})                   //使用自定义foramtter
	level, err := logrus.ParseLevel(global.CONFIG.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level) //设置最低level
	return mlog
}
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)                            //设置输出类型
	logrus.SetReportCaller(global.CONFIG.Logger.Show_line) //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                   //使用自定义foramtter
	level, err := logrus.ParseLevel(global.CONFIG.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低level
}

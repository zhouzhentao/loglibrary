/*
	改善能直接调用并且可以自由选择终端或者文件或者网络方式format输出日志
	设计思路：
		易用性封装
		默认情况下执行直接输出到终端，如果读取到配置文件里面有logpath的情况下 输出到logpath，如果有kafka配置就输出到kakfa
		输出文件名行号
		文件输出不同 errorlog 输出errorlog文件 常规输出常规
*/

package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var Log Logprint

func InitLoger(name string) {
	switch name {
	case "file":
		f, err := os.OpenFile("./a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 655)
		if err != nil {
			return
		}
		Log = Newfilelog(f, "./a.log")
	case "console":
		Log = Newconsolelog()
	default:
		Log = Newconsolelog()
	}
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

//  2020/09/21 11:18:01.521895

/*
	优化 异步写入 防止文件锁死占用
	1、当业务打日志的时候调用方法，我们把需要写入日志的相关数据写入到chan（队列）
	2、然后我们在后台起一个线程不断从chan里面获取日志，最终写入文件中
*/
func Writefilelog(format string, Logdatachan chan *string, arg ...interface{}) {
	msg := fmt.Sprintf(format, arg...)
	//_, err := fmt.Fprintf(file, format, arg...)
	Logdatachan <- &msg
}

func Gettime() string {
	now := time.Now()
	timeformart := now.Format("2006/01/02 15:04:05.000")
	//timeformart := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%d", now.Year(), now.Month(), now.Day(), now.Hour(),
	//	now.Minute(),
	//	now.Second(), now.Nanosecond()/1000)
	return timeformart
}

func Logformat(level string) string {
	filename, line := GetLanno()
	logformat := fmt.Sprintf("%s [%s] %s:%d", Gettime(), level, filename, line)
	return logformat + " %s"
}

func GetLanno() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
}

type Logprint interface {
	Debug(arg ...interface{}) error
	Info(arg ...interface{}) error
	Warn(arg ...interface{}) error
	Error(arg ...interface{}) error
	Fatal(arg ...interface{}) error
}

package logger

import (
	"fmt"
	"log"
	"os"
)

var loger *log.Logger

func InitLogger() {
	//创建输出日志文件
	//logFile, err := os.Create("log/" + time.Now().Format("20060102") + ".log")
	logFile, err := os.Create("log/debug.log")
	if err != nil {
		fmt.Println(err)
	}
	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger = log.New(logFile, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	//Flags返回Logger的输出选项

	//SetFlags设置输出选项
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func SetLogger(id int, msg string) {

	logList := []string{"Info", "Warning", "Error"}
	loger.Output(2, logList[id]+": "+msg)
}

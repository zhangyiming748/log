package log

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger // 仅打印到屏幕
	Debug *log.Logger // 打印屏幕并保存到文件
	Warn  *log.Logger // 出现一般性错误,打印屏幕并保存到文件
	TTY   *log.Logger
)

func init() {
	file := "program.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("打开日志文件错误")
	}
	Info = log.New(os.Stdout, "INFO:", log.Ltime)
	Debug = log.New(io.MultiWriter(logf, os.Stdout), "DEBUG:", log.Lshortfile)
	Warn = log.New(io.MultiWriter(logf, os.Stdout), "WARN:", log.Ltime|log.Llongfile)
	TTY = log.New(os.Stdout, "TTY:", 0)
}

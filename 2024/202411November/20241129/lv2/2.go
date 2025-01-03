package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Path := flag.String("Path", "test3.txt", "输入生成日志的路径：")
	Mode := flag.String("Mode", "ALL", "请输入日志的模式(ALL,ONLY_TIME,PROCESS):")
	TFormat := flag.String("TFormat", "2006年01月02日---15点04分05秒", "时间格式(没有则默认)")

	flag.Parse()
	// 打开一个日志文件，如果文件不存在则创建，追加写入
	file, err := os.OpenFile(*Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	// 创建一个带时间戳的写入器
	logWriter := newTimestampWriter(file, *Mode, *TFormat)

	// 模拟用户操作并记录日志
	fmt.Fprintln(logWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作B")
}

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间和时间戳
type timestampWriter struct {
	logFile io.Writer
	Mode    string
	TFormat string
}

// 传入一个io.writer,file实现了io.writer接口
func newTimestampWriter(w io.Writer, mode string, TFormat string) *timestampWriter {
	return &timestampWriter{logFile: w, Mode: mode, TFormat: TFormat}
}

func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 添加时间戳和时间
	t1 := time.Now().Unix()
	t2 := time.Now()
	var Combine string
	if tw.Mode == "ALL" {
		Combine = fmt.Sprintf("时间：%s，unix时间戳为%d,信息：%s\n", t2.Format(tw.TFormat), t1, p)
	} else if tw.Mode == "ONLY_TIME" {
		Combine = fmt.Sprintf("时间：%s，unix时间戳为%d\n", t2.Format(tw.TFormat), t1)
	} else if tw.Mode == "PROCESS" {
		Combine = fmt.Sprintf("信息：%s\n", p)
	}

	// 输出到文件
	Writers := io.MultiWriter(tw.logFile, os.Stdout)

	return Writers.Write([]byte(Combine))
}

package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CopyFile(SourceP string, TargetP string) error {
	PFile1, err := os.OpenFile(SourceP, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("打开失败目录2,", err)
	}
	defer PFile1.Close()
	PFile2, err2 := os.OpenFile(TargetP, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err2 != nil {
		log.Fatal("打开失败目录2,", err2)
	}

	defer PFile2.Close()
	_, err3 := io.Copy(PFile2, PFile1)

	return err3
}

func Check(SourceP string, TargetP string, fileModTimes map[string]time.Time) {
	PFile1, err := os.ReadDir(SourceP)
	if err != nil {
		log.Fatal("同步失败,", err)
	}

	for _, file := range PFile1 {
		File1name := filepath.Join(SourceP, file.Name())
		File2name := filepath.Join(TargetP, file.Name())

		FileInfo, err := file.Info()
		if err != nil {
			log.Fatal("获取更新时间失败,", err)
		}
		modTime := FileInfo.ModTime()

		if modTime.After(fileModTimes[file.Name()]) {
			err := CopyFile(File1name, File2name)
			if err != nil {
				log.Fatal("同步失败：", err)
			}
			fileModTimes[file.Name()] = modTime
		}

	}

}

func main() {
	//SorPath := flag.String("OJPath", "", "请输入要监视的文件路径")
	//TarPath := flag.String("SYPath", "", "请输入要监视的文件路径")
	SourPath := "D:\\New\\1\\1\\"
	TarPath := "D:\\New\\1\\2\\"
	//flag.Parse()

	fileModTimes := make(map[string]time.Time)
	P1, err := os.ReadDir(SourPath)
	if err != nil {
		log.Fatal("同步失败,", err)
	}
	for _, file := range P1 {
		fileName := file.Name()
		fileModTimes[fileName] = time.Now()
	}
	Ticker := time.NewTicker(time.Second * 2)
	defer Ticker.Stop()
	go func() {
		for {
			<-Ticker.C
			Check(SourPath, TarPath, fileModTimes)
		}
	}()

	select {}
}

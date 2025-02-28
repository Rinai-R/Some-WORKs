package Handle

import (
	pb "Golang/2025/02February/20250228/file/service/file/proto"
	"errors"
	"fmt"
	"io"
	"os"
)

func (*FileService) DownloadFile(stream pb.File_DownloadFileServer) error {
	var outFile *os.File
	var buffer = make([]byte, 1024)
	fmt.Println("进入下载")
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	if _, err = os.Stat(req.Path); os.IsNotExist(err) {
		return errors.New("file not exist")
	}
	outFile, err = os.OpenFile(fmt.Sprintf("%s/%s", req.Path, req.Filename), os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer outFile.Close()
	i := 0
	for {
		fmt.Println("读取文件流 ", i)
		i++
		n, err := outFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := &pb.DownloadRsp{
			Content: buffer[:n],
			Base: &pb.BaseRsp{
				Message: "ok",
				Success: true,
			},
		}
		err = stream.Send(chunk)
		if err != nil {
			return err
		}
	}

	return nil
}

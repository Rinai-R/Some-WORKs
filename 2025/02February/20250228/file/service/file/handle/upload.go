package Handle

import (
	pb "Golang/2025/02February/20250228/file/service/file/proto"
	"fmt"
	"io"
	"os"
)

type FileService struct {
	pb.UnimplementedFileServer
}

func (*FileService) UploadFile(stream pb.File_UploadFileServer) error {
	var filePath string
	var outFile *os.File
	fmt.Println("访问服务")
	var filedata []byte
	for {
		filechunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if filePath == "" {
			filePath = filechunk.Path
			if _, err = os.Stat(filePath); os.IsNotExist(err) {
				err = os.MkdirAll(filePath, 0755)
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
			outFile, err = os.Create(fmt.Sprintf("%s/%s", filePath, filechunk.FileName))
			defer outFile.Close()
		}
		_, err = outFile.Write(filechunk.Content)
		if err != nil {
			fmt.Println(err)
			return err
		}

		filedata = append(filedata, filechunk.Content...)
		fmt.Printf("Received chunk %d, size: %d\n", filechunk.Index, len(filechunk.Content))
	}
	return stream.SendAndClose(&pb.BaseRsp{
		Success: true,
		Message: "File uploaded successfully",
	})
}

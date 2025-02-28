package handle

import (
	"Golang/2025/02February/20250228/file/service/api/fileclient"
	pb "Golang/2025/02February/20250228/file/service/api/fileclient/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Download(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	path := ctx.PostForm("path")
	//向服务发送请求，建立流连接

	//建立流连接
	stream, err := fileclient.FileClient.DownloadFile(context.Background(), &pb.DownloadReq{
		Path:     path,
		Filename: filename,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 499,
			"msg":  err.Error(),
		})
		return
	}
	defer stream.CloseSend()

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Type", "application/octet-stream")
	for {
		//循环读取文件流
		rsp := &pb.DownloadRsp{}
		rsp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(500, gin.H{
				"code": 501,
				"msg":  err.Error(),
			})
			return
		}
		if _, err = ctx.Writer.Write(rsp.Content); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 503,
				"msg":  err.Error(),
			})
			return
		}
		ctx.Writer.Flush()
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})

}

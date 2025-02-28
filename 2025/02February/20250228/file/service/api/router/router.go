package router

import (
	"Golang/2025/02February/20250228/file/service/api/fileclient"
	pb "Golang/2025/02February/20250228/file/service/api/fileclient/proto"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func InitRouter() {
	router := gin.Default()

	router.PUT("/FILE", func(ctx *gin.Context) {
		file, _, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  err.Error(),
			})
			return
		}
		path := ctx.PostForm("path")
		filename := ctx.PostForm("filename")
		stream, err := fileclient.FileClient.UploadFile(context.Background())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}

		buffer := make([]byte, 1024)
		chunkIndex := 0
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 401,
					"msg":  err.Error(),
				})
				return
			}

			// 发送文件块到 gRPC 服务
			if err := stream.Send(&pb.FileChunk{
				Content:  buffer[:n],
				Index:    int32(chunkIndex),
				Path:     path,
				FileName: filename,
			}); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 402,
					"msg":  err.Error(),
				})
				return
			}

			chunkIndex++
		}
		resp, err := stream.CloseAndRecv()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 403,
				"msg":  err.Error(),
			})
			return
		}
		log.Printf("gRPC Response: %s", resp.Message)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  resp,
		})
	})

	err := router.Run(":8181")
	if err != nil {
		return
	}
}

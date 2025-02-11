package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger() {
	file, _ := os.OpenFile("./2025/02February/20250210/kitex-etcd/App/Initialize/Logger/log/logger.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	writerSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	Logger = zap.New(core)

}

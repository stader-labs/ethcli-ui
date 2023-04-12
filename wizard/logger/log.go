package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func init() {
	logFile := "./wizard.log"
	config := zap.NewProductionConfig()

	// Log.Out = os.Stdout
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{f.Name()}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	Log = logger.Sugar()
}

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// log ...
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		// if there is issue init, the don't start app.. 
		panic(err)
	}
}

// // GetLogger ...
// func GetLogger() {
// 	return log
// }

// Info ...
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

// Error ...
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}
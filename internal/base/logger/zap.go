package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// InitLogger initialises a logger while also setting up the global logger.
func InitLogger(level string) (theLogger Logger, cleanupFunc func()) {
	parsedLevel, err := zap.ParseAtomicLevel(level)
	if err != nil {
		panic(fmt.Sprintf("Invalid level provided: %s", level))
	}

	cfg := zap.Config{
		Level: parsedLevel,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:          "json",
		DisableStacktrace: false,
		DisableCaller:     false,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		Development:       false,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	zapLogger := zap.Must(cfg.Build())
	defer zapLogger.Sync()

	cleanupFunc = zap.ReplaceGlobals(zapLogger)
	theLogger = zapLogger.Sugar()
	theLogger.Info("logger construction succeeded")
	return theLogger, cleanupFunc
}

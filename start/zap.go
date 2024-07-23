package start

import "go.uber.org/zap"

func Zap() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{
		"../../logs/log.log",
		"stdout",
	}
	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
		return
	}
	zap.ReplaceGlobals(logger)
}

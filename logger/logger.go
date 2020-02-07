package logger

import (
	"errors"

	"go.uber.org/zap"
)

func GetLogger(detail, path string) (*zap.Logger, error) {
	var (
		err    error
		config zap.Config
	)
	switch detail {
	case "prod":
		config = zap.NewProductionConfig()
	case "dev":
		config = zap.NewDevelopmentConfig()
	default:
		return nil, errors.New("Unknow detail")
	}
	if path != "" {
		config.OutputPaths = []string{
			path,
		}
	}
	logger, err := config.Build()
	if err != nil {
		return nil, errors.New("Error build logger")
	}
	defer logger.Sync()
	return logger, nil
}

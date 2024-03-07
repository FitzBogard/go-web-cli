package initialize

import (
	"os"

	"github.com/spf13/viper"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func Logger() error {
	level := viper.GetString("log.level")
	levelStr, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}
	logger.WithLevel(levelStr)
	return nil
}

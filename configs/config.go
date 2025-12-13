package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadConfig() (*viper.Viper, error) {
	logger, _ := zap.NewProduction()
	localViper := viper.New()
	localViper.SetConfigName("config")
	localViper.SetConfigType("yaml")
	localViper.AddConfigPath("./configs")

	if err := localViper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			logger.Info("Config file not found, using default settings")
		} else {
			return localViper, fmt.Errorf("found config file, but encountered an error : %v", err)
		}
	}
	return localViper, nil
}

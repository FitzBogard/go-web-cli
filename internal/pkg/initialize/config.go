package initialize

import (
	"fmt"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(confPath string) error {
	configDir, fileName := filepath.Split(confPath)
	if fileName == "" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName(fileName)
	}
	if configDir == "" {
		viper.AddConfigPath("./")
	} else {
		viper.AddConfigPath(configDir)
	}
	viper.SetConfigType("yaml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed: %w \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed...")
	})
	return nil
}

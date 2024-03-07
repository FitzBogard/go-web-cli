package initialize

import (
	"context"
	redis2 "go-web-cli/internal/pkg/redis"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type RedisSetting struct {
	ID       string `mapstructure:"id"`
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func Redis() error {
	var settings []*RedisSetting
	err := viper.UnmarshalKey("redis", &settings)
	if err != nil {
		return err
	}

	for _, setting := range settings {
		rdb := redis.NewClient(&redis.Options{
			Addr:     setting.Addr,
			Password: setting.Password,
			DB:       setting.DB,
		})

		err = rdb.Ping(context.Background()).Err()
		if err != nil {
			return err
		}
		redis2.Set(setting.ID, rdb)
	}

	return nil
}

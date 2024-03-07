package cmd

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"time"

	"go-web-cli/internal/pkg/initialize"
	"go-web-cli/pkg/biz_name/delivery/http"
	"go-web-cli/pkg/biz_name/repository"
	"go-web-cli/pkg/biz_name/usecase"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	app      *server.Hertz
	ConfPath string
)

func init() {
	rootCmd.Flags().StringVar(&ConfPath, "config", "config.yaml", "use config file (default config.yaml)")
}

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// defer stdlog.Flush()
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if ok {
					stdlog.Panic("error: ", err)
				}
				stdlog.Panic(fmt.Errorf("unknown error: %v", err))
			}
		}()

		err := startup(ctx)
		if err != nil {
			stdlog.Panic(fmt.Errorf("fail to startup: %v", err))
			return
		}
		err = shutdown(ctx)
		if err != nil {
			stdlog.Println(fmt.Errorf("fail to shutdown: %v", err))
			return
		}
	},
}

func startup(ctx context.Context) error {
	// config
	err := initialize.InitConfig(ConfPath)
	if err != nil {
		stdlog.Printf("startup: fail to initialize config, err:%v\n", err)
		return err
	}

	// log
	err = initialize.Logger()
	if err != nil {
		return err
	}

	// metric
	err = initialize.Metric()
	if err != nil {
		return err
	}

	// mysql
	err = initialize.Mysql()
	if err != nil {
		return err
	}

	// redis
	err = initialize.Redis()
	if err != nil {
		return err
	}

	// handlers
	repo := repository.NewRepository()
	uc := usecase.NewUsecase(repo)
	handler := http.NewHandler(uc)

	// hertz web engine
	app = initialize.Hertz(handler)
	app.Spin()
	return nil
}

func shutdown(ctx context.Context) error {
	log.Info().Msg("service is shutting down...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			app.Shutdown(ctx)

			log.Info().Msg("service was shutdown successfully")
			return nil
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

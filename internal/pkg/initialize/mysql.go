package initialize

import (
	"context"
	"database/sql"
	"go-web-cli/internal/pkg/mysql"
	"go-web-cli/pkg/biz_name/domain"
	"sync"

	"github.com/spf13/viper"
)

func Mysql() error {
	var options []*mysql.Options
	err := viper.UnmarshalKey("mysql", &options)
	if err != nil {
		return err
	}

	for _, opts := range options {
		db, err := mysql.InitMysqlDB(opts)
		if err != nil {
			return err
		}
		set(opts.ID, db)
	}

	return nil
}

var (
	sqlInstance = make(map[string]*sql.DB)
	mu          sync.RWMutex
)

func set(name string, db *sql.DB) {
	mu.Lock()
	defer mu.Unlock()
	sqlInstance[name] = db
}

func Get(name string) (*sql.DB, error) {
	mu.RLock()
	defer mu.RUnlock()

	client, found := sqlInstance[name]
	if found {
		return client, nil
	}

	return nil, domain.ErrNotFound
}

var (
	ctxKey = &struct {
		name string
	}{
		name: "mysql",
	}
)

func FromContext(ctx context.Context) *sql.DB {
	client, ok := ctx.Value(ctxKey).(*sql.DB)
	if !ok {
		return nil
	}
	return client
}

func ToContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, ctxKey, db)
}

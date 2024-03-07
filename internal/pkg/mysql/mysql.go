package mysql

import (
	"database/sql"
	"fmt"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type Options struct {
	ID              string        `mapstructure:"id"`
	DSN             string        `mapstructure:"dsn"`
	Database        string        `mapstructure:"database"`
	Host            string        `mapstructure:"host"`
	Port            int64         `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_life_time"`
}

func InitMysqlDB(opts *Options) (*sql.DB, error) {
	dsn := opts.DSN
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", opts.User, opts.Password, opts.Host, opts.Port, opts.Database)
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(opts.MaxIdleConns)
	db.SetMaxOpenConns(opts.MaxOpenConns)
	db.SetConnMaxIdleTime(opts.ConnMaxIdleTime)
	db.SetConnMaxLifetime(opts.ConnMaxLifetime)

	return db, nil
}

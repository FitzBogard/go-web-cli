package domain

import (
	"context"
)

// business usecase and repo interface definition
type Usecase interface {
	Ping(ctx context.Context) (interface{}, error)
}
type Repository interface {
	Ping(ctx context.Context) (interface{}, error)
}

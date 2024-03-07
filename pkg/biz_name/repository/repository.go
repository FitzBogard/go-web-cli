package repository

import (
	"context"
)

type Repository struct{}

func (u *Repository) Ping(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func NewRepository() *Repository {
	return &Repository{}
}

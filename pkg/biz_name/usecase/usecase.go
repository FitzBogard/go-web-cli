package usecase

import (
	"context"
	"go-web-cli/pkg/biz_name/domain"
)

type Usecase struct {
	repo domain.Repository
}

func (u *Usecase) Ping(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func NewUsecase(repo domain.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

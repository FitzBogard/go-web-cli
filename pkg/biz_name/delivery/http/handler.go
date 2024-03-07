package http

import (
	"context"
	"go-web-cli/pkg/biz_name/domain"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Handler struct {
	usecase domain.Usecase
}

func NewHandler(usecase domain.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// Ping .
func (h *Handler) Ping(ctx context.Context, c *app.RequestContext) {
	_, err := h.usecase.Ping(ctx)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}

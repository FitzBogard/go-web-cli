package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-web-cli/pkg/biz_name/domain"
	"log"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	errors2 "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ErrorHandle(ctx context.Context, c *app.RequestContext) {

	defer func() {
		if err := recover(); err != nil {
			log.Panic(fmt.Errorf("hertz: error handled: %v", err))
			c.JSON(http.StatusInternalServerError, domain.ErrInternal)
			panic(domain.ErrInternal) // throw it, recovery.Recovery() will catch it.
		}
	}()

	c.Next(ctx)

	if c.Response.StatusCode() == consts.StatusNotFound {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	err := c.Errors.Last()
	if err == nil {
		return
	}

	var Herr *errors2.Error
	if errors.As(err, &Herr) {
		if Herr.IsType(errors2.ErrorTypeBind) {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
	}

	log.Panic(fmt.Errorf("hertz: unknown err: %v", err))
	c.JSON(http.StatusInternalServerError, domain.ErrInternal)
}

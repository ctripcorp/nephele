package handler

import "github.com/nephele/context"

type HealthCheckHandler struct {
}

func (h *HealthCheckHandler) Handler() HandlerFunc {
	return func(ctx *context.Context) {

	}
}

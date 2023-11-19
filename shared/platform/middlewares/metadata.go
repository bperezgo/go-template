package middlewares

import (
	"github.com/bperezgo/go-template/shared/platform/handler"
	"github.com/bperezgo/go-template/shared/platform/handlertypes"
	"github.com/google/uuid"
)

type MetadataMiddleware struct {
	handler handler.Handler
}

func NewMetadataMiddleware(handler handler.Handler) *MetadataMiddleware {
	return &MetadataMiddleware{
		handler: handler,
	}
}

func (h *MetadataMiddleware) GetMethod() handler.HandlerMethod {
	return h.handler.GetMethod()
}

func (h *MetadataMiddleware) GetPath() string {
	return h.handler.GetPath()
}

func (h *MetadataMiddleware) Function(req handlertypes.Request) handlertypes.Response {
	req.Meta = &handlertypes.Meta{
		RequestId: uuid.NewString(),
	}

	return h.handler.Function(req)
}

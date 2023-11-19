package handler

import "github.com/bperezgo/go-template/shared/platform/handlertypes"

type HandlerMethod int64

const (
	GET HandlerMethod = iota
	Any
	DELETE
	HEAD
	POST
	OPTIONS
	PATCH
	PUT
)

func (s HandlerMethod) String() string {
	switch s {
	case GET:
		return "GET"
	case Any:
		return "Any"
	case DELETE:
		return "DELETE"
	case HEAD:
		return "HEAD"
	case POST:
		return "POST"
	case OPTIONS:
		return "OPTIONS"
	case PATCH:
		return "PATCH"
	case PUT:
		return "PUT"
	}
	return "unknown"
}

type Function func(req handlertypes.Request) (res handlertypes.Response)

type Handler interface {
	GetMethod() HandlerMethod
	GetPath() string
	Function(req handlertypes.Request) (res handlertypes.Response)
}

type BaseHandler struct{}

func (h *BaseHandler) GetEmptyRequest() handlertypes.Request {
	return handlertypes.Request{}
}

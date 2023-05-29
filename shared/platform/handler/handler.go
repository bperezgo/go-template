package handler

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

type Response struct {
	Body       interface{}
	HttpStatus int
}

type Request struct {
	Body   interface{}
	Query  interface{}
	Params interface{}
}

type Function func(req Request) (res Response)

type Handler interface {
	GetMethod() HandlerMethod
	GetPath() string
	Function(req Request) (res Response)
	GetEmptyRequest() Request
}

type BaseHandler struct{}

func (h *BaseHandler) GetEmptyRequest() Request {
	return Request{}
}

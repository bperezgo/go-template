package handler

type BasisHandler struct {
	HandlerMethod HandlerMethod
	Path          string
}

func (h *BasisHandler) GetMethod() HandlerMethod {
	return h.HandlerMethod
}

func (h *BasisHandler) GetPath() string {
	return h.Path
}

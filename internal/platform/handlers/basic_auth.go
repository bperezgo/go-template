package handlers

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/bperezgo/go-template/internal/domain/ports"
	"github.com/bperezgo/go-template/shared/platform/handler"
	"github.com/bperezgo/go-template/shared/platform/handlertypes"
)

type BasicAuthRequest struct{}

type BasicAuthResponse struct {
	Token    string `json:"token"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type BasicAuthMiddleware struct {
	userRepository ports.UserRepository
}

func NewBasicAuthMiddleware(userRepository ports.UserRepository) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{
		userRepository: userRepository,
	}
}

func (h *BasicAuthMiddleware) GetMethod() handler.HandlerMethod {
	return handler.POST
}

func (h *BasicAuthMiddleware) GetPath() string {
	return "/auth"
}

// This can be the middleware basic auth function
func (h *BasicAuthMiddleware) Function(req handlertypes.Request) handlertypes.Response {
	authentication := req.Headers.Authorization

	authentication = authentication[6:]

	userAndPassword, err := base64.StdEncoding.DecodeString(authentication)
	if err != nil {
		return handlertypes.Response{
			Body:       err.Error(),
			HttpStatus: http.StatusBadRequest,
		}
	}

	userAndPasswordSplited := strings.Split(string(userAndPassword), ":")
	if len(userAndPasswordSplited) != 2 {
		return handlertypes.Response{
			Body:       "Invalid credentials",
			HttpStatus: http.StatusBadRequest,
		}
	}

	user := userAndPasswordSplited[0]
	password := userAndPasswordSplited[1]

	return handlertypes.Response{
		Body: BasicAuthResponse{
			Token:    "token",
			User:     user,
			Password: password,
		},
		HttpStatus: http.StatusOK,
	}
}

func (h *BasicAuthMiddleware) GetEmptyRequest() handlertypes.Request {
	return handlertypes.Request{
		Body: BasicAuthRequest{},
	}
}

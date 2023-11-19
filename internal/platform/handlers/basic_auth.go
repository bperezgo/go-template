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
	handler.BasisHandler
}

func NewBasicAuthMiddleware(userRepository ports.UserRepository) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{
		userRepository: userRepository,
		BasisHandler: handler.BasisHandler{
			HandlerMethod: handler.POST,
			Path:          "/auth",
		},
	}
}

// This can be the middleware basic auth function
func (h *BasicAuthMiddleware) Function(req handlertypes.Request) handlertypes.Response {
	authentication := req.Headers.Authorization

	authentication = authentication[6:] // After the word "Basic " (6 characters) comes the user and password encoded in base64

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

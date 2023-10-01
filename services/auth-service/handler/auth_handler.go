package handler

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/jnst/monorepo-go/protos/gen/auth/v1"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) SendPasscode(
	_ context.Context,
	_ *connect.Request[v1.SendPasscodeRequest],
) (*connect.Response[v1.SendPasscodeResponse], error) {
	return connect.NewResponse(&v1.SendPasscodeResponse{
		Success: true,
	}), nil
}

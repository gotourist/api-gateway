package handler

import (
	"github.com/iman_task/api-gateway/api/handlers/post"
)

// HandlerV1 ...
type HandlerV1 struct {
	PostHandler post.PostHandler
}

func New(h *HandlerV1) *HandlerV1 {
	return &HandlerV1{
		PostHandler: h.PostHandler,
	}
}

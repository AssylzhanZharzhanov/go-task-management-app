package http

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {

}

func (h *Handler) List(c *gin.Context) {

}

func (h *Handler) GetByID(c *gin.Context) {

}

func (h *Handler) Update(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}

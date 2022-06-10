package http

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.RouterGroup, service user.Service) {
	h := NewHandler(service)

	users := router.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("", h.List)
		users.GET(":id", h.GetByID)
		users.PUT(":id", h.Update)
		users.DELETE("", h.Delete)
	}
}

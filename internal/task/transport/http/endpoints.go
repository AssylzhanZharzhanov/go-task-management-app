package http

import (
	"github.com/AssylzhanZharzhanov/task-management-app/internal/task"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.RouterGroup, service task.Service) {
	h := NewHandler(service)

	users := router.Group("/tasks")
	{
		users.POST("", h.Create)
		users.GET("", h.List)
		users.GET(":id", h.GetByID)
		users.PUT(":id", h.Update)
		users.DELETE(":id", h.Delete)
	}
}

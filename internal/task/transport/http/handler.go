package http

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/task"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/error"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/task"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service task.Service
}

func NewHandler(service task.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var input *domain.CreateTaskDTO
	if err := c.BindJSON(&input); err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	taskID, err := h.service.Create(input)
	if err != nil {
		error.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, taskID)

}

func (h *Handler) List(c *gin.Context) {
	tasks, err := h.service.List()
	if err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	taskID, err := h.service.GetByID(domain.TaskID(id))
	if err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, taskID)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input *domain.UpdateTaskDTO
	if err := c.BindJSON(&input); err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.ID = domain.TaskID(id)
	taskID, err := h.service.Update(input)
	if err != nil {
		error.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, taskID)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.Delete(domain.TaskID(id))
	if err != nil {
		error.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

package http

import (
	domain "github.com/AssylzhanZharzhanov/task-management-app/internal/domain/user"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/error"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var input *domain.CreateUserDTO
	if err := c.BindJSON(&input); err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, err := h.service.Create(input)
	if err != nil {
		error.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, userID)
}

func (h *Handler) List(c *gin.Context) {
	users, err := h.service.List()
	if err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	userID, err := h.service.GetByID(domain.UserID(id))
	if err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userID)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input *domain.UpdateUserDTO
	if err := c.BindJSON(&input); err != nil {
		error.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.ID = domain.UserID(id)
	userID, err := h.service.Update(input)
	if err != nil {
		error.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, userID)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.Delete(domain.UserID(id))
	if err != nil {
		error.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

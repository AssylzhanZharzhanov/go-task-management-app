package app

import (
	"context"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/task"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server

	taskService task.Service
	userService task.Service
	Port        string
}

func NewApp(port string) *App {
	//add use cases

	return &App{
		Port: port,
	}
}

func (s *App) Run() error {

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//register endpoints

	s.httpServer = &http.Server{
		Addr:           ":" + s.Port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *App) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

package app

import (
	"context"
	"net/http"
	"time"

	taskRepo "github.com/AssylzhanZharzhanov/task-management-app/internal/task/repository"
	tasksService "github.com/AssylzhanZharzhanov/task-management-app/internal/task/service"
	tasksEndpoints "github.com/AssylzhanZharzhanov/task-management-app/internal/task/transport/http"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/user"
	userRepo "github.com/AssylzhanZharzhanov/task-management-app/internal/user/repository"
	usersService "github.com/AssylzhanZharzhanov/task-management-app/internal/user/service"
	usersEndpoints "github.com/AssylzhanZharzhanov/task-management-app/internal/user/transport/http"

	"github.com/AssylzhanZharzhanov/task-management-app/internal/task"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	httpServer *http.Server

	DB          *gorm.DB
	Port        string

	userService user.Service
	taskService task.Service
}

func NewApp(db *gorm.DB, port string) *App {

	// Repositories
	usersRepository := userRepo.NewPostgresRepository(db)
	tasksRepository := taskRepo.NewPostgresRepository(db)

	return &App{
		DB:   db,
		Port: port,

		// Services
		userService: usersService.NewService(usersRepository),
		taskService: tasksService.NewService(tasksRepository),
	}
}

func (s *App) Run() error {

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")

	//register endpoints
	usersEndpoints.RegisterEndpoints(api, s.userService)
	tasksEndpoints.RegisterEndpoints(api, s.taskService)

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

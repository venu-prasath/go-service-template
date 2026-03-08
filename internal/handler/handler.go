package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/venumohan/go-service-template/internal/middleware"
	"github.com/venumohan/go-service-template/internal/service"
)

type Handler struct {
	userService *service.UserService
	jwtSecret   string
}

func New(userService *service.UserService, jwtSecret string) *Handler {
	return &Handler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *Handler) SetupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/healthz", h.HealthCheck)

	api := r.Group("/api/v1")
	{
		api.POST("/auth/register", h.Register)
		api.POST("/auth/login", h.Login)

		protected := api.Group("")
		protected.Use(middleware.Auth(h.jwtSecret))
		{
			protected.GET("/users/me", h.GetCurrentUser)
			protected.GET("/users/:id", h.GetUser)
		}
	}

	return r
}

package handler

import (
	"segmentation-service/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/:id", h.createUser)
		user.GET("/", h.getAllUsers)
		user.GET("/:id", h.getUserById)
		user.PATCH("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}

	return router
}


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
		user.GET("/", h.getAllUsers)
		user.POST("/", h.createUser)
		user.DELETE("/:id", h.deleteUser)
	}

	segment := router.Group("/segment")
	{
		segment.GET("/", h.getAllSegments)
		segment.POST("/", h.createSegment)
		segment.PATCH(":id", h.updateSegment)
		segment.DELETE("/", h.deleteSegment)
	}

	return router
}


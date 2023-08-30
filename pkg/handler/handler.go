package handler

import (
	"segmentation-service/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	user_segment := router.Group("/user-segment")
	{
		user_segment.POST("/", h.createUserSegment)
		user_segment.DELETE("/", h.deleteUserSegment)
		user_segment.GET("/:userId", h.getSegmentsByUserId)
	}

	return router
}

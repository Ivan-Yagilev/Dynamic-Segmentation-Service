package handler

import "github.com/gin-gonic/gin"

type Handler struct {

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


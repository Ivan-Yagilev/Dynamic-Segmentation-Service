package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {

}

func (h *Handler) getAllUsers(c *gin.Context) {

}

func (h *Handler) getUserById(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {
	// userId, err := getUserId(c)
	// if err != nil {
	// 	return
	// }

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
	// 	return
	// }

	// 

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
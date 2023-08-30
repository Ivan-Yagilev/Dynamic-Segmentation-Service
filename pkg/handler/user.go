package handler

import (
	"net/http"
	segmentation_service "segmentation-service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllUsersResponse struct {
	Data []segmentation_service.UserResponse `json:"data"`
}

// @Summary Get All Users
// @Tags Users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllUsersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	lists, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: lists,
	})
}

// @Summary Create user
// @Tags Users
// @Description create user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body segmentation_service.User true "user info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [post]
func (h *Handler) createUser(c *gin.Context) {
	var input segmentation_service.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Delete user
// @Tags Users
// @Description delete user by id
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	err = h.services.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

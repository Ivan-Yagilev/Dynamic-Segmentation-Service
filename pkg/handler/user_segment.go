package handler

import (
	"net/http"
	segmentation_service "segmentation-service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create user-segment
// @Tags UserSegment
// @Description create user-segment by user id and list of segments
// @ID create-user-segment
// @Accept  json
// @Produce  json
// @Param input body segmentation_service.UserSegment true "user-segment info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user-segment/ [post]
func (h *Handler) createUserSegment(c *gin.Context) {
	var input segmentation_service.UserSegment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.CreateUserSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete user-segment
// @Tags UserSegment
// @Description delete user-segment by user id and list of segments
// @ID delete-user-segment
// @Accept  json
// @Produce  json
// @Param input body segmentation_service.UserSegment true "user-segment info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user-segment/ [delete]
func (h *Handler) deleteUserSegment(c *gin.Context) {
	var input segmentation_service.UserSegment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.DeleteUserSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getSegmentsByUserIdResponse struct {
	Data []segmentation_service.Segment `json:"data"`
}

// @Summary Get User Segments
// @Tags UserSegment
// @Description get active user segments
// @ID get-all-user-segments
// @Accept  json
// @Produce  json
// @Param id path int true "Segment ID" Format(int64)
// @Success 200 {object} getSegmentsByUserIdResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user-segment/{id} [get]
func (h *Handler) getSegmentsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	lists, err := h.services.UserSegment.GetAllSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getSegmentsByUserIdResponse{
		Data: lists,
	})
}

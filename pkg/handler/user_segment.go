package handler

import (
	"net/http"
	segmentation_service "segmentation-service"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

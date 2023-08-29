package handler

import (
	"net/http"
	segmentation_service "segmentation-service"

	"github.com/gin-gonic/gin"
)

type getAllSegmentsResponse struct {
	Data []segmentation_service.Segment `json:"data"`
}

func (h *Handler) getAllSegments(c *gin.Context) {
	lists, err := h.services.GetAllSegments()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSegmentsResponse{
		Data: lists,
	})
}

func (h *Handler) createSegment(c *gin.Context) {
	var input segmentation_service.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {
	var input segmentation_service.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.DeleteSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
package handler

import (
	"net/http"
	segmentation_service "segmentation-service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUserSegment(c *gin.Context) {
	var input segmentation_service.UserSegment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUserSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
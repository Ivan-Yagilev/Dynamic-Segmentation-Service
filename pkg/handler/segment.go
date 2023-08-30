package handler

import (
	"net/http"
	segmentation_service "segmentation-service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllSegmentsResponse struct {
	Data []segmentation_service.Segment `json:"data"`
}

// @Summary Get All Segments
// @Tags Segments
// @Description get all segments
// @ID get-all-segments
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllSegmentsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/ [get]
func (h *Handler) getAllSegments(c *gin.Context) {
	lists, err := h.services.Segment.GetAllSegments()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSegmentsResponse{
		Data: lists,
	})
}

// @Summary Update Segment
// @Tags Segments
// @Description update segmentname by id
// @ID update-segment
// @Accept  json
// @Produce  json
// @Param id path int true "Segment ID" Format(int64)
// @Param input body segmentation_service.UpdateSegmentInput true "segmentname"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/{id} [patch]
func (h *Handler) updateSegment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input segmentation_service.UpdateSegmentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateSegment(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}


// @Summary Create segment
// @Tags Segments
// @Description create segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body segmentation_service.SegmentWOId true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/ [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input segmentation_service.SegmentWOId

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

// @Summary Delete segment
// @Tags Segments
// @Description delete segment by name
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body segmentation_service.SegmentWOId true "segmentname"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment/ [delete]
func (h *Handler) deleteSegment(c *gin.Context) {
	var input segmentation_service.SegmentWOId

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

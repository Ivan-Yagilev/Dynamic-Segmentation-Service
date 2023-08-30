package handler

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// const segmentCtx = "segmentId"

// func getSegmentId(c *gin.Context) (int, error) {
// 	id, ok := c.Get(segmentCtx)
// 	if !ok {
// 		newErrorResponse(c, http.StatusInternalServerError, "no segment id")
// 		return 0, errors.New("no segment id")
// 	}

// 	idInt, ok := id.(int)
// 	if !ok {
// 		newErrorResponse(c, http.StatusInternalServerError, "invalid id type")
// 		return 0, errors.New("invalid id type")
// 	}

// 	return idInt, nil
// }

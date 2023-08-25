package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "no user id")
		return 0, errors.New("no user id")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id type")
		return 0, errors.New("invalid id type")
	}

	return idInt, nil
}
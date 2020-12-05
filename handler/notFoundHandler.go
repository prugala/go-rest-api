package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(
		http.StatusNotFound,
		model.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
		},
	)
}

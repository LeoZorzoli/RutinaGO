package Middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rutina.api/Contracts"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var statusCode int
			switch err.(type) {
			case *gin.Error:
				statusCode = http.StatusBadRequest
			default:
				statusCode = http.StatusInternalServerError
			}

			response := Contracts.BaseResponse{
				Success: false,
				Message: err.Error(),
			}

			c.JSON(statusCode, response)
			c.Abort()
		}
	}
}

package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func SuccessMessage(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data: data,
	})
}

func ErrorMessage(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, APIResponse{
		Success: false,
		Message: message,
		Error: data.(string),
	})
}
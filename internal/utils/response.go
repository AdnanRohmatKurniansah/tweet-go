package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type Pagination struct {
	Data interface{} `json:"data"`
	Total int64 `json:"total"`
	Page int `json:"page"`
	Limit int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

type PaginatedAPIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Total int64 `json:"total"`
	Page int `json:"page"`
	Limit int `json:"limit"`
	TotalPages int `json:"total_pages"`
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
		Error: data,
	})
}

func PaginationMessage(c *gin.Context, status int, message string, data interface{}, total int64, page, limit, totalPages int) {
	c.JSON(status, PaginatedAPIResponse{
		Success: true,
		Message: message,
		Data: data,
		Total: total,
		Page: page,
		Limit: limit,
		TotalPages: totalPages,
	})
}
package util

import "github.com/gin-gonic/gin"

// utilities which are used to return responses on successful API calls

// Response struct
// returns JSON Formatted success message with Struct data
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// returns JSON Formatted error response
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

// returns JSON formatted success message
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}
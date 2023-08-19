package helpers

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
}

func HandleError(c *gin.Context, statusCode int, errMsg string, err error) {
    c.JSON(statusCode, ErrorResponse{
        Error:   errMsg,
        Message: err.Error(),
    })
}
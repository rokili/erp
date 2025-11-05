package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemStatusResponse struct {
	IsOpened bool   `json:"is_opened"`
	Message  string `json:"message"`
}

// GetSystemStatus 获取系统状态
func GetSystemStatus(c *gin.Context) {
	c.JSON(http.StatusOK, SystemStatusResponse{
		IsOpened: false,
		Message:  "系统未开账",
	})
}

// OpenSystem 系统开账
func OpenSystem(c *gin.Context) {
	// TODO: 实现系统开账逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "系统开账成功",
	})
}

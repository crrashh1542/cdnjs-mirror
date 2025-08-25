package handlers

import (
	"github.com/gin-gonic/gin"
)

var (
	siteURL string
)

func SetSiteURL(url string) {
	siteURL = url
}

// getStatus 用于返回站点域名及端口号的信息
func HandleGetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"site":    siteURL,
	})
}
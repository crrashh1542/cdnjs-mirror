package handlers

import (
	"github.com/gin-gonic/gin"
	"cdnjs-mirror/config"
)

var (
	siteURL string
	buildId = "dev"
)

func SetSiteURL(url string) {
	siteURL = url
}
func SetBuildId(id string) {
	buildId = id
}

// getStatus 用于返回站点域名及端口号的信息
func HandleGetStatus(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		return
	}
	
	c.JSON(200, gin.H{
		"code":    200,
		"site":    siteURL,
		"version": config.Version,
		"build":   buildId,
	})
}
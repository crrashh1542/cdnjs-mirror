package router

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"cdnjs-mirror/config"
	"cdnjs-mirror/handlers"
	"cdnjs-mirror/utils"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 1）index & icon
	indexFile := filepath.Join(config.StaticDir, "index.html")
	r.StaticFile("/favicon.ico", filepath.Join(config.StaticDir, "favicon.ico"))
	if utils.IsFileExists(indexFile) {
		r.StaticFile("/", indexFile)
		r.StaticFile("/index.html", indexFile)
	} else {
		fmt.Printf("[%s] 未检测到主页文件，使用默认主页\n", time.Now().Format("06-01-02 15:04:05"))
		// 提供一个简单的默认主页
		r.GET("/", func(c *gin.Context) {
			c.String(200, `<html>
<head>
    <title>CDNJS Mirror</title>
    <meta charset="utf-8" />
</head>
<body>
    <h1>CDNJS Mirror</h1>
    <p>服务正在运行中！</p>
</body>
</html>`)
		})
	}

	// 2）_assets
	if utils.IsDirExists(config.AssetsDir) {
		fmt.Printf("[%s] 正在启动静态资源服务\n", time.Now().Format("06-01-02 15:04:05"))
		r.Static("/_assets", config.AssetsDir)
	} else {
		fmt.Printf("[%s] 未检测到静态资源目录，准备创建\n", time.Now().Format("06-01-02 15:04:05"))
	}

	// 3）getStatus
	r.GET("/getStatus", handlers.HandleGetStatus)

	// 4）CDNJS
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		
		// 排除特殊路径
		if path == "/" || path == "/index.html" || path == "/getStatus" || 
		   strings.HasPrefix(path, "/_assets") {
			c.Status(http.StatusNotFound)
			return
		}
		
		// 处理CDN请求
		handlers.HandleCDNJSRequest(c)
	})

	return r
}
package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"cdnjs-mirror/handlers"
)

func SetupRouter(staticFS embed.FS) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	fmt.Printf("[%s] 检查嵌入的静态文件...\n", time.Now().Format("06-01-02 15:04:05"))
	
	// 1) index.html
	if _, err := fs.Stat(staticFS, "static/index.html"); err == nil {
		r.GET("/", func(c *gin.Context) {
			content, err := fs.ReadFile(staticFS, "static/index.html")
			if err != nil {
				fmt.Printf("[%s] 读取嵌入文件失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
				c.String(500, "无法读取主页文件")
				return
			}
			c.Data(200, "text/html; charset=utf-8", content)
		})
		
		r.GET("/index.html", func(c *gin.Context) {
			content, err := fs.ReadFile(staticFS, "static/index.html")
			if err != nil {
				fmt.Printf("[%s] 读取嵌入文件失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
				c.String(500, "无法读取主页文件")
				return
			}
			c.Data(200, "text/html; charset=utf-8", content)
		})
	} else {
		fmt.Printf("[%s] 未检测到嵌入的主页文件: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
		// 如果出错导致没有 index.html 被打包，提供一个简单的默认主页
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

	// 2）favicon.ico
	if _, err := fs.Stat(staticFS, "static/favicon.ico"); err == nil {
		r.GET("/favicon.ico", func(c *gin.Context) {
			content, err := fs.ReadFile(staticFS, "static/favicon.ico")
			if err != nil {
				fmt.Printf("[%s] 读取嵌入文件失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
				c.String(500, "无法读取图标文件")
				return
			}
			c.Data(200, "image/x-icon", content)
		})
	}

	// 3）_assets 静态资源
	if _, err := fs.Stat(staticFS, "static/_assets"); err == nil {
		r.GET("/_assets/*filepath", func(c *gin.Context) {
			filepathParam := c.Param("filepath")
			if filepathParam == "" {
				c.Status(http.StatusNotFound)
				return
			}
			filepathParam = strings.TrimPrefix(filepathParam, "/")
			
			// 注：使用 path.Join 而不是 filepath.Join 来确保使用正斜杠
			// Windows 路径为反斜杠所以会导致 _assets 目录打包出错
			fullPath := path.Join("static", "_assets", filepathParam)
			
			content, err := fs.ReadFile(staticFS, fullPath)
			if err != nil {
				c.Status(http.StatusNotFound)
				return
			}
			
			var contentType string
			switch {
			case strings.HasSuffix(filepathParam, ".css"):
				contentType = "text/css"
			case strings.HasSuffix(filepathParam, ".js"):
				contentType = "application/javascript"
			case strings.HasSuffix(filepathParam, ".png"):
				contentType = "image/png"
			case strings.HasSuffix(filepathParam, ".jpg") || strings.HasSuffix(filepathParam, ".jpeg"):
				contentType = "image/jpeg"
			case strings.HasSuffix(filepathParam, ".ico"):
				contentType = "image/x-icon"
			default:
				contentType = "application/octet-stream"
			}
			c.Data(200, contentType, content)
		})
	} else {
		fmt.Printf("[%s] 未检测到嵌入的静态资源目录: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
	}

	// 4）getStatus
	r.GET("/getStatus", handlers.HandleGetStatus)

	// 5）CDNJS
	r.NoRoute(func(c *gin.Context) {
		pathParam := c.Request.URL.Path
		
		// 排除特殊路径
		if pathParam == "/" || pathParam == "/index.html" || pathParam == "/getStatus" || 
		   strings.HasPrefix(pathParam, "/_assets") {
			// 这些路径应该已经被其他路由处理了
			return
		}
		
		// 处理CDN请求
		handlers.HandleCDNJSRequest(c)
	})

	return r
}
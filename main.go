package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"cdnjs-mirror/config"
	"cdnjs-mirror/handlers"
	"cdnjs-mirror/router"
	"cdnjs-mirror/utils"
)

var (
	siteURL string
)

func main() {
	flag.StringVar(&siteURL, "h", "http://localhost:23467", "站点URL，例如: https://cdn.example.com")
	flag.Parse()
	siteURL = strings.TrimRight(siteURL, "/")
	handlers.SetSiteURL(siteURL)

	gin.SetMode(gin.ReleaseMode)
	utils.CreateDirs([]string{
		config.StaticDir, 
		config.AssetsDir, 
		config.LocalCacheDir,
	})
	fmt.Printf("[%s] 准备启动HTTP服务......\n", time.Now().Format("06-01-02 15:04:05"))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CDNJS 镜像路由
	r = router.SetupRouter()
	fmt.Printf("[%s] 正在启动CDN路由服务\n", time.Now().Format("06-01-02 15:04:05"))

	// 启动 Server
	port := "23657"
	addr := ":" + port
	fmt.Println("监听地址: http://localhost:23657")
	fmt.Printf("自定义访问地址: %s\n", siteURL)
	fmt.Printf("获取站点信息: %s/getStatus\n", siteURL)

	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
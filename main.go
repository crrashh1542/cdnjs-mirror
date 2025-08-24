package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	originalCDNJS = "https://cdnjs.cloudflare.com"
	localCacheDir = "./cdn"
	staticDir = "./static"
	assetsDir = "./static/_assets"
)

var (
	siteURL string
)

func main() {
	flag.StringVar(&siteURL, "h", "http://localhost:23467", "站点URL，例如: https://cdn.example.com")
	flag.Parse()
	siteURL = strings.TrimRight(siteURL, "/")

	gin.SetMode(gin.ReleaseMode)
	createDirs()
	fmt.Printf("[%s] 准备启动HTTP服务......\n", time.Now().Format("06-01-02 15:04:05"))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// static/_assets
	if isDirExists(assetsDir) {
		fmt.Printf("[%s] 正在启动静态资源服务\n", time.Now().Format("06-01-02 15:04:05"))
		r.Static("/_assets", assetsDir)
	} else {
		fmt.Printf("[%s] 未检测到静态资源目录，准备创建\n", time.Now().Format("06-01-02 15:04:05"))
	}

	// index.html
	indexFile := filepath.Join(staticDir, "index.html")
	r.StaticFile("/favicon.ico", filepath.Join(staticDir, "favicon.ico"))
	if isFileExists(indexFile) {
		r.StaticFile("/", indexFile)
		r.StaticFile("/index.html", indexFile)
	} else {
		fmt.Printf("[%s] 未检测到主页文件，使用默认主页\n", time.Now().Format("06-01-02 15:04:05"))
		// 提供一个简单的默认主页
		r.GET("/", func(c *gin.Context) {
			c.String(200, `<html>
<head>
    <title>CDNJS镜像</title>
    <link rel="stylesheet" href="/_assets/style.css">
</head>
<body>
    <h1>CDNJS镜像服务器</h1>
    <p>服务器运行正常</p>
    <p>CDN资源访问示例: /jquery/3.6.0/jquery.min.js</p>
    <script src="/_assets/main.js"></script>
</body>
</html>`)
		})
	}

	// getSite 用于返回站点域名及端口号的信息
	r.GET("/getSite", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"site": siteURL,
		})
	})

	// CDNJS 镜像路由
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		
		// 排除特殊路径
		if path == "/" || path == "/index.html" || path == "/health" || 
		   strings.HasPrefix(path, "/_assets") {
			// 这些路径应该已经被其他路由处理了
			c.Status(http.StatusNotFound)
			return
		}
		
		// 处理 CDN 请求
		handleCDNJSRequest(c)
	})
	fmt.Printf("[%s] 正在启动CDN路由服务\n", time.Now().Format("06-01-02 15:04:05"))

	// 启动 Server
	port := "23657"
	addr := ":" + port
	fmt.Println("监听地址: http://localhost:23657")
	fmt.Printf("自定义访问地址: %s\n", siteURL)
	fmt.Printf("健康检查: %s/health\n", siteURL)
	fmt.Printf("获取站点信息: %s/getSite\n", siteURL)

	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// 没有必须的目录就创建
func createDirs() {
	dirs := []string{staticDir, assetsDir, localCacheDir}
	for _, dir := range dirs {
		if !isDirExists(dir) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("创建目录失败 %s: %v\n", dir, err)
			} else {
				fmt.Printf("创建目录: %s\n", dir)
			}
		} else {
			fmt.Printf("目录已存在: %s\n", dir)
		}
	}
}
// 判断必需目录是否已存在
func isDirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
// 判断文件是否存在
func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// 处理 CDNJS 请求
func handleCDNJSRequest(c *gin.Context) {
	// 获取请求的文件路径（去掉开头的/）
	filePath := strings.TrimPrefix(c.Request.URL.Path, "/")
	if filePath == "" {
		fmt.Printf("[%s] 请求路径为空\n", time.Now().Format("06-01-02 15:04:05"))
		c.Status(http.StatusNotFound)
		return
	}

	// 构造本地文件路径
	localFilePath := filepath.Join(localCacheDir, filePath)

	// 检查本地是否已有缓存
	if isFileExists(localFilePath) {
		// 如果本地存在，直接返回
		c.File(localFilePath)
		return
	}

	// 如果本地不存在，从 CDNJS 缓存
	originalURL := originalCDNJS + "/ajax/libs/" + filePath
	fmt.Printf("[%s] 正在缓存: %s\n", time.Now().Format("06-01-02 15:04:05"), originalURL)

	// 创建目录
	dir := filepath.Dir(localFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("[%s] 创建目录失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
		c.String(http.StatusInternalServerError, "创建目录失败: %v", err)
		return
	}

	// 下载文件
	if err := downloadFile(originalURL, localFilePath); err != nil {
		fmt.Printf("[%s] 下载失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
		c.String(http.StatusInternalServerError, "下载文件失败: %v", err)
		return
	}

	fmt.Printf("[%s] 下载完成: %s\n", time.Now().Format("06-01-02 15:04:05"), localFilePath)
	// 返回下载的文件
	c.File(localFilePath)
}

func downloadFile(url, filepath string) error {
	fmt.Printf("[%s] 开始下载: %s\n", time.Now().Format("06-01-02 15:04:05"), url)
	
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("[%s] HTTP请求失败: %v", time.Now().Format("06-01-02 15:04:05"), err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[%s] HTTP状态码: %d", time.Now().Format("06-01-02 15:04:05"), resp.StatusCode)
	}

	// 创建本地文件
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("[%s] 创建文件失败: %v", time.Now().Format("06-01-02 15:04:05"), err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("[%s] 写入文件失败: %v", time.Now().Format("06-01-02 15:04:05"), err)
	}
	
	return nil
}
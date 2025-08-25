package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"cdnjs-mirror/config"
	"cdnjs-mirror/utils"
)

func HandleCDNJSRequest(c *gin.Context) {
	// 获取请求的文件路径（去掉开头的/）
	filePath := strings.TrimPrefix(c.Request.URL.Path, "/")
	if filePath == "" {
		fmt.Printf("[%s] 请求路径为空\n", time.Now().Format("06-01-02 15:04:05"))
		c.Status(http.StatusNotFound)
		return
	}

	// 构造本地文件路径
	localFilePath := filepath.Join(config.LocalCacheDir, filePath)

	// 检查本地是否已有缓存
	if utils.IsFileExists(localFilePath) {
		// 如果本地存在，直接返回
		c.File(localFilePath)
		return
	}

	// 如果本地不存在，从 CDNJS 缓存
	originalURL := config.OriginalCDNJS + "/ajax/libs/" + filePath
	fmt.Printf("[%s] 正在缓存: %s\n", time.Now().Format("06-01-02 15:04:05"), originalURL)

	// 创建目录
	dir := filepath.Dir(localFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("[%s] 创建目录失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
		c.String(http.StatusInternalServerError, "创建目录失败: %v", err)
		return
	}

	// 下载文件
	if err := utils.DownloadFile(originalURL, localFilePath); err != nil {
		fmt.Printf("[%s] 下载失败: %v\n", time.Now().Format("06-01-02 15:04:05"), err)
		c.String(http.StatusInternalServerError, "下载文件失败: %v", err)
		return
	}

	fmt.Printf("[%s] 下载完成: %s\n", time.Now().Format("06-01-02 15:04:05"), localFilePath)
	// 返回下载的文件
	c.File(localFilePath)
}
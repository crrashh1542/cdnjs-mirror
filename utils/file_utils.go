package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func IsDirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func CreateDirs(dirs []string) {
	for _, dir := range dirs {
		if !IsDirExists(dir) {
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

func DownloadFile(url, filepath string) error {
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
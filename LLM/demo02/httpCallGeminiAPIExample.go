package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("GEMINI_API_KEY 环境变量未设置")
		return
	}

	// 设置代理
	proxyUrl, err := url.Parse("http://127.0.0.1:10808")
	if err != nil {
		fmt.Printf("解析代理地址失败: %v\n", err)
		return
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Transport: transport,
	}

	// 设置请求 URL 和请求体
	geminiUrl := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + apiKey
	reqBody := `{
	"contents": [{
    	"parts":[{"text": "Explain how AI works"}]
    }]
   }`

	// 创建 POST 请求
	req, err := http.NewRequestWithContext(ctx, "POST", geminiUrl, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应
	fmt.Println(string(body))
}

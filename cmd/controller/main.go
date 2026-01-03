package main

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/api"
	"github.com/wangn9900/UltraForward/internal/assets"
	"github.com/wangn9900/UltraForward/internal/database"
)

func main() {
	log.Println("[UltraForward] 正在初始化旗舰控制中心...")

	// 1. 初始化数据库
	database.InitDB()
	log.Println("[UltraForward] 数据库连接成功！")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New() // 使用 New 减少默认中间件干扰
	r.Use(gin.Recovery())

	// 2. 极其简化的跨域处理
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 3. 准备前端文件系统
	// 直接从嵌入的 WebDist 中剥离 web-dist 根目录
	feFS, err := fs.Sub(assets.WebDist, "web-dist")
	if err != nil {
		log.Fatalf("[FATAL] 前端资源剥离失败: %v", err)
	}
	staticServer := http.FileServer(http.FS(feFS))

	// 4. API 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", api.RegisterHandler)
		v1.POST("/auth/login", api.LoginHandler)
		v1.GET("/stats", api.GetDashboardStats)
		v1.GET("/plans", api.GetPlans)
		v1.POST("/admin/plans", api.CreatePlan)
		v1.DELETE("/admin/plans/:id", api.DeletePlan)
	}

	// 5. 终极静态资源 & SPA 兜底逻辑
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 检查嵌入的文件系统中是否存在该文件
		_, err := fs.Stat(feFS, strings.TrimPrefix(path, "/"))

		if err == nil {
			// 文件存在，直接按静态文件返回
			staticServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		// 如果文件不存在且不是 API 请求，则返回 index.html (SPA 路由)
		if !strings.HasPrefix(path, "/api/") {
			indexData, err := fs.ReadFile(feFS, "index.html")
			if err != nil {
				c.String(404, "UI Bundle Incomplete")
				return
			}
			c.Data(200, "text/html; charset=utf-8", indexData)
			return
		}

		c.JSON(404, gin.H{"error": "Path not found"})
	})

	port := ":8080"
	log.Printf("[UltraForward] 服务已就绪，请访问: http://0.0.0.0%s", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}

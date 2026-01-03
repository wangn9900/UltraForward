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
	database.InitDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// CORS 设置
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 获取嵌入式前端
	publicFS, err := fs.Sub(assets.WebDist, "web-dist")
	if err != nil {
		log.Fatal("Failed to load embedded frontend:", err)
	}

	// API 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", api.RegisterHandler)
		v1.POST("/auth/login", api.LoginHandler)
		v1.GET("/stats", api.GetDashboardStats)
		v1.GET("/plans", api.GetPlans)
		v1.POST("/admin/plans", api.CreatePlan)
		v1.DELETE("/admin/plans/:id", api.DeletePlan)
	}

	// 统一处理前端资源 (SPA 路由适配)
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果是请求具体资源文件 (assets, images 等)
		if strings.Contains(path, ".") || strings.HasPrefix(path, "/assets/") {
			http.FileServer(http.FS(publicFS)).ServeHTTP(c.Writer, c.Request)
			return
		}

		// 否则，返回 index.html 以支持 Vue Router
		indexData, err := fs.ReadFile(publicFS, "index.html")
		if err != nil {
			c.String(404, "Frontend fully missing. Please check build.")
			return
		}
		c.Data(200, "text/html; charset=utf-8", indexData)
	})

	port := ":8080"
	log.Printf("[UltraForward] 旗舰控制中心已上线: http://0.0.0.0%s", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}

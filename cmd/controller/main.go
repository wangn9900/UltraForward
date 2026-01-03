package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/api"
	"github.com/wangn9900/UltraForward/internal/database"
)

//go:embed all:web-dist
var embeddedFiles embed.FS

func main() {
	database.InitDB()

	// 设置为发布模式，减少冗余日志
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// CORS 跨域支持
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

	// 获取嵌入式前端文件系统
	publicFS, err := fs.Sub(embeddedFiles, "web-dist")
	if err != nil {
		log.Fatal("Failed to load embedded frontend:", err)
	}

	// 注册 API 接口
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", api.RegisterHandler)
		v1.POST("/auth/login", api.LoginHandler)
		v1.GET("/stats", api.GetDashboardStats)
		v1.GET("/plans", api.GetPlans)

		// Admin
		v1.POST("/admin/plans", api.CreatePlan)
		v1.DELETE("/admin/plans/:id", api.DeletePlan)
	}

	// 智能静态资源分发 (支持 Vue SPA 路由)
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果请求的是带后缀的文件 (css, js, png等)，由 FileServer 提供
		if strings.Contains(path, ".") {
			http.FileServer(http.FS(publicFS)).ServeHTTP(c.Writer, c.Request)
			return
		}

		// 否则，其余所有路径全部返回 index.html (SPA 路由需求)
		indexData, _ := fs.ReadFile(publicFS, "index.html")
		c.Data(200, "text/html; charset=utf-8", indexData)
	})

	port := ":8080"
	log.Printf("[UltraForward] 旗舰控制中心已上线: http://0.0.0.0%s", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}

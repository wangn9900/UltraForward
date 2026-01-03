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

	// CORS
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

	// 使用新的 assets 包加载嵌入网页
	publicFS, err := fs.Sub(assets.WebDist, "web-dist")
	if err != nil {
		log.Fatal("Failed to load embedded frontend:", err)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", api.RegisterHandler)
		v1.POST("/auth/login", api.LoginHandler)
		v1.GET("/stats", api.GetDashboardStats)
		v1.GET("/plans", api.GetPlans)
		v1.POST("/admin/plans", api.CreatePlan)
		v1.DELETE("/admin/plans/:id", api.DeletePlan)
	}

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.Contains(path, ".") {
			http.FileServer(http.FS(publicFS)).ServeHTTP(c.Writer, c.Request)
			return
		}
		indexData, _ := fs.ReadFile(publicFS, "index.html")
		c.Data(200, "text/html; charset=utf-8", indexData)
	})

	port := ":8080"
	log.Printf("[UltraForward] 旗舰控制中心已上线: http://0.0.0.0%s", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}

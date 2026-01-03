package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/api"
	"github.com/wangn9900/UltraForward/internal/database"
)

func main() {
	database.InitDB()
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// 静态资源路由 (SPA 模式)
	webRoot := "./web-dist"
	if _, err := os.Stat(filepath.Join(webRoot, "index.html")); err == nil {
		r.Static("/assets", filepath.Join(webRoot, "assets"))
		r.StaticFile("/", filepath.Join(webRoot, "index.html"))
		r.NoRoute(func(c *gin.Context) {
			if !strings.HasPrefix(c.Request.URL.Path, "/api/v1") {
				c.File(filepath.Join(webRoot, "index.html"))
				return
			}
			c.Status(404)
		})
	}

	// API 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("/stats", api.GetDashboardStats)
		v1.GET("/rules", api.GetUserRules)
		v1.GET("/lines", api.GetLines)
		v1.GET("/plans", api.GetPlans)
	}

	log.Println("[UltraForward Controller] Starting on :8080...")
	r.Run(":8080")
}

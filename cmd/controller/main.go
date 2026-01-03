package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/api"
	"github.com/wangn9900/UltraForward/internal/database"
)

func main() {
	// 初始化数据库
	database.InitDB()

	r := gin.Default()

	// 允许跨域 (开发环境)
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

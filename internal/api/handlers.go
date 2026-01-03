package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/database"
	"github.com/wangn9900/UltraForward/internal/models"
)

// GetDashboardStats 获取商用面板概览
func GetDashboardStats(c *gin.Context) {
	// 模拟当前用户 ID 为 1
	var user models.User
	database.DB.First(&user, 1)

	var ruleCount int64
	database.DB.Model(&models.Rule{}).Where("user_id = ?", 1).Count(&ruleCount)

	c.JSON(http.StatusOK, gin.H{
		"user":         user,
		"active_rules": ruleCount,
	})
}

// GetUserRules 获取用户的转发规则
func GetUserRules(c *gin.Context) {
	var rules []models.Rule
	database.DB.Preload("Line").Where("user_id = ?", 1).Find(&rules)
	c.JSON(http.StatusOK, rules)
}

// GetLines 获取所有可用线路 (商用名称)
func GetLines(c *gin.Context) {
	var lines []models.Line
	database.DB.Find(&lines)
	c.JSON(http.StatusOK, lines)
}

// GetPlans 获取商城套餐
func GetPlans(c *gin.Context) {
	var plans []models.Plan
	database.DB.Find(&plans)
	c.JSON(http.StatusOK, plans)
}

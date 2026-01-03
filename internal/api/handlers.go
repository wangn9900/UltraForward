package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/database"
	"github.com/wangn9900/UltraForward/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterHandler 处理用户注册
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 检查用户是否已存在
	var exists models.User
	if err := database.DB.Where("username = ?", req.Username).First(&exists).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "该用户名已被注册"})
		return
	}

	// 密码哈希
	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	// 创建新用户 (初始化默认值)
	user := models.User{
		Username:     req.Username,
		Password:     string(hashed),
		Balance:      0,
		TotalTraffic: 0,
		UsedTraffic:  0,
		MaxRules:     3,
		ExpiredAt:    time.Now().AddDate(0, 0, 7), // 赠送 7 天试用
		AutoRenewal:  false,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败，请稍后再试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功", "user_id": user.ID})
}

// LoginHandler 处理用户登录
func LoginHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	// TODO: 生成 JWT Token
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"balance":  user.Balance,
		},
		"token": "simulated_jwt_token_for_now",
	})
}

// GetDashboardStats 获取商用面板概览
func GetDashboardStats(c *gin.Context) {
	// 获取当前登录用户 (此处暂硬编码 ID 1)
	var user models.User
	database.DB.First(&user, 1)

	var ruleCount int64
	database.DB.Model(&models.Rule{}).Where("user_id = ?", user.ID).Count(&ruleCount)

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

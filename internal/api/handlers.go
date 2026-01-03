package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wangn9900/UltraForward/internal/database"
	"github.com/wangn9900/UltraForward/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// --- 用户 Auth 相关 ---

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	var exists models.User
	if err := database.DB.Where("username = ?", req.Username).First(&exists).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "该用户名已被注册"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	user := models.User{
		Username:  req.Username,
		Password:  string(hashed),
		Balance:   0,
		ExpiredAt: time.Now().AddDate(0, 0, 7),
		MaxRules:  3,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "user": user, "token": "simulated_token"})
}

// --- 用户业务接口 ---

func GetDashboardStats(c *gin.Context) {
	var user models.User
	database.DB.First(&user)
	c.JSON(http.StatusOK, gin.H{"user": user, "active_rules": 0})
}

func GetPlans(c *gin.Context) {
	var plans []models.Plan
	database.DB.Find(&plans)
	c.JSON(http.StatusOK, plans)
}

// --- 管理员接口 (Admin) ---

// CreatePlan 管理员创建套餐
func CreatePlan(c *gin.Context) {
	var plan models.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&plan)
	c.JSON(http.StatusOK, gin.H{"message": "套餐创建成功", "plan": plan})
}

func DeletePlan(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Plan{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "套餐已删除"})
}

// CRUD for Lines and Nodes (后续可以继续补充)

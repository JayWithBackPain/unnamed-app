package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qimen/backend/internal/models"
	"github.com/qimen/backend/pkg/auth"
	"github.com/qimen/backend/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required,min=6"`
	Email     string    `json:"email" binding:"required,email"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserHandler 用戶處理器
type UserHandler struct {
	// TODO: 添加需要的服務
}

// NewUserHandler 創建用戶處理器
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Register 處理用戶註冊
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求數據"})
		return
	}

	// 檢查用戶名是否已存在
	var existingUser models.User
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用戶名已存在"})
		return
	}

	// 加密密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密碼加密失敗"})
		return
	}

	user := models.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		Email:     req.Email,
		BirthDate: req.BirthDate,
		Balance:   100, // 新用戶贈送 100 虛擬貨幣
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "創建用戶失敗"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "註冊成功"})
}

// Login 處理用戶登入
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的請求數據"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用戶名或密碼錯誤"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用戶名或密碼錯誤"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"balance":    user.Balance,
			"birth_date": user.BirthDate,
		},
	})
}

// GetProfile 獲取用戶資料
func (h *UserHandler) GetProfile(c *gin.Context) {
	// TODO: 實現獲取資料邏輯
}

// UpdateProfile 更新用戶資料
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	// TODO: 實現更新資料邏輯
}

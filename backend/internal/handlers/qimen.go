package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qimen/backend/internal/services"
)

// QimenHandler 奇門盤處理器
type QimenHandler struct {
	qimenService *services.QimenService
}

// NewQimenHandler 創建奇門盤處理器
func NewQimenHandler() *QimenHandler {
	return &QimenHandler{
		qimenService: services.NewQimenService(),
	}
}

// GenerateChart 生成奇門盤
func (h *QimenHandler) GenerateChart(c *gin.Context) {
	// 從上下文中獲取用戶ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "未授權",
		})
		return
	}

	// 獲取請求參數
	var req struct {
		DateTime string `json:"date_time" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "無效的請求參數",
		})
		return
	}

	// 解析時間
	dateTime, err := time.Parse(time.RFC3339, req.DateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "無效的時間格式",
		})
		return
	}

	// 生成奇門盤
	chart, err := h.qimenService.CalculateChart(userID.(uint), dateTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "生成奇門盤失敗",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": chart,
	})
}

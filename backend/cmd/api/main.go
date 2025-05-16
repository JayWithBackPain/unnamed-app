package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/qimen/backend/internal/handlers"
	"github.com/qimen/backend/internal/middleware"
	"github.com/qimen/backend/pkg/database"
)

func main() {
	// 載入環境變數
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 設置 Gin 模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化資料庫
	database.InitDB()

	// 初始化路由
	r := gin.Default()

	// 設置 CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 創建處理器
	userHandler := handlers.NewUserHandler()
	qimenHandler := handlers.NewQimenHandler()

	// 公開路由
	public := r.Group("/api")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	// 需要認證的路由
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		// 用戶相關
		authorized.GET("/user/profile", userHandler.GetProfile)
		authorized.PUT("/user/profile", userHandler.UpdateProfile)

		// 奇門盤相關
		authorized.POST("/qimen/chart", qimenHandler.GenerateChart)
	}

	// 健康檢查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 啟動服務器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

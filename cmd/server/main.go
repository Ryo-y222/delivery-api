package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryo-y222/delivery-api/internal/handler"
	"github.com/ryo-y222/delivery-api/internal/middleware"
	"github.com/ryo-y222/delivery-api/internal/model"
	"github.com/ryo-y222/delivery-api/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	const version = "0.1.0"
	startedAt := time.Now()

	// 環境変数からDB接続情報を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// MySQL接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ DB接続に失敗しました:", err)
	}

	log.Println("✅ DB接続成功！")

	//AutoMigrate
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("❌ マイグレーション失敗:", err)
	}
	log.Println("✅ マイグレーション完了！")

	// Repojitory
	userRepo := repository.NewUserRepository(db)
	// Handler
	authHandler := handler.NewAuthHandler(userRepo)

	// Ginルーター作成
	r := gin.Default()

	//CORSミドルウェア
	r.Use(middleware.CORSMiddleware())

	//ヘルスチェック
	healthHandler := handler.NewHealthHandler(db, version, startedAt)
	r.GET("/health", healthHandler.Check)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Delivery API!",
		})
	})

	//認証ルート
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", authHandler.Register)
	}

	// ポート8080で起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	_ = db // 後のIssueで使う
	r.Run(":" + port)

}

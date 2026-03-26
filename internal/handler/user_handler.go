package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-y222/delivery-api/internal/repository"
)

// UserHandler はユーザー関連のリクエストを処理する
type UserHandler struct {
	userRepo *repository.UserRepository
}

// NewUserHandler は新しい UserHandler を生成する
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// GetMe は認証済みユーザーの自分のプロフィールを返す
func (h *UserHandler) GetMe(c *gin.Context) {
	// ミドルウェアがセットしたuser_idを取得
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が見つかりません"})
		return
	}
	// uintに型変換
	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDの変換に失敗しました。"})
		return
	}
	//DBからユーザーを取得
	user, err := h.userRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースエラー"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません。"})
		return
	}
	// レスポンスを返す
	c.JSON(http.StatusOK, gin.H{"user": user})

}

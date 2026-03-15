package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryo-y222/delivery-api/internal/model"
	"github.com/ryo-y222/delivery-api/internal/repository"
	"github.com/ryo-y222/delivery-api/internal/util"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		UserRepo: userRepo,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=shipper transport_company"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	//1.リクエストボディをパース、バリデーション
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//2.メールアドレスの重複チェック
	existingUser, err := h.UserRepo.GetByEmail(req.Email) //&user,errが返る
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースエラー"})
		return
	}

	if existingUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "このメールアドレスは既に登録されています。"})
		return
	}

	// 3. パスワードをハッシュ化する
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "パスワードのハッシュ化に失敗しました"})
		return
	}

	// 4. ユーザーを作成する
	user := &model.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Name:         req.Name,
		Role:         req.Role,
		Company:      req.Company,
		Phone:        req.Phone,
	}

	if err := h.UserRepo.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーの作成に失敗しました。"})
		return
	}

	// 5. JWTトークンを生成する
	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := util.GenerateToken(user.ID, user.Email, user.Role, jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}

	// 6. HTTP-only Cookie にトークンをセットする
	c.SetCookie(
		"auth_token", // Cookie名
		token,        // 値（JWTトークン）
		86400,        // 有効期限（秒）= 24時間
		"/",          // パス（全てのパスで有効）
		"",           // ドメイン（空 = 現在のドメイン）
		false,        // Secure（HTTPS限定）※ 開発中は false、本番は true
		true,         // HttpOnly ⭐ JavaScriptからアクセス不可
	)

	//7.レスポンスを返す(トークンはcookieに入っているのでボディには含めない。)
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

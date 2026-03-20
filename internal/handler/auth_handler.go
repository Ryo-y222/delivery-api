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
	userRepo *repository.UserRepository
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
	}
}

// ユーザー登録リクエストの構造体
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=shipper transport_company"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
}

// ログインリクエストの構造体
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// クッキーセットのヘルパー関数
func setAuthCookie(c *gin.Context, token string) {
	c.SetCookie(
		"auth_token", // Cookie名
		token,        // 値（JWTトークン）
		86400,        // 有効期限（秒）= 24時間
		// 300,
		"/",   // パス（全てのパスで有効）
		"",    // ドメイン（空 = 現在のドメイン）
		false, // Secure（HTTPS限定）※ 開発中は false、本番は true
		true,  // HttpOnly ⭐ JavaScriptからアクセス不可
	)
}

func (h *AuthHandler) Register(c *gin.Context) {
	//1.リクエストボディをパース、バリデーション
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//2.メールアドレスの重複チェック
	existingUser, err := h.userRepo.GetByEmail(req.Email) //&user,errが返る
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

	if err := h.userRepo.Create(user); err != nil {
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
	setAuthCookie(c, token)

	//7.レスポンスを返す(トークンはcookieに入っているのでボディには含めない。)
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

// ログイン処理
func (h *AuthHandler) Login(c *gin.Context) {
	// 1. リクエストボディをパース→構造体変換
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. メールアドレスでユーザーを検索する
	user, err := h.userRepo.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースエラー"})
		return
	}
	if user == nil {
		// ⭐ 「メールが見つからない」とは言わない（セキュリティ上の理由）
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが正しくありません"})
		return
	}

	// 3. パスワードを検証する
	if err := util.CheckPassword(req.Password, user.PasswordHash); err != nil {
		// ⭐ 「パスワードが違う」とは言わない（セキュリティ上の理由）
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが正しくありません"})
		return
	}

	// 4. JWTトークンを生成する
	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := util.GenerateToken(user.ID, user.Email, user.Role, jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}

	// 5. HTTP-only Cookie にトークンをセットする
	setAuthCookie(c, token)

	// 6. レスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

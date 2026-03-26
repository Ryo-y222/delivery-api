package handler

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryo-y222/delivery-api/internal/model"
	"github.com/ryo-y222/delivery-api/internal/service"
)

type AuthHandler struct {
	authService  *service.AuthService
	secureCookie bool
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		secureCookie: os.Getenv("ENV") == "production",
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=72"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=shipper transport_company"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func toUserResponse(u *model.User) UserResponse {
	return UserResponse{ID: u.ID, Email: u.Email, Name: u.Name, Role: u.Role}
}

func setAuthCookie(c *gin.Context, token string, secure bool) {
	c.SetCookie(
		"auth_token",
		token,
		86400,
		"/",
		"",
		secure,
		true,
	)
}

func (h *AuthHandler) Register(c *gin.Context) {
	// 1.リクエストをパース
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力内容に不備があります。"})
		return
	}

	user, token, err := h.authService.Register(req.Email, req.Password, req.Name, req.Role, req.Company, req.Phone)

	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		log.Printf("[ERROR] Register failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラー"})
		return
	}

	setAuthCookie(c, token, h.secureCookie)

	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

func (h *AuthHandler) Login(c *gin.Context) {

	// 1.リクエストをパース
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力内容に不備があります。"})
		return
	}
	user, token, err := h.authService.Login(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		log.Printf("[ERROR] Login failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラー"})
		return
	}

	setAuthCookie(c, token, h.secureCookie)

	c.JSON(http.StatusOK, gin.H{"user": toUserResponse(user)})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie(
		"auth_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "ログアウト成功"})
}

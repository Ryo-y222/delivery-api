package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-y222/delivery-api/internal/util"
)

// AuthMiddleware はCookieからJWTを読み取り、認証を検証するミドルウェア
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//クッキーからJWTを取り出す。
		tokenString, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です。"})
			c.Abort()
			return
		}
		//トークンを検証
		claims, err := util.ParseToken(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "トークンが無効です。"})
			c.Abort()
			return
		}
		// トークンからユーザー情報を取り出す。ハンドラーに渡して使う。
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RoleMiddleware は指定されたロールのみアクセスを許可するミドルウェア
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//ロールを取得
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です。"})
			c.Abort()
			return
		}

		// 比較するために文字列に、interfaceとの比較
		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ローカル情報の取得に失敗しました。"})
			c.Abort()
			return
		}

		//スライスの中に含まれているかチェック
		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "この操作を行う権限がありません。"})
		c.Abort()

	}
}

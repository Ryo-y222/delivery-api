package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db        *gorm.DB
	startedAt time.Time
	version   string
	now       func() time.Time //
}

func NewHealthHandler(db *gorm.DB, version string, startedAt time.Time) *HealthHandler {
	return &HealthHandler{
		db:        db,
		version:   version,
		startedAt: startedAt,
		now:       time.Now,
	}
}

// APIサーバーとDBの状態を返す
func (h *HealthHandler) Check(c *gin.Context) {
	uptimeSeconds := int64(h.now().Sub(h.startedAt).Seconds())
	//
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":         "unhealthy",
			"database":       "disconnected",
			"version":        h.version,
			"uptime_seconds": uptimeSeconds,
			"error":          err.Error(),
		})
		return
	}

	//Pingチェック
	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":         "unhealthy",
			"database":       "unreachable",
			"version":        h.version,
			"uptime_seconds": uptimeSeconds,
			"error":          err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         "healthy",
		"database":       "connected",
		"version":        h.version,
		"uptime_seconds": uptimeSeconds,
	})
}

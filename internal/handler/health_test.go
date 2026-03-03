package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm + sqlmock のDBを作る（Ping期待は一切しない）
func setupMockGormDB(t *testing.T) (*gorm.DB, func(), func()) {
	t.Helper()

	sqlDB, _, err := sqlmock.New() // MonitorPingsOption は使わない
	if err != nil {
		t.Fatalf("sqlmock.New failed: %v", err)
	}

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open failed: %v", err)
	}

	cleanup := func() { _ = sqlDB.Close() }
	// Ping失敗を作りたい時用（Closeすると Ping が失敗する）
	closeDB := func() { _ = sqlDB.Close() }

	return gdb, cleanup, closeDB
}

func TestHealth_Check_Healthy(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, cleanup, _ := setupMockGormDB(t)
	defer cleanup()

	h := NewHealthHandler(db)

	r := gin.New()
	r.GET("/health", h.Check)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}

	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("json unmarshal failed: %v", err)
	}

	if body["status"] != "healthy" {
		t.Fatalf("status=%v", body["status"])
	}

	// ✅ あなたのcurl結果に合わせてチェック
	if body["database"] != "connected" {
		t.Fatalf("database=%v", body["database"])
	}

	// version/uptime_seconds が入っていることだけ確認（値は環境により変わるので固定しない）
	if v, ok := body["version"].(string); !ok || v == "" {
		t.Fatalf("version missing or empty: %v", body["version"])
	}
	if u, ok := body["uptime_seconds"].(float64); !ok || u < 0 {
		t.Fatalf("uptime_seconds missing or invalid: %v", body["uptime_seconds"])
	}
}

func TestHealth_Check_Unhealthy_WhenPingFails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, cleanup, closeDB := setupMockGormDB(t)
	defer cleanup()

	// Ping を失敗させる（DBをCloseすると Ping が connection refused 相当で失敗する）
	closeDB()

	h := NewHealthHandler(db)

	r := gin.New()
	r.GET("/health", h.Check)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}

	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("json unmarshal failed: %v", err)
	}

	if body["status"] != "unhealthy" {
		t.Fatalf("status=%v", body["status"])
	}

	// ここは実装により "unreachable" / "disconnected" どちらもあり得るので許容
	// ※あなたのコードでは Ping失敗時は "unreachable" のはず
	dbVal, _ := body["database"].(string)
	if dbVal != "unreachable" && dbVal != "disconnected" {
		// もしここで落ちたら、health.go のキーが "databese" のtypoになってる可能性が高いです
		t.Fatalf("database=%v (maybe typo key 'databese'?)", body["database"])
	}

	if v, ok := body["version"].(string); !ok || v == "" {
		t.Fatalf("version missing or empty: %v", body["version"])
	}
	if u, ok := body["uptime_seconds"].(float64); !ok || u < 0 {
		t.Fatalf("uptime_seconds missing or invalid: %v", body["uptime_seconds"])
	}

	// エラー文が返ること（中身は環境で変わるので空じゃないことだけ）
	if e, ok := body["error"].(string); !ok || e == "" {
		t.Fatalf("error missing or empty: %v", body["error"])
	}
}

package seed

import (
	"log"

	"github.com/ryo-y222/delivery-api/internal/model"
	"gorm.io/gorm"
)

// SeedBlockedPatterns は初期のフィルタリングパターンを投入する
func SeedBlockedPatterns(db *gorm.DB) {
	patterns := []model.BlockedPattern{
		{Pattern: `\d{2,4}-\d{2,4}-\d{4}`, Description: "電話番号（ハイフンあり）", IsActive: true},
		{Pattern: `0\d{9,10}`, Description: "電話番号（ハイフンなし）", IsActive: true},
		{Pattern: `[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+`, Description: "メールアドレス", IsActive: true},
		{Pattern: `https?://\S+`, Description: "URL", IsActive: true},
	}

	for _, p := range patterns {
		// 既に同じパターンが存在する場合はスキップ
		var count int64
		db.Model(&model.BlockedPattern{}).Where("pattern = ?", p.Pattern).Count(&count)
		if count == 0 {
			db.Create(&p)
			log.Printf("✅ パターン追加: %s", p.Description)
		}
	}
}

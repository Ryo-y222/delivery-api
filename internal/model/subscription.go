package model

import "time"

// Subscription plan constants
const (
	PlanFree     = "free"     // 無料: ソロモードのみ、月3件マッチング
	PlanStandard = "standard" // ¥9,800/月: マッチング無制限、チャット
	PlanPremium  = "premium"  // ¥29,800/月: 全機能、API連携、優先表示
)

// Subscription status constants
const (
	SubscriptionStatusActive    = "active"
	SubscriptionStatusCancelled = "cancelled"
	SubscriptionStatusPastDue   = "past_due"
)

// Subscription は運送会社の月額サブスクリプションを表す
type Subscription struct {
	ID                   uint      `json:"id" gorm:"primaryKey"`
	CompanyID            uint      `json:"company_id" gorm:"index:idx_subscriptions_company"`
	Company              Company   `json:"company" gorm:"foreignKey:CompanyID" binding:"-"`
	Plan                 string    `json:"plan" gorm:"type:varchar(20);default:free;not null"`
	StripeSubscriptionID string    `json:"stripe_subscription_id" gorm:"type:varchar(255)"`
	Status               string    `json:"status" gorm:"type:varchar(20);default:active;not null"`
	CurrentPeriodStart   time.Time `json:"current_period_start"`
	CurrentPeriodEnd     time.Time `json:"current_period_end"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}

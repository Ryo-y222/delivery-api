package model

import "time"

// Review はマッチング完了後の相互評価を表す
type Review struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	MatchID           uint      `json:"match_id" gorm:"uniqueIndex:idx_review_unique,priority:1"`
	Match             Match     `json:"match" gorm:"foreignKey:MatchID" binding:"-"`
	ReviewerID        uint      `json:"reviewer_id" gorm:"uniqueIndex:idx_review_unique,priority:2"`
	Reviewer          User      `json:"reviewer" gorm:"foreignKey:ReviewerID" binding:"-"`
	RevieweeCompanyID uint      `json:"reviewee_company_id" gorm:"index:idx_reviews_company"`
	RevieweeCompany   Company   `json:"reviewee_company" gorm:"foreignKey:RevieweeCompanyID" binding:"-"`
	Rating            int       `json:"rating" binding:"required,min=1,max=5"`
	Comment           string    `json:"comment"`
	CreatedAt         time.Time `json:"created_at"`
}

func (Review) TableName() string {
	return "reviews"
}

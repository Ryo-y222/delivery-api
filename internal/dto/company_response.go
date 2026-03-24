package dto

// CompanySearchResponse は検索時の会社情報（最小限）
type CompanySearchResponse struct {
	DisplayName string `json:"display_name"`
	// 電話番号・正式名称・住所は含めない
}

// CompanyMatchedResponse はマッチング申請後の会社情報
type CompanyMatchedResponse struct {
	DisplayName string  `json:"display_name"`
	RatingAvg   float64 `json:"rating_avg"`
	TotalDeals  int     `json:"total_deals"`
}

// CompanyApprovedResponse は承認後の会社情報
type CompanyApprovedResponse struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	RatingAvg   float64 `json:"rating_avg"`
	TotalDeals  int     `json:"total_deals"`
}

// CompanyFullResponse は決済完了後の会社情報（フル開示）
type CompanyFullResponse struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	RatingAvg   float64 `json:"rating_avg"`
	TotalDeals  int     `json:"total_deals"`
}

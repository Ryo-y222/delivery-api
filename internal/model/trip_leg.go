package model

import "time"

// TripLeg cargo status constants
const (
	CargoStatusLoaded       = "loaded"        // 荷物あり（自社案件）
	CargoStatusEmptySeeking = "empty_seeking" // 空車・荷物募集中
	CargoStatusEmptyPrivate = "empty_private" // 空車・募集しない
)

// TripLeg visibility constants
const (
	VisibilityPrivate     = "private"      // 自社のみ閲覧可
	VisibilityCompanyOnly = "company_only" // 運送会社のみに公開
	VisibilityPublic      = "public"       // 全ユーザーに公開
)

// TripLeg status constants
const (
	TripLegStatusScheduled = "scheduled"
	TripLegStatusInTransit = "in_transit"
	TripLegStatusCompleted = "completed"
	TripLegStatusCancelled = "cancelled"
)

// TripLeg leg type constants
const (
	LegTypeOutbound = "outbound" // 往路
	LegTypeReturn   = "return"   // 帰路
)

// TripLeg は配車計画を構成する1つの区間を表す
type TripLeg struct {
	ID               uint         `json:"id" gorm:"primaryKey"`
	DispatchPlanID   uint         `json:"dispatch_plan_id" gorm:"index:idx_trip_legs_dispatch_plan"`
	DispatchPlan     DispatchPlan `json:"dispatch_plan" gorm:"foreignKey:DispatchPlanID" binding:"-"`
	LegOrder         int          `json:"leg_order"`
	OriginAddress    string       `json:"origin_address" binding:"required"`
	OriginLat        float64      `json:"origin_lat"`
	OriginLng        float64      `json:"origin_lng"`
	DestAddress      string       `json:"dest_address" binding:"required"`
	DestLat          float64      `json:"dest_lat"`
	DestLng          float64      `json:"dest_lng"`
	DepartureAt      time.Time    `json:"departure_at" gorm:"index:idx_trip_legs_departure"`
	ArrivalAt        time.Time    `json:"arrival_at"`
	LegType          string       `json:"leg_type" gorm:"type:varchar(20);default:outbound"`
	CargoStatus      string       `json:"cargo_status" gorm:"type:varchar(20);default:loaded;index:idx_trip_legs_visibility_cargo"`
	CargoDescription string       `json:"cargo_description"`
	CargoWeight      float64      `json:"cargo_weight"`
	AvailableWeight  float64      `json:"available_weight"`
	Price            int          `json:"price"`
	Visibility       string       `json:"visibility" gorm:"type:varchar(20);default:private;index:idx_trip_legs_visibility_cargo"`
	RoutePolyline    string       `json:"route_polyline" gorm:"type:text"`
	RouteDurationSec int          `json:"route_duration_sec" gorm:"default:0"`
	RouteStepsJSON   string       `json:"route_steps_json" gorm:"type:mediumtext"`
	DelayMinutes     int          `json:"delay_minutes" gorm:"default:0"`
	Status           string       `json:"status" gorm:"type:varchar(20);default:scheduled;index:idx_trip_legs_status"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
}

func (TripLeg) TableName() string {
	return "trip_legs"
}

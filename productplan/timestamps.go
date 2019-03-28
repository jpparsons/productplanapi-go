package productplan

import "time"

// Timestamps represents a timestamp
type Timestamps struct {
	CreatedAt time.Time `json:"created_at,string"`
	UpdatedAt time.Time `json:"updated_at,string"`
}

package postarticle_model

import "time"

type Post struct {
	Id          int
	Title       string
	Content     string
	Category    string
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate *time.Time
	Status      string
}

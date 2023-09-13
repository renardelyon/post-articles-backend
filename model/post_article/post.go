package postarticle_model

import "time"

type Post struct {
	Id          int        `json:"-"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Category    string     `json:"category"`
	CreatedDate time.Time  `json:"-" gorm:"column:created_date;autoCreateTime"`
	UpdatedDate *time.Time `json:"-"`
	Status      string     `json:"status"`
}

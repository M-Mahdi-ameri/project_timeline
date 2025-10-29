package domain

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AuthorID  uint      `json:"author_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

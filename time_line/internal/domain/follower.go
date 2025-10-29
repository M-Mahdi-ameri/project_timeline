package domain

type Follower struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	FollowerID uint
}

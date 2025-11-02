package repository

import (
	"context"

	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"gorm.io/gorm"
)

type GormFollowerRepo struct {
	db *gorm.DB
}

func NewGormFollowerRepo(db *gorm.DB) *GormFollowerRepo {
	return &GormFollowerRepo{db: db}
}

func (r *GormFollowerRepo) Follow(ctx context.Context, userID, followerID uint) error {
	f := domain.Follower{UserID: userID, FollowerID: followerID}
	return r.db.WithContext(ctx).Create(&f).Error
}

func (r *GormFollowerRepo) Unfollow(ctx context.Context, userID, followerID uint) error {
	return r.db.WithContext(ctx).Where("follower_id = ? AND user_id = ?", followerID, userID).Delete(&domain.Follower{}).Error

}
func (r *GormFollowerRepo) GetFollowers(ctx context.Context, userID uint) ([]uint, error) {
	var followers []uint
	if err := r.db.WithContext(ctx).Model(&domain.Follower{}).Where("user_id = ?", userID).Pluck("follower_id", &followers).Error; err != nil {
		return nil, err
	}
	return followers, nil
}
func (r *GormFollowerRepo) GetFollowing(ctx context.Context, followerID uint) ([]uint, error) {
	var following []uint
	if err := r.db.WithContext(ctx).Model(&domain.Follower{}).Where("follower_id = ?", followerID).Pluck("user_id", &following).Error; err != nil {
		return nil, err
	}
	return following, nil
}

package repository

import (
	"context"

	"github.com/M-Mahdi-ameri/time_line/internal/domain"
	"gorm.io/gorm"
)

type GormPostRepo struct {
	db *gorm.DB
}

func NewGormPostRepo(db *gorm.DB) *GormPostRepo {
	return &GormPostRepo{db: db}
}

func (r *GormPostRepo) Create(ctx context.Context, post *domain.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *GormPostRepo) GetByIDp(ctx context.Context, id uint) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *GormPostRepo) GetPostByIDs(ctx context.Context, ids []uint) ([]domain.Post, error) {
	posts := make([]domain.Post, 0, len(ids))
	if len(ids) == 0 {
		return posts, nil
	}
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *GormPostRepo) GetPostsByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error) {
	var posts []domain.Post
	if err := r.db.WithContext(ctx).Where("author_id = ?", authorID).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil

}

func (r *GormPostRepo) Deletep(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Post{}, id).Error
}

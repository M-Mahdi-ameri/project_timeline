package domain

import (
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetByIDp(ctx context.Context, id uint) (*Post, error)
	GetPostByIDs(ctx context.Context, ids []uint) ([]Post, error)
	Deletep(ctx context.Context, id uint) error
	GetPostsByAuthor(ctx context.Context, authorID uint) ([]Post, error)
}

type UserRepository interface {
	GetByIDu(ctx context.Context, id uint) (*User, error)
	Deleteu(ctx context.Context, id uint) error
}

type FollowerRepository interface {
	Follow(ctx context.Context, followerID, userID uint) error
	Unfollow(ctx context.Context, followerID, userID uint) error
	GetFollowers(ctx context.Context, userID uint) ([]uint, error)
	GetFollowing(ctx context.Context, followerID uint) ([]uint, error)
}

package repository

import (
	"context"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *Post) (int64, error)
	DeletePost(ctx context.Context, postID int64) error
	GetPost(ctx context.Context, postID int64) (*Post, error)
	UpdatePost(ctx context.Context, post *Post) error
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *Comment) (int64, error)
	DeleteCommentByPostID(ctx context.Context, postID int64) error
	GetCommentsByPostID(ctx context.Context, postID int64) ([]Comment, error)
}

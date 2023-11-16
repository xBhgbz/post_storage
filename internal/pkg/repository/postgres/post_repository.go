package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"storage/internal/pkg/repository"
)

type PostRepository struct {
	db databaseConnector
}

func NewPostRepository(db databaseConnector) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetPost(ctx context.Context, postID int64) (*repository.Post, error) {
	var p repository.Post
	err := r.db.Get(ctx, &p, `SELECT id, text, likes, created_at FROM post WHERE id=$1;`, postID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrObjectNotFound
	}
	return &p, err
}

func (r *PostRepository) CreatePost(ctx context.Context, post *repository.Post) (int64, error) {
	row := r.db.ExecQueryRow(ctx, `INSERT INTO post (text, likes) VALUES ($1, $2) RETURNING id;`, post.Text, post.Likes)
	var ID int64
	if err := row.Scan(&ID); err != nil {
		return 0, err
	}
	return ID, nil
}

func (r *PostRepository) UpdatePost(ctx context.Context, post *repository.Post) error {
	tag, err := r.db.Exec(ctx, `UPDATE post SET text=$2, likes=$3 WHERE id=$1;`, post.ID, post.Text, post.Likes)
	if tag.RowsAffected() == 0 {
		return repository.ErrObjectNotFound
	}
	return err
}

func (r *PostRepository) DeletePost(ctx context.Context, postID int64) error {
	tag, err := r.db.Exec(ctx, "DELETE FROM post WHERE id=$1;", postID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return repository.ErrObjectNotFound
	}
	return nil
}

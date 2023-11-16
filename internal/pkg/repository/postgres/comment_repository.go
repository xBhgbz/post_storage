package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"storage/internal/pkg/repository"
)

type CommentRepository struct {
	db databaseConnector
}

func NewCommentRepository(db databaseConnector) *CommentRepository {
	return &CommentRepository{db: db}
}

func (c *CommentRepository) CreateComment(ctx context.Context, comment *repository.Comment) (int64, error) {
	row := c.db.ExecQueryRow(ctx, `INSERT INTO comment (text, likes, post_id) VALUES ($1, $2, $3) RETURNING id;`, comment.Text, comment.Likes, comment.PostID)
	var ID int64
	if err := row.Scan(&ID); err != nil {
		return 0, err
	}
	return ID, nil
}

func (c *CommentRepository) DeleteCommentByPostID(ctx context.Context, postID int64) error {
	_, err := c.db.Exec(ctx, "DELETE FROM comment WHERE post_id=$1;", postID)
	return err
}

func (c *CommentRepository) GetCommentsByPostID(ctx context.Context, postID int64) ([]repository.Comment, error) {
	var comments []repository.Comment
	err := c.db.Select(ctx, &comments, "SELECT id, text, likes, created_at, post_id FROM comment WHERE post_id=$1;", postID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return comments, nil
}

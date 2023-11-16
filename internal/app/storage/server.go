package storage

import (
	"storage/internal/pkg/repository"
	"storage/pkg/posting"
)

type Storage struct {
	posting.UnimplementedPostingServer
	post    repository.PostRepository
	comment repository.CommentRepository
}

func NewStorage(post repository.PostRepository, comment repository.CommentRepository) *Storage {
	return &Storage{post: post, comment: comment}
}

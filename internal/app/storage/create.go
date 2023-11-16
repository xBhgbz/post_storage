package storage

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"storage/internal/pkg/logger"
	"storage/internal/pkg/repository"
	"storage/pkg/posting"
)

func (s *Storage) CreatePost(ctx context.Context, req *posting.CreatePostRequest) (*posting.CreatePostResponse, error) {
	ctx = logger.ToContext(ctx, logger.FromContext(ctx).With(zap.String("method", "create")))
	span, ctx := opentracing.StartSpanFromContext(ctx, "creating post")
	defer span.Finish()

	post := MakeRepositoryPost(req.Post)

	addSpan, ctx := opentracing.StartSpanFromContext(ctx, "adding post")
	ID, err := s.post.CreatePost(ctx, post)
	addSpan.Finish()
	if err != nil {
		logger.Errorf(ctx, err.Error())
		return &posting.CreatePostResponse{}, err
	}

	comments := MakeRepositoryComments(req.Comments, ID)
	addSpan, ctx = opentracing.StartSpanFromContext(ctx, "adding comments")
	for i := range comments {
		_, err = s.comment.CreateComment(ctx, &comments[i])
		if err != nil {
			logger.Errorf(ctx, err.Error())
			return &posting.CreatePostResponse{}, err
		}
	}
	addSpan.Finish()

	logger.Infof(ctx, "creation success")
	return &posting.CreatePostResponse{PostId: ID}, nil
}

func MakeRepositoryPost(post *posting.Post) *repository.Post {
	return &repository.Post{
		Text:  post.GetText(),
		Likes: post.GetLikes(),
	}
}

func MakeRepositoryComments(comments []*posting.Comment, postID int64) []repository.Comment {
	newComments := make([]repository.Comment, len(comments))
	for i := range comments {
		newComments[i] = repository.Comment{
			Text:   comments[i].GetText(),
			Likes:  comments[i].GetLikes(),
			PostID: postID,
		}
	}
	return newComments
}

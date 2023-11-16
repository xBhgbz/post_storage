package storage

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"storage/internal/pkg/logger"
	"storage/internal/pkg/repository"
	"storage/pkg/posting"
)

func (s *Storage) GetPost(ctx context.Context, req *posting.GetPostRequest) (*posting.GetPostResponse, error) {
	ctx = logger.ToContext(ctx, logger.FromContext(ctx).With(zap.String("method", "get")).With(zap.Int64("ID", req.GetId())))
	span, ctx := opentracing.StartSpanFromContext(ctx, "getting post")
	defer span.Finish()

	searchSpan, ctx := opentracing.StartSpanFromContext(ctx, "searching for post in repository")
	post, err := s.post.GetPost(ctx, req.GetId())
	if errors.Is(err, repository.ErrObjectNotFound) {
		logger.Infof(ctx, "no post with such id")
		return &posting.GetPostResponse{}, posting.ErrObjectNotFound
	}
	if err != nil {
		logger.Errorf(ctx, err.Error())
		return &posting.GetPostResponse{}, err
	}
	searchSpan.Finish()

	searchSpan, ctx = opentracing.StartSpanFromContext(ctx, "searching for comments in repository")
	comments, err := s.comment.GetCommentsByPostID(ctx, req.GetId())
	if err != nil {
		logger.Errorf(ctx, err.Error())
		return &posting.GetPostResponse{}, err
	}
	searchSpan.Finish()

	logger.Infof(ctx, "get success")
	return &posting.GetPostResponse{
		Post:     MakeGRPCPost(post),
		Comments: MakeGRPCComments(comments),
	}, nil
}

func MakeGRPCComments(comments []repository.Comment) []*posting.Comment {
	responseComments := make([]*posting.Comment, len(comments))
	for i := range comments {
		responseComments[i] = &posting.Comment{
			Likes: comments[i].Likes,
			Text:  comments[i].Text,
		}
	}
	return responseComments
}

func MakeGRPCPost(post *repository.Post) *posting.Post {
	return &posting.Post{
		Likes: post.Likes,
		Text:  post.Text,
	}
}

package storage

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"storage/internal/pkg/logger"
	"storage/internal/pkg/repository"
	"storage/pkg/posting"
)

func (s *Storage) DeletePost(ctx context.Context, req *posting.DeletePostRequest) (*emptypb.Empty, error) {
	ctx = logger.ToContext(ctx, logger.FromContext(ctx).With(zap.String("method", "delete")).With(zap.Int64("ID", req.GetId())))
	span, ctx := opentracing.StartSpanFromContext(ctx, "delete post")
	defer span.Finish()

	deleteSpan, ctx := opentracing.StartSpanFromContext(ctx, "deleting comments")
	err := s.comment.DeleteCommentByPostID(ctx, req.GetId())
	if err != nil {
		logger.Errorf(ctx, err.Error())
		return &emptypb.Empty{}, err
	}
	deleteSpan.Finish()

	deleteSpan, ctx = opentracing.StartSpanFromContext(ctx, "deleting post")
	err = s.post.DeletePost(ctx, req.GetId())
	if errors.Is(err, repository.ErrObjectNotFound) {
		logger.Errorf(ctx, err.Error())
		return &emptypb.Empty{}, posting.ErrObjectNotFound
	}
	if err != nil {
		logger.Errorf(ctx, err.Error())
		return &emptypb.Empty{}, err
	}
	deleteSpan.Finish()

	logger.Infof(ctx, "delete success")
	return &emptypb.Empty{}, nil
}

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

func (s *Storage) UpdatePost(ctx context.Context, req *posting.UpdatePostRequest) (*emptypb.Empty, error) {
	ctx = logger.ToContext(ctx, logger.FromContext(ctx).With(zap.String("method", "update")).With(zap.Int64("ID", req.GetId())))
	span, ctx := opentracing.StartSpanFromContext(ctx, "getting post")
	defer span.Finish()

	updateSpan, ctx := opentracing.StartSpanFromContext(ctx, "updating post in repository")
	post := MakeRepositoryPost(req.Post)
	post.ID = req.GetId()
	err := s.post.UpdatePost(ctx, post)
	if errors.Is(err, repository.ErrObjectNotFound) {
		logger.Infof(ctx, "no post with such id")
		return &emptypb.Empty{}, posting.ErrObjectNotFound
	} else if err != nil {
		logger.Errorf(ctx, err.Error())
		return &emptypb.Empty{}, err
	}
	updateSpan.Finish()

	logger.Infof(ctx, "update success")
	return &emptypb.Empty{}, nil
}

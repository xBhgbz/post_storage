package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"storage/internal/app/storage"
	"storage/internal/app/tracing"
	"storage/internal/pkg/environment"
	"storage/internal/pkg/logger"
	"storage/internal/pkg/repository/postgres"
	"storage/internal/pkg/repository/postgres/pgconnector"
	"storage/pkg/posting"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := environment.LoadEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	portNumber := environment.GetAddr()

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	logger.SetGlobal(zapLogger)

	tracer := tracing.SetGlobalTracer("posting-service")
	defer tracer.Close()

	database, err := pgconnector.ConnectToPostgresDB(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Close()

	implementation := storage.NewStorage(postgres.NewPostRepository(database), postgres.NewCommentRepository(database))

	grpcServer := grpc.NewServer()
	posting.RegisterPostingServer(grpcServer, implementation)

	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Infof(ctx, "grpc server working")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

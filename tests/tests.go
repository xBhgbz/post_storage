//go:build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"storage/internal/pkg/environment"
	"storage/pkg/posting"
	"testing"
)

func clientSetUp() posting.PostingClient {
	port := environment.GetAddr()
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return posting.NewPostingClient(conn)
}

func TestCreateGetPost(t *testing.T) {
	defer db.Clear(t)
	client := clientSetUp()
	request := posting.CreatePostRequest{
		Post:     Post().Valid().Pointer(),
		Comments: []*posting.Comment{Comment().Valid().Pointer()},
	}

	createResponse, err := client.CreatePost(context.Background(), &request)
	getResponse, err2 := client.GetPost(context.Background(), &posting.GetPostRequest{Id: createResponse.GetPostId()})

	require.Nil(t, err)
	require.Nil(t, err2)

	assert.Equal(t, Post().Valid().Value().Text, getResponse.Post.GetText())
	assert.Equal(t, Post().Valid().Value().Likes, getResponse.Post.GetLikes())

	require.Equal(t, 1, len(getResponse.Comments))
	assert.Equal(t, Comment().Valid().Value().Text, getResponse.Comments[0].Text)
	assert.Equal(t, Comment().Valid().Value().Likes, getResponse.Comments[0].Likes)
}

func TestUpdatePost(t *testing.T) {
	defer db.Clear(t)
	client := clientSetUp()
	createRequest := posting.CreatePostRequest{
		Post:     Post().Text("Hello World").Likes(543).Pointer(),
		Comments: []*posting.Comment{},
	}
	updateRequest := posting.UpdatePostRequest{
		Post: Post().Text("World Hello").Likes(345).Pointer(),
		Id:   0,
	}

	createResponse, err := client.CreatePost(context.Background(), &createRequest)
	updateRequest.Id = createResponse.GetPostId()
	_, err1 := client.UpdatePost(context.Background(), &updateRequest)
	getResponse, err2 := client.GetPost(context.Background(), &posting.GetPostRequest{Id: createResponse.GetPostId()})

	require.Nil(t, err)
	require.Nil(t, err1)
	require.Nil(t, err2)

	assert.Equal(t, updateRequest.GetPost().GetText(), getResponse.Post.GetText())
	assert.Equal(t, updateRequest.GetPost().GetLikes(), getResponse.Post.GetLikes())
}

func TestDeletePost(t *testing.T) {
	defer db.Clear(t)
	client := clientSetUp()
	createRequest := posting.CreatePostRequest{
		Post:     Post().Text("Hello World").Likes(543).Pointer(),
		Comments: []*posting.Comment{},
	}

	createResponse, err := client.CreatePost(context.Background(), &createRequest)
	_, err1 := client.DeletePost(context.Background(), &posting.DeletePostRequest{Id: createResponse.GetPostId()})
	_, err2 := client.GetPost(context.Background(), &posting.GetPostRequest{Id: createResponse.GetPostId()})

	require.Nil(t, err)
	require.Nil(t, err1)
	require.NotNil(t, err2)
}

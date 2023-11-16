package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"storage/pkg/posting"
)

const port = ":9000"

func main() {
	ctx := context.Background()
	flag.Parse()

	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := posting.NewPostingClient(conn)

	switch flag.Arg(0) {
	case "create":
		var id *posting.CreatePostResponse
		id, err = client.CreatePost(ctx, &posting.CreatePostRequest{
			Post: &posting.Post{
				Likes: 228,
				Text:  "Hello world",
			},
			Comments: []*posting.Comment{{
				Likes: 2,
				Text:  "great",
			}, {
				Likes: 23,
				Text:  "lol",
			}},
		})
		fmt.Println(id.GetPostId())
	case "delete":
		_, err = client.DeletePost(ctx, &posting.DeletePostRequest{Id: 0})
	case "update":
		_, err = client.UpdatePost(ctx, &posting.UpdatePostRequest{Post: &posting.Post{
			Likes: 2000,
			Text:  "World hello",
		}, Id: 0})
	case "get":
		data, err := client.GetPost(ctx, &posting.GetPostRequest{Id: 0})
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Printf("Got data: %v\n", data)
		}
	}

	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print("Success")
	}
}

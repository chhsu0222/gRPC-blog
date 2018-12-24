package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chhsu0222/gRPC-blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// Create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "CH",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	cRes, err := c.CreateBlog(
		context.Background(),
		&blogpb.CreateBlogRequest{Blog: blog},
	)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", cRes)
}

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
	fmt.Printf("Blog has been created: %v\n", cRes)

	// Read Blog
	blogID := cRes.GetBlog().GetId()
	fmt.Println("Creating the blog")

	_, err = c.ReadBlog(
		context.Background(),
		&blogpb.ReadBlogRequest{BlogId: "5c20a997b8b15c59c7d987a3"},
	)
	if err != nil {
		fmt.Printf("Error happened while reading blog: %v\n", err)
	}

	rRes, err := c.ReadBlog(
		context.Background(),
		&blogpb.ReadBlogRequest{BlogId: blogID},
	)
	if err != nil {
		fmt.Printf("Error happened while reading blog: %v\n", err)
	}

	fmt.Printf("Blog was read: %v", rRes.GetBlog())
}

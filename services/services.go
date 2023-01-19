package services

import (
	"context"
	"fmt"
	"golang-mongo/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(ctx context.Context, client *mongo.Client) {
	collections := client.Database("blog").Collection("posts")

	docs := []interface{}{
		bson.D{{Key: "title", Value: "world"}, {Key: "body", Value: "Hello World"}},
		bson.D{{Key: "title", Value: "Mars"}, {Key: "body", Value: "Hello Mars"}},
		bson.D{{Key: "title", Value: "Pluto"}, {Key: "body", Value: "Hello Pluto"}},
	}

	res, err := collections.InsertMany(ctx, docs)
	if err != nil {
		log.Fatalln(err)
	}
	println(res.InsertedIDs)
}

func FindAll(ctx context.Context, client *mongo.Client, filter interface{}) {
	collections := client.Database("blog").Collection("posts")
	cur, err := collections.Find(ctx, filter)
	if err != nil {
		log.Fatalln(err)
	}
	defer cur.Close(ctx)

	var posts []models.Post
	err = cur.All(ctx, &posts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(posts)

}

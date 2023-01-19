package main

import (
	"context"
	"golang-mongo/database"
	"golang-mongo/services"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := database.MakeMongoDbConnection(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer client.Disconnect(ctx)

	services.CreatePost(ctx, client)
	services.FindAll(ctx, client, struct{}{})
}

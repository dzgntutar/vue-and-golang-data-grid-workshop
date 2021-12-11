package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	client, ctx, cancel, err := connectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer closeMongo(client, ctx, cancel)

	pingMongo(client, ctx)

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.SendString("Hello from fiber")
		return nil
	})

	app.Listen(":3000")
}

func connectMongo(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err
}

func pingMongo(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("mongo baglantisi basarili..")
	return nil
}

func closeMongo(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

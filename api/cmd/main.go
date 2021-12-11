package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	client, c, cancel, err := connectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer closeMongo(client, c, cancel)

	pingMongo(client, c)

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		var document interface{}

		document = bson.D{
			{"name", "Computer"},
			{"price", 80},
			{"count", 10},
			{"category", 1},
		}

		insertOneResult, err := insertOne(client, c, "myDb",
			"product", document)

		if err != nil {
			panic(err)
		}

		fmt.Println("InsertOne-->")
		fmt.Println(insertOneResult.InsertedID)

		var documents []interface{}

		documents = []interface{}{
			bson.D{
				{"name", "Phone"},
				{"price", 25},
				{"count", 5},
				{"category", 1},
			},
			bson.D{
				{"name", "Console"},
				{"price", 44},
				{"count", 2},
				{"category", 2},
			},
		}

		insertManyResult, err := insertMany(client, c, "myDb",
			"product", documents)

		if err != nil {
			panic(err)
		}

		fmt.Println("InsertMany-->")

		for id := range insertManyResult.InsertedIDs {
			fmt.Println(id)
		}

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
	fmt.Println("Mongo db bağlantısı başarılı..")
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

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

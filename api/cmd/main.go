package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongoSetting "vue-and-golang-data-grid-workshop/pkg/mongo"
	"vue-and-golang-data-grid-workshop/pkg/repository"
)

type MainApp struct {
	fiber  *fiber.App
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func main() {
	app := MainApp{}

	app.Initialize()

	app.CreateRoute()

	app.Run(":3000")
}

func (app *MainApp) Initialize() {

	client, c, cancel, err := mongoSetting.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	app.client = client
	app.ctx = c
	app.cancel = cancel

	if err := mongoSetting.PingMongo(client, c); err != nil {
		panic(err)
	}

	app.fiber = fiber.New()
}

func (app *MainApp) CreateRoute() {

	var productRepository = repository.ProductRepository{
		Client: app.client,
		Ctx:    app.ctx,
		Cancel: app.cancel,
	}

	app.fiber.Post("/", func(ctx *fiber.Ctx) error {

		var document interface{}

		document = bson.D{
			{"name", "Computer"},
			{"price", 80},
			{"count", 10},
			{"category", 1},
		}

		result, err := productRepository.InsertOne(document)

		if err != nil {
			fmt.Println("Hata-->")
			panic(err)
		}

		fmt.Println("InsertOne-->")
		fmt.Println(result.InsertedID)

		/*
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
			}*/

		return nil
	})
}

func (app *MainApp) Run(port string) {
	if err := app.fiber.Listen(port); err != nil {
		panic(err)
	}
}

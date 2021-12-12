package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"vue-and-golang-data-grid-workshop/pkg/mongo"
)

type MainApp struct {
	fiber *fiber.App
}

func main() {
	app := MainApp{}

	app.Initialize()

	app.CreateRoute()

	app.Run(":3000")
}

func (app *MainApp) Initialize() {

	client, c, cancel, err := mongo.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer mongo.CloseMongo(client, c, cancel)

	if err := mongo.PingMongo(client, c); err != nil {
		panic(err)
	}

	app.fiber = fiber.New()
}

func (app *MainApp) CreateRoute() {
	app.fiber.Get("/", func(ctx *fiber.Ctx) error {

		/*var document interface{}

		document = bson.D{
			{"name", "Computer"},
			{"price", 80},
			{"count", 10},
			{"category", 1},
		}

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
		}*/

		fmt.Println("hiiii")

		return nil
	})
}

func (app *MainApp) Run(port string) {
	if err := app.fiber.Listen(port); err != nil {
		panic(err)
	}
}

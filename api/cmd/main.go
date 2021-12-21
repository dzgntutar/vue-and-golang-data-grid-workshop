package main

import (
	"context"
	"fmt"
	"vue-and-golang-data-grid-workshop/pkg/model"
	mongoSetting "vue-and-golang-data-grid-workshop/pkg/mongo"
	"vue-and-golang-data-grid-workshop/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		product := model.Product{}
		if err := ctx.BodyParser(&product); err != nil {
			return err
		}

		var document interface{}

		document = bson.D{
			{"name", product.Name},
			{"price", product.Price},
			{"count", product.Count},
			{"category", product.Category},
		}

		result, err := productRepository.InsertOne(document)

		if err != nil {
			fmt.Println("Hata-->")
			panic(err)
		}

		fmt.Println("InsertOne-->")
		fmt.Println(result.InsertedID)

		return nil
	})

	app.fiber.Post("/insertMany", func(ctx *fiber.Ctx) error {

		var products []model.Product
		if err := ctx.BodyParser(&products); err != nil {
			return err
		}

		fmt.Println(products)

		documents := []bson.D{}

		for _, product := range products {
			documents = append(documents, bson.D{
				{"name", product.Name},
				{"price", product.Price},
				{"count", product.Count},
				{"category", product.Category},
			})
		}

		fmt.Println(documents)

		return nil
	})
}

func (app *MainApp) Run(port string) {
	if err := app.fiber.Listen(port); err != nil {
		panic(err)
	}
}

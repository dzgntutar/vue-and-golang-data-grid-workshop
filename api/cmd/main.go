package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"vue-and-golang-data-grid-workshop/pkg/model"
	"vue-and-golang-data-grid-workshop/pkg/repository"

	"github.com/gofiber/fiber/v2"
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

var productRepository repository.ProductRepository

func (app *MainApp) Initialize() {

	productRepository = repository.ProductRepository{}

	app.fiber = fiber.New()
}

func (app *MainApp) CreateRoute() {

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

		_, err := productRepository.InsertOne(document)

		if err != nil {
			panic(err)
		}

		return nil
	})

	app.fiber.Post("/insertMany", func(ctx *fiber.Ctx) error {

		var products []model.Product
		if err := ctx.BodyParser(&products); err != nil {
			return err
		}

		var documents []interface{}

		for _, product := range products {
			documents = append(documents, bson.D{
				{"name", product.Name},
				{"price", product.Price},
				{"count", product.Count},
				{"category", product.Category},
			})
		}

		_, err := productRepository.InsertMany(documents)

		if err != nil {
			panic(err)
		}

		return nil
	})

	app.fiber.Get("/", func(ctx *fiber.Ctx) error {

		products, err := productRepository.GetAll()

		if err != nil {
			panic(err)
		}
		if err := ctx.Status(fiber.StatusOK).JSON(products); err != nil {
			return err
		}

		return nil
	})

	app.fiber.Get("/getWithPagination", func(ctx *fiber.Ctx) error {

		var pageModel = model.PageModel{}

		if err := ctx.BodyParser(&pageModel); err != nil {
			return err
		}

		products, err := productRepository.GetAllWithPagination(pageModel)

		if err != nil {
			panic(err)
		}

		if err := ctx.Status(fiber.StatusOK).JSON(products); err != nil {
			return err
		}

		return nil
	})

}

func (app *MainApp) Run(port string) {
	if err := app.fiber.Listen(port); err != nil {
		panic(err)
	}
}

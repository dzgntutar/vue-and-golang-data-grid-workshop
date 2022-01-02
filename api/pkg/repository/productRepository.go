package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"vue-and-golang-data-grid-workshop/pkg/model"
	mongoSetting "vue-and-golang-data-grid-workshop/pkg/mongo"
)

var (
	dataBase = "myDb"
	col      = "product"
)

type ProductRepository struct {
}

func (r ProductRepository) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {
	client, ctx, cancel, err := mongoSetting.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	if err := mongoSetting.PingMongo(client, ctx); err != nil {
		fmt.Println("PingMongo")
		panic(err)
	}

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	collection := client.Database(dataBase).Collection(col)

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	result, err := collection.InsertOne(ctx, doc)

	return result, err
}

func (r ProductRepository) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
	client, ctx, cancel, err := mongoSetting.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	if err := mongoSetting.PingMongo(client, ctx); err != nil {
		fmt.Println("PingMongo")
		panic(err)
	}

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	collection := client.Database(dataBase).Collection(col)

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	result, err := collection.InsertMany(ctx, docs)

	return result, err
}

func (r ProductRepository) GetAllWithPagination(pageModel model.PageModel) ([]bson.M, error) {

	client, ctx, cancel, err := mongoSetting.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	if err := mongoSetting.PingMongo(client, ctx); err != nil {
		fmt.Println("PingMongo")
		panic(err)
	}

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	collection := client.Database(dataBase).Collection(col)

	skip := int64((pageModel.Page - 1) * pageModel.Count)
	limit := int64(pageModel.Count)

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := collection.Find(ctx, bson.D{}, &opts)

	var products []bson.M
	var product bson.M
	for cursor.Next(ctx) {

		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return products, nil

}

func (r ProductRepository) GetAll() ([]bson.M, error) {
	client, ctx, cancel, err := mongoSetting.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	if err := mongoSetting.PingMongo(client, ctx); err != nil {
		fmt.Println("PingMongo")
		panic(err)
	}

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	collection := client.Database(dataBase).Collection(col)

	defer mongoSetting.CloseMongo(client, ctx, cancel)

	cursor, err := collection.Find(ctx, bson.M{})

	var products []bson.M

	for cursor.Next(ctx) {
		var product bson.M
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
		fmt.Println(product)
	}
	cursor.Close(ctx)
	return products, nil
}

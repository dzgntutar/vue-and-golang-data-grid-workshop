package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	mongoSetting "vue-and-golang-data-grid-workshop/pkg/mongo"
)

var (
	dataBase = "myDb"
	col      = "product"
)

type ProductRepository struct {
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (r ProductRepository) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	defer mongoSetting.CloseMongo(r.Client, r.Ctx, r.Cancel)

	result, err := collection.InsertOne(r.Ctx, doc)

	return result, err
}

func (r ProductRepository) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	defer mongoSetting.CloseMongo(r.Client, r.Ctx, r.Cancel)

	result, err := collection.InsertMany(r.Ctx, docs)

	return result, err
}

func (r ProductRepository) GetAll() ([]bson.M, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	defer mongoSetting.CloseMongo(r.Client, r.Ctx, r.Cancel)

	cursor, err := collection.Find(r.Ctx, bson.M{})

	var products []bson.M

	for cursor.Next(r.Ctx) {
		var product bson.M
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		fmt.Println(product)
		products = append(products, product)
	}

	return products, nil

}

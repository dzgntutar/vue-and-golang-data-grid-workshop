package repository

import (
	"context"
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
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (r ProductRepository) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	//defer mongoSetting.CloseMongo(r.Client, r.Ctx, r.Cancel)

	result, err := collection.InsertOne(r.Ctx, doc)

	return result, err
}

func (r ProductRepository) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	//defer mongoSetting.CloseMongo(r.Client, r.Ctx, r.Cancel)

	result, err := collection.InsertMany(r.Ctx, docs)

	return result, err
}

func (r ProductRepository) GetAllWithPagination(pageModel model.PageModel) ([]bson.M, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	//defer mongoSetting.CloseMongo(r.Client, r.Ctx, r.Cancel)

	skip := int64((pageModel.Page - 1) * pageModel.Count)
	limit := int64(pageModel.Count)

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	//opts.SetSort(bson.D{{"_id", -1}})
	//son.D{{"price", bson.D{{"$gt", 50}}}}
	cursor, err := collection.Find(r.Ctx, bson.D{}, &opts)

	defer cursor.Close(r.Ctx)

	var products []bson.M

	for cursor.Next(r.Ctx) {
		var product bson.M
		if err = cursor.Decode(&product); err != nil {
			fmt.Println("In Cursor")
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return products, nil

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
		products = append(products, product)
	}

	return products, nil

}

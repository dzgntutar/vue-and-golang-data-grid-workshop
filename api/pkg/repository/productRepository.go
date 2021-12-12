package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dataBase = "myDb"
	col      = "product"
)

type ProductRepository struct {
	Client *mongo.Client
	Ctx    context.Context
}

func (r ProductRepository) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {
	collection := r.Client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(r.Ctx, doc)
	return result, err
}

func (r ProductRepository) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {

	collection := r.Client.Database(dataBase).Collection(col)

	result, err := collection.InsertMany(r.Ctx, docs)
	return result, err
}

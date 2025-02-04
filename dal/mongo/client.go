package mongo

import (
	"context"
	"fmt"
	"metadata/conf"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Database

func InitMongoDb() {
	dsn := fmt.Sprintf(conf.GetConfMongo().DbTemplate, conf.GetConfMongo().Username, conf.GetConfMongo().Passwd,
		conf.GetConfMongo().Host, conf.GetConfMongo().Port)

	println(dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
	mongoDB = client.Database(conf.GetConfMongo().Database)
}

func GetMongoDb() *mongo.Database {
	return mongoDB
}

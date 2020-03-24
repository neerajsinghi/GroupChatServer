package model

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataPoints struct {
	ScreenerName  string   `bson:"Screener_name" json:"Screener_name,omitempty"`
	Date          string   `json:"date,omitempty"`
	Dates         []string `json:"dates,omitempty"`
	Tradingsymbol string   `json:"tradingsymbol,omitempty"`
}

var collections *mongo.Collection

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	var err error
	//dbPort := os.Getenv("db_port")
	dbHost := os.Getenv("db_host")
	db := os.Getenv("db")
	dbName := os.Getenv("db_name")
	collName := os.Getenv("collection_name")
	userName := os.Getenv("user_name")
	password := os.Getenv("password")
	uri := db + "://" + userName + ":" + password + "@" + dbHost + "/test?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")
	//client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	//client.Connect(nil)
	collections = client.Database(dbName).Collection(collName)
}

func createCollation() *options.Collation {
	var collation options.Collation
	collation.Locale = "en"
	collation.CaseLevel = false
	return &collation
}

//FindPost find list of documents from db
func Find(filter, projection, sort bson.M, limit, skip int64) (*mongo.Cursor, error) {
	return collections.Find(nil, filter, options.Find().SetProjection(projection), options.Find().SetSort(sort), options.Find().SetSkip(skip), options.Find().SetLimit(limit))
}
func FindCase(filter, projection, sort bson.M, limit, skip int64) (*mongo.Cursor, error) {
	return collections.Find(nil, filter, options.Find().SetProjection(projection), options.Find().SetCollation(createCollation()), options.Find().SetSort(sort), options.Find().SetSkip(skip), options.Find().SetLimit(limit))
}

//FindPostOne find one  document from db
func FindOne(filter, projection, sort bson.M) *mongo.SingleResult {
	return collections.FindOne(nil, filter, options.FindOne().SetProjection(projection), options.FindOne().SetSort(sort))
}

//InsertMany Insert many documents to db
func InsertMany(document []interface{}) (*mongo.InsertManyResult, error) {
	return collections.InsertMany(nil, document)
}

//InsertOne Insert one document to db
func InsertOne(document bson.M) (*mongo.InsertOneResult, error) {
	return collections.InsertOne(nil, document)
}

//UpdateOne update one document
func UpdateOne(filter, set bson.M) (*mongo.UpdateResult, error) {
	return collections.UpdateOne(nil, filter, set)
}
func Aggregate(pipeline bson.A) (*mongo.Cursor, error) {
	return collections.Aggregate(nil, pipeline)
}

func Count() (int64, error) {
	return collections.CountDocuments(nil, bson.M{})
}
func Distict(field string, filter bson.M) ([]interface{}, error) {
	return collections.Distinct(nil, field, filter)
}

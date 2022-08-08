package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var result struct {
	ShortUrl    string
	OriginalUrl string
}

func ConnectDB() *mongo.Client {
	godotenv.Load()
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

var client *mongo.Client = ConnectDB()

func findOriginalLink(shortUrl string) (string, error) {
	collection := client.Database("url_shorter").Collection("urls")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.D{{Key: "shortUrl", Value: shortUrl}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("url doesn't exists")
	} else if err != nil {
		return "", err
	}
	return result.OriginalUrl, nil
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func createShortLink(url string) (string, error) {
	collection := client.Database("url_shorter").Collection("urls")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := collection.FindOne(ctx, bson.D{{Key: "url", Value: url}})
	if result.Err() == mongo.ErrNoDocuments {
		shortUrl := RandStringBytes(6)
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		_, err := collection.InsertOne(ctx, bson.D{{Key: "originalUrl", Value: url}, {Key: "shortUrl", Value: shortUrl}})
		if err != nil {
			return "", err
		}
		return shortUrl, nil
	} else if result.Err() != nil {
		return "", result.Err()
	}
	return "", nil
}

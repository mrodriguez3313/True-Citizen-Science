package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"encoding/json"
	"time"
	"context"
	"regexp"
	"reflect"
	// "io"
	// "io/ioutil"
	"os"
	// "os/exec"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IPFSContent struct {
	Filename string
}



func main() {
	// initializing client connection to mongodb service
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil { log.Fatal(err) }
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil { log.Fatal(err) }
	fmt.Println("connected to MongoDB!")

	// This is correct. Lines is an array of strings where the index has the content of each file.
	lines, err := scanLines("/home/marco/Documents/InsightDC/hash_outputs")
	if err != nil { log.Fatal(err) }
	tokenizer := regexp.MustCompile(`;`)

	// ipfscontent := make(map[string]interface{})

	var collection *mongo.Collection
	// For every line in file: tokenize the line, remove the tag, create object, insert to mongodb in the 'tag' collection
	for idx, _ := range lines {
		// each iteration tag is being reinitialized [tag, string w/o tag]
		split_input := tokenizer.Split(string(lines[idx]), 2)
		tag := split_input[0]
		// fmt.Println("tag: ", tag)

		// Creating database called PrivateIPFSDB. with table called Files
		collection := client.Database("PrivateIPFSDB").Collection(tag)
		// fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")
		jsonIpfsContent, err := json.Marshal(strings.TrimLeft(split_input[1], " "))
		if err != nil { log.Fatal(err) }

		// fmt.Println("jsonIpfsContent: ", string(jsonIpfsContent))
		ipfsFile := &IPFSContent{ string(jsonIpfsContent) }
		// fmt.Println("ifpsfile: ", ipfsFile)
		insertRes, err := collection.InsertOne(context.TODO(), ipfsFile)
		if err != nil { log.Fatal(err) }

		fmt.Println("inserted one file: ", insertRes.InsertedID)
	}

	findOptions := options.Find()
	collection = client.Database("PrivateIPFSDB").Collection("Beaver")
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		log.Fatal(err)
		defer cursor.Close(ctx)
	}
		// iterate over docs using Next()
	for cursor.Next(ctx) {
		// Declare a result BSON object
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println("cursor.Next() error:", err)
			log.Fatal(err)
		}
		fmt.Println("\nresult type:", reflect.TypeOf(result))
		fmt.Println("result:", result)

	}

	err = client.Disconnect(context.TODO())
	if err != nil { log.Fatal(err)}

	fmt.Println("Connection to MongoDB closed.")

	}

	func scanLines(path string) ([]string, error) {

  file, err := os.Open(path)
  if err != nil {
     return nil, err
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanLines)

  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, nil
}

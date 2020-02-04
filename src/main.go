package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"context"
	"regexp"
	"os"
	// "reflect"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IPFSContent struct {
	Projectname string
	Filecontent string
}



func main() {
	// initializing client connection to mongodb service
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil { log.Fatal(err) }


	// Connect to MongoDB
	err = client.Connect(context.Background())
	if err != nil { log.Fatal(err) }
	fmt.Println("connected to MongoDB!")


	db := client.Database("PrivateIPFSDB")
	collection := db.Collection("AllProjects")
	var insertRes *mongo.InsertOneResult


	// lines is an array of strings where the index has the content of each file.
	lines, err := scanLines("/home/marco/Documents/InsightDC/hash_outputs")
	if err != nil { log.Fatal(err) }
	tokenizer := regexp.MustCompile(`;`)

	// For every line in file: tokenize the line, remove the tag, create object, insert to mongodb
	for idx, _ := range lines {
		// 52: This line will split file input between projectname and filecontent
		split_input := tokenizer.Split(string(lines[idx]), 2)
		// 54: assigns projectname to tag
		tag := split_input[0]
		// 56: initializes object to insert into db.
		ipfsFile := &IPFSContent{
			Projectname: tag,
			Filecontent: strings.TrimLeft(split_input[1], " "),
		}

		//inserts object into db, under collection "AllProjects"
		insertRes, err = collection.InsertOne(context.TODO(), ipfsFile)
		if err != nil { log.Fatal(err) }

		fmt.Println("inserted one file: ", insertRes.InsertedID)
	}

	fmt.Println("now were here.")

	err = client.Disconnect(context.TODO())
	if err != nil { log.Fatal(err)}

	fmt.Println("Connection to MongoDB closed.")

}

// This function takes in a file path and returns a string where each index in the array is every line in the file
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

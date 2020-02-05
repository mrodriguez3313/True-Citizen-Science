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

	// We check this again later on, but this is here to short circuit that check. In case something was provided but it was not a file.
	if len(os.Args) == 1 {
		fmt.Println("Please provide 1 or more files to add to mongodb. This file may be a file with many elements or a single element.")
		os.Exit(0)
	}
	// initializing client connection to mongodb service
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil { log.Fatal(err) }


	// Connect to MongoDB
	err = client.Connect(context.Background())
	if err != nil { log.Fatal(err) }
	fmt.Println("connected to MongoDB!")

	// creating Database and collection for easily importing to it later on.
	db := client.Database("PrivateIPFSDB")
	collection := db.Collection("AllProjects")

	// Input checking if args provided are valid files. Else the program exits.
	Arg := os.Args[1:]
	if len(Arg) >= 1 {
		for idx, file := range Arg {
			if _, err := os.Stat(file); err == nil {
				// file exists, add it into the db
				import_file(file, db, collection, client)
			} else if os.IsNotExist(err) {
					// file doesn't exist
					if idx == len(Arg)-1 {
						fmt.Printf("File %s does not exist.\n", file)
						os.Exit(1)
					} else {
					fmt.Printf("Sorry file %s does not exist. Trying next file...\n", file)
					continue
				}
			} else {
				// mmmmm something bad happened see error for details
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println("Please provide 1 or more files to add to mongodb.")
		os.Exit(0)
	}
	os.Exit(0)
}


// This function will insert a file into the database.
// INPUT: a file or path to file to parse each line and insert each line into db.
// Closes connection after done inserting every line in file.
func import_file(file string, db *mongo.Database, collection *mongo.Collection, client *mongo.Client) {
	// 71: lines is will hold an array of strings where the index has the content of each file.
	lines, err := scanLines(file)
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
		insertRes, err := collection.InsertOne(context.TODO(), ipfsFile)
		if err != nil { log.Fatal(err) }

		fmt.Println("inserted one file: ", insertRes.InsertedID)
	}

	err = client.Disconnect( context.TODO() )
	if err != nil { log.Fatal(err) }

	fmt.Println("Connection to MongoDB closed.")
}


// This function takes in a file path and returns a list of strings where each index in the array is every line in the file
func scanLines(path string) ([]string, error) {
	// Does not do any formatting checking. File provided MUST be a new file on each line. Each line must be in format: '<Projectname>; <Filecontent>'
	// A solution is to have the data, (BEFORE ADDING TO IPFS), in json format.
	//  This could be achieved by providing the user/citizen a simple n-steps, step-by-step process to add the data into a json file; where n is determined by the researcher for how many samples/data points they want to collect.
	file, err := os.Open(path)
  if err != nil {
     return nil, err
  }

  defer file.Close()
	// Create scanner object to look through file
  scanner := bufio.NewScanner(file)
	// scanner will now only look line by line in the file
  scanner.Split(bufio.ScanLines)

  var lines []string
	// iterate through each line in file
  for scanner.Scan() {
		// pull out the text
		text := scanner.Text()
		// creating list of every line in file
    lines = append(lines, text)
		// fmt.Println("current line: ", text)
  }

  return lines, nil
}

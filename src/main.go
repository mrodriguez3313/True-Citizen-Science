package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	// "encoding/json"
	"context"
	"regexp"
	// "reflect"
	"os"
	// "os/exec"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
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
	findOptions := options.Find()
	findOptions.SetLimit(5)
	// ipfscontent := []interface{}

	// lines is an array of strings where the index has the content of each file.
	lines, err := scanLines("/home/marco/Documents/InsightDC/hash_outputs")
	if err != nil { log.Fatal(err) }
	tokenizer := regexp.MustCompile(`;`)

	// For every line in file: tokenize the line, remove the tag, create object, insert to mongodb
	for idx, _ := range lines {
		// 51: This line will split file input between projectname and filecontent
		split_input := tokenizer.Split(string(lines[idx]), 2)
		// 53: assigns projectname to tag
		tag := split_input[0]
		// 54: initializes object to insert into db.
		ipfsFile := &IPFSContent{
			Projectname: tag,
			Filecontent: strings.TrimLeft(split_input[1], " "),
		}

		insertRes, err = collection.InsertOne(context.TODO(), ipfsFile)
		if err != nil { log.Fatal(err) }

		fmt.Println("inserted one file: ", insertRes.InsertedID)
	}

	fmt.Println("now were here.")
	// collection = client.Database("PrivateIPFSDB").Collection("Beaver")
	// cur, err := *collection.Find(context.Background(), bson.D{})
	// check := FindRecords(collection)
	// fmt.Println("a element: ", check[0])
	// if err != nil {
	// 	fmt.Println("Finding all documents ERROR:", err)
	// 	log.Fatal(err)
	// }
	// defer cur.Close(context.Background())
	// fmt.Println("passed error check of cursor.")
	// 	// iterate over docs using Next()
	// for cur.Next(context.TODO()) {
	// 	fmt.Println("inside for loop,")
	// 	// Declare a result BSON object
	// 	var result bson.M
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		fmt.Println("cursor.Next() error:", err)
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("\nresult type:", reflect.TypeOf(result))
	// 	fmt.Println("result:", result)
	//
	// }

	// err = client.Disconnect(context.TODO())
	// if err != nil { log.Fatal(err)}
	//
	// fmt.Println("Connection to MongoDB closed.")

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



//Find multiple documents
// func FindRecords(db mongo.Collection) ([]IPFSContent){
//     // err := godotenv.Load()
// 		//
//     // if err != nil {
//     //     fmt.Println(err)
//     // }
// 		//
//     // //Get database settings from env file
//     // //dbUser := os.Getenv("db_username")
//     // //dbPass := os.Getenv("db_pass")
//     // dbName := os.Getenv("db_name")
//     // docCollection := "retailMembers"
// 		//
//     // dbHost := os.Getenv("db_host")
//     // dbPort := os.Getenv("db_port")
//     // dbEngine := os.Getenv("db_type")
// 		//
//     // //set client options
//     // clientOptions := options.Client().ApplyURI("mongodb://" + dbHost + ":" + dbPort)
//     // //connect to MongoDB
//     // client, err := mongo.Connect(context.TODO(), clientOptions)
//     // if err != nil {
//     //     log.Fatal(err)
//     // }
// 		//
//     // //check the connection
//     // err = client.Ping(context.TODO(), nil)
//     // if err != nil {
//     //     log.Fatal(err)
//     // }
// 		//
//     // fmt.Println("Connected to " + dbEngine)
//     // db := client.Database(dbName).Collection(docCollection)
//
//     //find records
//     //pass these options to the Find method
//     //Set the limit of the number of record to find
//
//     //Define an array in which you can store the decoded documents
//     var results []IPFSContent
//
//     //Passing the bson.D{{}} as the filter matches  documents in the collection
//     cur, err := db.Find(context.TODO(), bson.D{{}}, findOptions)
//     if err !=nil {
//         log.Fatal(err)
//     }
//
//     //Finding multiple documents returns a cursor
//     //Iterate through the cursor allows us to decode documents one at a time
//     for cur.Next(context.TODO()) {
//         //Create a value into which the single document can be decoded
//         var elem IPFSContent
//         err := cur.Decode(&elem)
//         if err != nil {
//             log.Fatal(err)
//         }
//
//         results=append(results, elem)
//
//     }
//
//     if err := cur.Err(); err != nil {
//         log.Fatal(err)
//     }
//
//     //Close the cursor once finished
//     cur.Close(context.TODO())
//
//     fmt.Printf("Found multiple documents: %+v\n", results)
//
// 		return results
// }

package main

// go run get_url.go  http://localhost:3000/posts/2 Correia

import (
	"os"
	"net/http"
	"log"
	"encoding/json"
	"bytes"
)

func main() {
	
	authorData      := os.Args[1]
	titleData       := os.Args[2]
	descriptionData := os.Args[3]
	target          := os.Args[4]

	type book struct {
		author string
		title string
		description string
	} // book

	// Another data structure:
	//book := map[string]string{"author": authorData, "title": titleData, "description": descriptionData}

	newBook := book{
		author: authorData,
		title:  titleData,
		description: descriptionData,
	} // newBook

	json_bytes, _ := json.Marshal(newBook)
	
	response, err := http.Post(target, "application/json", bytes.NewBuffer(json_bytes))

	if err != nil {
		log.Fatal(err)
	} // if

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
	
} // main()

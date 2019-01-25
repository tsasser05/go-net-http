package main

// go run get_url.go  http://localhost:3000/posts/2 Correia

import (
	"fmt"
	"os"
	"net/http"
	"bytes"
	"encoding/json"
	"log"
	"io/ioutil"
)

func main() {
	
	authorData      := os.Args[1]
	titleData       := os.Args[2]
	reviewData := os.Args[3]
	target          := os.Args[4]

	type book struct {
		Author string
		Title string
		Review string
	} // book

	newBook := &book{
		Author: authorData,
		Title:  titleData,
		Review: reviewData,
	} // newBook

	jsonBook, _ := json.Marshal(newBook)
	
	resp, err := http.Post(target, "application/json", bytes.NewBuffer(jsonBook))

	if err != nil {
		log.Fatal(err)
	} // if

	defer resp.Body.Close()

	fmt.Println("Check HTTP response code:  ", resp.Status, "\n")

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}
	
	var y map[string]interface{}
	json.Unmarshal(body, &y)
	fmt.Printf("%+v\n\n", y)	

	for key, val := range y {
		if (key == "id") {
			continue
		} // if

		fmt.Println(key, "\t", val)

		if (val == "Correia") {
			fmt.Println("Found Correia in the response ", key, ": ", val, " \n")
		} else {
			fmt.Println("Did not find Correia in the response ", key, ": ", val, "\n")
		} // if

	} // for
	
} // main()

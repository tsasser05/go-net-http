package main

// go run get_all_items.go <type>
//
// Types:
// comments
// posts
// profile

//{
//  "posts": [
//    {
//      "id": 1,
//      "title": "json-server",
//      "author": "typicode"
//    }
//  ],
//  "comments": [
//    {
//      "id": 1,
//      "body": "some comment",
//      "postId": 1
//    }
//  ],
//  "profile": {
//    "name": "typicode"
//  }
//}


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
}


func main() {
    url := "http://localhost:3000/posts"
    fmt.Println("URL is: ", url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	} // if

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal(err)
	} // if

	fmt.Println("Body content containing all posts:")
    fmt.Println("-------------------------------------------------------")
	fmt.Printf("%s\n", body)
    fmt.Println("-------------------------------------------------------")

	var responseData Response
    json.Unmarshal(body, &responseData)

	fmt.Println("Show all Titles:")

	//fmt.Println(responseData)

    //for i := 0; i < len(responseData.Post); i++ {
	//	fmt.Println(responseData.Post[i])
	//} 

} // main()

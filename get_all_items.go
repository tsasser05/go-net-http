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
	"os"
	"net/http"
	"io/ioutil"
	"log"
)

type Post struct {
  ID        int    `json:"id"`
  Title     string `json:"title"`
  Author    string `json:"author"`
}

type Comment struct {
  ID        int    `json:"id"`
  Body      string `json:"body"`
  PostId    int    `json:"postId"`
}

type Profile struct {
  Name    string    `json:"name"`
}




func main() {
    item_type := os.Args[1]
    url_base := "http://localhost:3000/"
    url := url_base + item_type
    fmt.Println("URL is: ", url)
    fmt.Println("-------------------------------------------------------")
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	} // if

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal(err)
	} // if

	fmt.Printf("%s\n", body)
    fmt.Println("-------------------------------------------------------")

} // main()

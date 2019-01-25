package main

// go run get_item.go comments 1
//
// Creates:
//
// http://localhost:3000/comments/1

// {
//   "id": 1,
//   "body": "some comment",
//   "postId": 1
// }

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
    item_type := os.Args[1]
	comment_num := os.Args[2]
    url_base := "http://localhost:3000/"
    separator := "/"
    url := url_base + item_type + separator + comment_num
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

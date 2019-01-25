package main

// go run get_comment.go 1

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

//type Comment struct {
//	ID      int       `json:"id"`
//    Body    string    `json:"body"`
//    PostId  int       `json:"postId"`
//} 

func main() {
	comment_num := os.Args[1]
    url_base := "http://localhost:3000/comments/"
    url := url_base + comment_num
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

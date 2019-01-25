package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	target := os.Args[1]
	fmt.Println("URL is:  ", target)

    fmt.Println(" ")
	
	response, err := http.Get(target)

	if err != nil {
		log.Fatal(err)
	} // if

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	} // if

    fmt.Println("-------------------------------------------------------")
	fmt.Printf("%s\n", body)
    fmt.Println("-------------------------------------------------------")

} // main()

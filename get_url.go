package main

// go run get_url.go  http://localhost:3000/posts/2 Correia

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	target := os.Args[1]
	pattern := os.Args[2]
	fmt.Println("URL is:  ", target)
	fmt.Println("Pattern is:  ", pattern)
	
	response, err := http.Get(target)

	if err != nil {
		log.Fatal(err)
	} // if

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	} // if

	fmt.Printf("%s", body)

	//matched, err := regexp.MatchString(pattern, body)
	matched, err := regexp.Match(pattern, body)

	if err != nil {
                fmt.Println("\nFailed to find the string", pattern, "in body")
        } else {
                fmt.Println("\nPattern:  ", pattern, "was found:  ", matched)
        } // if 
	
} // main()

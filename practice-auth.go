package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
)


/*******************************************************
*
* Data Structures
*
*******************************************************/

type NewUser struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first-name"`
	LastName  string    `json:"last-name"`
	Username  string    `json:"user-name"`
} // #NewUser


/*******************************************************
* 
* main()
* 
* Requires json-server-auth running inside json-server
* 
* https://github.com/jeremyben/json-server-auth
* 
* json-server user.json -m /usr/local/lib/node_modules/json-server-auth
* 
* user.json
* -------------------------------------------------------
* {
*   "users": []
* }
* -------------------------------------------------------
* 
*******************************************************/

func main() {
	url := "http://localhost:3000/register"

	userData := NewUser{
		Email: "tsasser05@gmail.com",
		Password: "rath5858",
		FirstName: "Tom",
		LastName: "Sasser",
		Username: "tsasser05",
	} 

	jsonData, err := json.Marshal(userData)

	if err != nil {
		fmt.Println("Error in JSON Marshal call")
	}

	fmt.Println("Verify JSON data was marshaled")
	fmt.Println(string(jsonData))
	fmt.Println("\n\n")

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Printf("HTTP request error: %s", err)
	} 

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Client error: %s", err)
	} 

	defer response.Body.Close()

} // main()

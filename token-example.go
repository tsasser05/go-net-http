/*******************************************************
*
* Example using an existing token after hitting the login API
*
*******************************************************/

func getAllContacts(token string, url string) {
	var client http.Client
	testPassed := false

    request, newRequestError := http.NewRequest("GET", url, nil)

    if newRequestError != nil {
        fmt.Printf("newRequestError:  %s", newRequestError)
    }

	var authStuff string = "Bearer " + token
	request.Header.Add("Authorization", authStuff)

	response, clientError := client.Do(request)
	
	if clientError != nil {
        fmt.Printf("clientError:  %s", clientError)
	} // if

	defer response.Body.Close()


-------------------------------------------------------

/*******************************************************
*
* login(loginURL string) string 
*
*******************************************************/

func login(loginURL string) string {

	var client http.Client
	var token string

	request, newRequestError := http.NewRequest("GET", loginURL, nil)

	if newRequestError != nil {
        fmt.Printf("newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil {
        fmt.Printf("clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

	type TokenResp struct {
		Token string `json:"token"`
	}


	if response.StatusCode == 200 {
		fmt.Println("HTTP 200 received\n")
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("\n-------------------------------------------------------\n")
		fmt.Println("\n\nResponse body:", string(body))
		fmt.Println("\n-------------------------------------------------------\n")

		var loginToken TokenResp
		json.Unmarshal([]byte(body), &loginToken)
		token = loginToken.Token
		fmt.Printf("login(): Token is: %s\n\n", token)

	} else {
		fmt.Println("Other code received\n")
	}	

	return token

} // login()

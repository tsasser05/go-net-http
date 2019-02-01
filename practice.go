package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
    "encoding/base64"
    "encoding/json"
	"bytes"
	"strconv"
)

 
/*******************************************************
*
* basicAuth
*
*******************************************************/

func basicAuth(request *http.Request) *http.Request  {
	username := "simuser"
    password := "simpassword"
    auth := username + ":" + password
	authStr :=  base64.StdEncoding.EncodeToString([]byte(auth))
    request.Header.Add("Authorization","Basic "+authStr)

	return request

} // basicAuth


/*******************************************************
*
* getAllContacts
*
*******************************************************/

func getAllContacts(dName string) {
	var client http.Client
	url := "http://localhost:3000/contacts"
	testPassed := false

    request, newRequestError := http.NewRequest("GET", url, nil)

    if newRequestError != nil {
        fmt.Printf("newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil {
        fmt.Printf("clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

	if response.StatusCode == 200 { // OK

		type ContactEmail struct {
			Address string `json:"address"`
		} 

		type ContactName struct {
			FirstName string `json:"first-name"`
			LastName  string `json:"last-name"`
            Prefix    string `json:"prefix"`
            FullName  string `json:"full-name"`
		}

		type ContactPhone struct {
			Formatted string `json:"formatted"`
			Label     string `json:"label"`
            Primary   bool   `json:"primary"`
		} 

		type Contact struct {
            ID          int            `json:"id"`
			DisplayName string          `json:"display-name"`
			Name        ContactName     `json:"name"`
			Phone       []ContactPhone  `json:"phone"`
			Email       []ContactEmail  `json:"email"`
		} 

		bodyBytes, ioutilError := ioutil.ReadAll(response.Body)

		if ioutilError != nil {
			ioErrStr := "ioutil ReadAll call on the response body failed"
			fmt.Println(ioErrStr)
		} 

		var contact []Contact
		json.Unmarshal([]byte(bodyBytes), &contact)

		for _, q := range contact {
			if q.DisplayName == dName {
				fmt.Printf("Contact: %+v\n", q)
				testPassed = true
			}
		} 

	} else {
		httpCode := string(response.StatusCode)
		errStr := "HTTP Response code was " + httpCode
		fmt.Println(errStr)
	} 

	if testPassed {
		fmt.Println("getAllContacts for a user passed\n")
	} else {
		fmt.Println("getAllContacts for a user failed\n")
	}

} // getAllContacts()


/*******************************************************
*
* getAllMessages(owner string)
*
* Get all messages for a user
*
* TBD .. Further abstraction necessary since code is redundant.
*
*******************************************************/

func getAllMessages(owner string) {
	var client http.Client
	url := "http://localhost:3000/messages"
	testPassed := false

    request, newRequestError := http.NewRequest("GET", url, nil)

    if newRequestError != nil {
        fmt.Printf("newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil {
		fmt.Printf("clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

	if response.StatusCode == 200 { // OK

		type Message struct {
			Owner string `json:"owner"`
			Text string `json:"text"`
			SenderID string `json:"sender-id"`
			MessageID string `json:"message-id"`
		} 

		bodyBytes, ioutilError := ioutil.ReadAll(response.Body)

		if ioutilError != nil {
			ioErrStr := "ioutil ReadAll call on the response body failed"
			fmt.Println(ioErrStr)
		} 

		var msg []Message
		json.Unmarshal([]byte(bodyBytes), &msg)

		for _, q := range msg {
			if q.Owner == owner {
				fmt.Printf("Message: %+v\n", q)
				testPassed = true
			}
		} 

	} else {
		httpCode := string(response.StatusCode)
		errStr := "HTTP Response code was " + httpCode
		fmt.Println(errStr)
	} 

	if testPassed {
		fmt.Println("getAllMessages for a user passed\n")
	} else {
		fmt.Println("getAllMessages for a user failed\n")
	}

} // getAllMessages()


/*******************************************************
*
* getMessageByID(messageID string)
*
* TBD .. Further abstraction necessary since code is redundant.
*
*******************************************************/

func getMessageByID(msgID string) {
	var client http.Client
	url := "http://localhost:3000/messages"
	testPassed := false

    request, newRequestError := http.NewRequest("GET", url, nil)

    if newRequestError != nil {
		fmt.Printf("newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil 
		fmt.Printf("clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

	if response.StatusCode == 200 { // OK

		type Message struct {
			Owner string `json:"owner"`
			Text string `json:"text"`
			SenderID string `json:"sender-id"`
			MessageID string `json:"message-id"`
		} 

		bodyBytes, ioutilError := ioutil.ReadAll(response.Body)

		if ioutilError != nil {
			ioErrStr := "ioutil ReadAll call on the response body failed"
			fmt.Println(ioErrStr)
		} 

		var msg []Message
		json.Unmarshal([]byte(bodyBytes), &msg)

		for _, q := range msg {
			if q.MessageID == msgID {
				fmt.Printf("Message: %+v\n", q)
				testPassed = true
			}
		} 

	} else {
		httpCode := string(response.StatusCode)
		errStr := "HTTP Response code was " + httpCode
		fmt.Println(errStr)
	} 

	if testPassed {
		fmt.Println("getMessageByID passed\n")
	} else {
		fmt.Println("getMessageByID failed\n")
	}

} // getMessageByID()



/*******************************************************
*
* getContactByID(contactID int)
*
* TBD .. Further abstraction necessary since code is redundant.
*
*******************************************************/

func getContactByID(contactID int) {
	var client http.Client
	url := "http://localhost:3000/contacts"
	testPassed := false

    request, newRequestError := http.NewRequest("GET", url, nil)

    if newRequestError != nil {
        fmt.Printf("newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil {
		fmt.Printf("clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

	if response.StatusCode == 200 { // OK

        type ContactEmail struct {
            Address string `json:"address"`
        }

        type ContactName struct {
            FirstName string `json:"first-name"`
            LastName  string `json:"last-name"`
            Prefix    string `json:"prefix"`
            FullName  string `json:"full-name"`
        }

        type ContactPhone struct {
            Formatted string `json:"formatted"`
            Label     string `json:"label"`
            Primary   bool   `json:"primary"`
        }

        type Contact struct {
            ID          int             `json:"id"`
            DisplayName string          `json:"display-name"`
            Name        ContactName     `json:"name"`
            Phone       []ContactPhone  `json:"phone"`
            Email       []ContactEmail  `json:"email"`
        }


		bodyBytes, ioutilError := ioutil.ReadAll(response.Body)

		if ioutilError != nil {
			ioErrStr := "ioutil ReadAll call on the response body failed"
			fmt.Println(ioErrStr)
		} 

		var con []Contact
		json.Unmarshal([]byte(bodyBytes), &con)

		for _, q := range con {
			if q.ID == contactID {
				fmt.Printf("Message: %+v\n", q)
				testPassed = true
			}
		} 

	} else {
		httpCode := string(response.StatusCode)
		errStr := "HTTP Response code was " + httpCode
		fmt.Println(errStr)
	} 

	if testPassed {
		fmt.Println("getContactByID passed\n")
	} else {
		fmt.Println("getContactByID failed\n")
	}

} // getContactByID()


/*******************************************************
*
* postContact()
*
* TBD .. abstract out the http stuff
*
*******************************************************/

func postContact() {
    url := "http://localhost:3000/contacts"

	// TBD should create this outside of main() and pass new instance into this function
	type ContactEmail struct {
		Address string `json:"address"`
	}
	
	type ContactName struct {
		FirstName string `json:"first-name"`
		LastName  string `json:"last-name"`
		Prefix    string `json:"prefix"`
		FullName  string `json:"full-name"`
	}
	
	type ContactPhone struct {
		Formatted string `json:"formatted"`
		Label     string `json:"label"`
		Primary   bool   `json:"primary"`
	}
	
	type Contact struct {
		ID          int             `json:"id"`
		DisplayName string          `json:"display-name"`
		Name        ContactName     `json:"name"`
		Phone       []ContactPhone  `json:"phone"`
		Email       []ContactEmail  `json:"email"`
	}

	// Should pass in from main() as arg to this function
	newContact := Contact{
		ID: 50,
		DisplayName: "foobarbaz",
		Name: ContactName {
			FirstName: "bar",
			LastName: "baz",
			Prefix: "foo",
			FullName: "foo bar baz",
		},
		Phone: 	[]ContactPhone{
			{
				Formatted: "777-888-9999",
				Label: "Foo Main Line",
				Primary: true,
			},
		},
		Email: []ContactEmail{
			{
				Address: "foo@gmail.com",
			},
		},
	} // struct newContact


    contactPostData, _ := json.Marshal(newContact)

	// Success ... use the following to see it
	fmt.Println("Verify json data was marshaled")
	fmt.Println(string(contactPostData))
	fmt.Println("\n\n")


	request, httpNewRequestErr := http.NewRequest("POST", url, bytes.NewBuffer(contactPostData))
	request.Header.Set("Content-Type", "application/json")

	if httpNewRequestErr != nil {
		fmt.Println("postContact():: new http request creation failed")
	}

    client := http.Client{}
	response, clientError := client.Do(basicAuth(request))

	if clientError != nil {
		fmt.Println("postContact() http Post call failed\n")
	} 

	defer response.Body.Close()

	fmt.Println("\n\nHTTP response status", response.Status)

	fmt.Println("\n\nResponse headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("\n\nResponse body:", string(body))

	var httpCode int = response.StatusCode

	fmt.Printf("\n\nThe httpCode was: %d\n\n", httpCode)

	if httpCode == 404 {
		fmt.Println("\nHTTP POST failed with HTTP 404 result.\n\n")
	} else if httpCode == 200|201 {
		fmt.Printf("\nHTTP POST successful with HTTP %d result.\n\n", httpCode)
	} else if httpCode == 500 {
		fmt.Println("\nHTTP POST failed with HTTP 500 due to duplicate ID.\n\n")
	}

} // postContact()



/*******************************************************
*
* postMessage(owner string)
*
* TBD .. abstract out the http stuff
*
*******************************************************/

func postMessage(owner string) {
    url := "http://localhost:3000/messages"

	// TBD should create this outside of main() and pass new instance into this function
	type Message struct {
		Owner string `json:"owner"`
		Text string `json:"text"`
		SenderID string `json:"sender-id"`
		MessageID string `json:"message-id"`
	}

	newMsg := Message{
		Owner: owner,
		Text: "Check out my new Youtube Music playlist!",
		SenderID: "db62e230-86cb-11e2-9e96-0800200c9a66-3",
		MessageID: "27a2ecfb-89c7-4854-9603-90d30f92dc4-d",
	}

	fmt.Printf("Owner is: %s\n", newMsg.Owner)

    msgPostData, err := json.Marshal(newMsg)

	if err != nil {
		fmt.Printf("JSON Marshal error %v", err)
	} 

	// Success ... use the following to see it
	fmt.Println("Verify json data was marshaled")
	fmt.Println(string(msgPostData))
	fmt.Println("\n\n")

	request, httpNewRequestErr := http.NewRequest("POST", url, bytes.NewBuffer(msgPostData))
	request.Header.Set("Content-Type", "application/json")

	if httpNewRequestErr != nil {
		fmt.Println("postMessage():: new http request message creation failed")
	}

    client := http.Client{}
	response, clientError := client.Do(basicAuth(request))

	if clientError != nil {
		fmt.Println("postMessage() http Post call failed\n")
	} 

	defer response.Body.Close()

	fmt.Println("\n\nHTTP response status", response.Status)

	fmt.Println("\n\nResponse headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("\n\nResponse body:", string(body))

	var httpCode int = response.StatusCode

	fmt.Printf("\n\nThe httpCode was: %d\n\n", httpCode)

	if httpCode == 404 {
		fmt.Println("\nHTTP POST failed with HTTP 404 result.\n\n")
	} else if httpCode == 200|201 {
		fmt.Printf("\nHTTP POST successful with HTTP %d result.\n\n", httpCode)
	} else if httpCode == 500 {
		fmt.Println("\nHTTP POST failed with HTTP 500.\n\n")
	}

} // postMessage()


/*******************************************************
*
* deleteContact(contactID int)
*
* TBD .. abstract out the http stuff
*
*******************************************************/

func deleteContact(contactID int) {
	var client http.Client
	id := strconv.Itoa(contactID)
	url := "http://localhost:3000/contacts/" + id

    request, newRequestError := http.NewRequest("DELETE", url, nil)

    if newRequestError != nil {
        fmt.Printf("deleteContact()::newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil {
        fmt.Printf("deleteContact()::clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

    fmt.Println("\n\nHTTP response status", response.Status)

    fmt.Println("\n\nResponse headers:", response.Header)
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println("\n\nResponse body:", string(body))

    var httpCode int = response.StatusCode

    fmt.Printf("\n\nThe httpCode was: %d\n\n", httpCode)

    if httpCode == 404 {
        fmt.Println("\nHTTP POST failed with HTTP 404 result.\n\n")
    } else if httpCode == 200|201 {
        fmt.Printf("\nHTTP POST successful with HTTP %d result.\n\n", httpCode)
    } else if httpCode == 500 {
        fmt.Println("\nHTTP POST failed with HTTP 500 due to duplicate ID.\n\n")
    }

} // deleteContact()


/*******************************************************
*
* deleteMessage(messageID string)
*
* TBD .. abstract out the http stuff
*
*******************************************************/

func deleteMessage(messageID string) {
	var client http.Client
	url := "http://localhost:3000/messages?message-id=" + messageID

	// Incorrect URL:
	//url := "http://localhost:3000/messages/" + messageID

	// Data is here but the call fails with HTTP 404:
    //curl --user simuser:simpassword --request GET http://localhost:3000/messages?message-id=27a2ecfb-89c7-4854-9603-90d30f92dc4-c

	fmt.Println("Constructed URL is:  ", url, "\n")

    request, newRequestError := http.NewRequest("DELETE", url, nil)

    if newRequestError != nil {
        fmt.Printf("deleteMessage()::newRequestError:  %s", newRequestError)
    }

	response, clientError := client.Do(basicAuth(request))
	
	if clientError != nil {
        fmt.Printf("deleteMessage()::clientError:  %s", clientError)
	} // if

	defer response.Body.Close()

    fmt.Println("\n\nHTTP response status", response.Status)

    fmt.Println("\n\nResponse headers:", response.Header)
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println("\n\nResponse body:", string(body))

    var httpCode int = response.StatusCode

    fmt.Printf("\n\nThe httpCode was: %d\n\n", httpCode)

    if httpCode == 404 {
        fmt.Println("\nHTTP POST failed with HTTP 404 result.\n\n")
    } else if httpCode == 200|201 {
        fmt.Printf("\nHTTP POST successful with HTTP %d result.\n\n", httpCode)
    } else if httpCode == 500 {
        fmt.Println("\nHTTP POST failed with HTTP 500 due to duplicate ID.\n\n")
    }

} // deleteMessage()


/*******************************************************
*
* updateContact()
*
* TBD .. abstract out the http stuff
*
*******************************************************/

func updateContact(contactID int) {
	id := strconv.Itoa(contactID)
    //url := "http://localhost:3000/contacts" + id 
	url := "http://localhost:3000/contacts?id=" + id

	fmt.Printf("url is %s\n", url)

	type NewContact struct {
		DisplayName string `json:"display-name"`
	} 

	newContact := NewContact{
		DisplayName: "tsasser055555",
	} 

    contactPostData, _ := json.Marshal(newContact)

	// Success ... use the following to see it
	fmt.Println("Verify json data was marshaled")
	fmt.Println(string(contactPostData))
	fmt.Println("\n\n")


	request, httpNewRequestErr := http.NewRequest("PUT", url, bytes.NewBuffer(contactPostData))
	request.Header.Set("Content-Type", "application/json")

	if httpNewRequestErr != nil {
		fmt.Println("updateContact():: new http request creation failed")
	}

    client := http.Client{}
	response, clientError := client.Do(basicAuth(request))

	if clientError != nil {
		fmt.Println("updateContact() http Post call failed\n")
	} 

	defer response.Body.Close()

	fmt.Println("\n\nHTTP response status", response.Status)

	fmt.Println("\n\nResponse headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("\n\nResponse body:", string(body))

	var httpCode int = response.StatusCode

	fmt.Printf("\n\nThe httpCode was: %d\n\n", httpCode)

	if httpCode == 404 {
		fmt.Println("\nHTTP PUT failed with HTTP 404 result.\n\n")
	} else if httpCode == 200|201 {
		fmt.Printf("\nHTTP PUT successful with HTTP %d result.\n\n", httpCode)
	} else if httpCode == 500 {
		fmt.Println("\nHTTP PUT failed with HTTP 500 due to duplicate ID.\n\n")
	}

} // updateContact()










/*******************************************************
*
* Main ..... now I can see it!
*
*******************************************************/

func main() {

	const debug int = 1
    const baseURL string = "http://localhost:3000/"
    const contacts_url string = baseURL + "contacts"
	const messages_url string = baseURL + "messages"

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 1:  Get all messages for user zeushamr")
	fmt.Println("----------------------------------------------------------------------\n")
	getAllMessages("zeushamr")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 2:  Get all messages for nonexistent user")
	fmt.Println("----------------------------------------------------------------------\n")
	getAllMessages("foo")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 3: Get a message by id")
	fmt.Println("----------------------------------------------------------------------\n")
    getMessageByID("27a2ecfb-89c7-4854-9603-90d30f92dc4-b")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 4: Get a message by id missing ID")
	fmt.Println("----------------------------------------------------------------------\n")
    getMessageByID("27a2ecfb-89c7-4854-9603-90d30f")

	fmt.Println("----------------------------------------------------------------------")
    fmt.Println("Test 5: Get all contacts for display-name=tsasser05")
	fmt.Println("----------------------------------------------------------------------\n")
	getAllContacts("tsasser05")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 6: Get all contacts for non-existent display-name=foo")
	fmt.Println("----------------------------------------------------------------------\n")
	getAllContacts("foo")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 7: Get contact by ID=1")
	fmt.Println("----------------------------------------------------------------------\n")
	getContactByID(1)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 8: Create new contact")
	fmt.Println("----------------------------------------------------------------------\n")
	postContact()

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 9 Get created Contact by ID 50 (foobarbaz)")
	fmt.Println("----------------------------------------------------------------------\n")
    getContactByID(50)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 10: POST Attempt to create existing contact again")
	fmt.Println("----------------------------------------------------------------------\n")
	postContact()

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 11: DELETE contact we created earlier ID 50.")
	fmt.Println("----------------------------------------------------------------------\n")
	deleteContact(50)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 12: Get non-existent contact by ID 10000")
	fmt.Println("----------------------------------------------------------------------\n")
	getContactByID(10000)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 13: Verify deleted contact is gone ID 50.")
	fmt.Println("----------------------------------------------------------------------\n")
	getContactByID(50)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 14: Try to delete a non-existent contact ID 5000.")
	fmt.Println("----------------------------------------------------------------------\n")
	deleteContact(5000)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 15: Post a message for owner tsasser05")
	fmt.Println("----------------------------------------------------------------------\n")
	//postMessage("tsasser05")
	fmt.Println("Code fails due to server error or possible struct problem.\n")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 16: Try to delete message with message-id 27a2ecfb-89c7-4854-9603-90d30f92dc4-b")
	fmt.Println("----------------------------------------------------------------------\n")
	//deleteMessage("27a2ecfb-89c7-4854-9603-90d30f92dc4-b")
	fmt.Println("Code fails with HTTP 404, but data is there in browser.  Likely node body response read problem in json-server per internet search.")

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Test 17: Update contact with new data.")
	fmt.Println("----------------------------------------------------------------------\n")
	//updateContact(1)
	fmt.Println("Code fails with HTTP 404, but data is there in browser.  Likely node body response read problem in json-server per internet search.")
	// tried and both fail
	// http://localhost:3000/contacts?id=1
	// http://localhost:3000/contacts/1



} // main()


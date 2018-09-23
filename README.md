# go-net-http

#### Example code using net/http

## Installation
Relies on local npm server. Read the docs here:  [https://jsonplaceholder.typicode.com/]

1) Install npm
2) Install json-server:
    ```
    sudo npm install -g json-server
    ```
3) Add some data to ~/npm/db.json
4) Create a ~/npm/db.json file with some data
```{
  "posts": [
    { "id": 1, "title": "json-server", "author": "typicode" }
  ],
  "comments": [
    { "id": 1, "body": "some comment", "postId": 1 }
  ],
  "profile": { "name": "typicode" }
}
```
5) Start json-server:  
```
json-server --watch db.json
```
6) Access in browser:  http://localhost:3000/posts/1
  
    
    

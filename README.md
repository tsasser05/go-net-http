# go-net-http

Example code using net/http

Relies on local npm server.

Install npm
Install json-server:
    sudo npm install -g json-server
Add some data to db.json in ~/npm

Create a db.json file with some data

{
  "posts": [
    { "id": 1, "title": "json-server", "author": "typicode" }
  ],
  "comments": [
    { "id": 1, "body": "some comment", "postId": 1 }
  ],
  "profile": { "name": "typicode" }
}

Start json-server:  json-server --watch db.json

http://localhost:3000/posts/1
  
    
    

# go-gin-gorm-postgres-in-cloud

Go : RESTful CRUD APIs operations using the Gin, GORM package for PostgreSQL as a Service

This repository contains developing of RESTful web service APIs in Go Programming Language using Gin Web Framework (Gin) and uses ORM library for Golang (GORM) package to connect with PostgreSQL as a Service Cloud Database for DB CRUD operations.

**Prerequisite:**
- Make sure GO already installed and had working environment.<br>
while developing this code go version used,<br>
```
$go version
go version go1.18.4 darwin/amd64 
```
- Create a PostgreSQL as a Service instance (Tiny Turtle - Free plan) from https://www.elephantsql.com/

**Steps to Run:**
- Run `go run main.go` in terminal
- Go to another command terminal and test the REST APIs using CURL (or use REST client like Postman).

**Test Results:**<br>
CRUD REST APIs test using CURL and it's Results:

```
% curl -X GET http://localhost:8080/ping
{"message":"pong"}
```
Also you can verify this using any browser by entering the `http://localhost:8080/ping` url.

Create an book entry
```
% curl -iX POST http://localhost:8080/book -H "Content-Type: application/json" -d @book1.json    
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 12 Mar 2023 02:52:41 GMT
Content-Length: 163

{"ID":7,"CreatedAt":"2023-03-11T18:52:41.913989-08:00","UpdatedAt":"2023-03-11T18:52:41.913989-08:00","DeletedAt":null,"title":"Go Programming Adv","price":"32 $"
```

List all book entries
```
% curl -X GET http://localhost:8080/book | jq                                                
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   816  100   816    0     0   5764      0 --:--:-- --:--:-- --:--:--  6135
{
  "books": [
    {
      "ID": 1,
      "CreatedAt": "2023-03-11T12:28:14.132098-08:00",
      "UpdatedAt": "2023-03-11T12:28:14.132098-08:00",
      "DeletedAt": null,
      "title": "C Programming",
      "price": "20 $"
    },
    {
      "ID": 2,
      "CreatedAt": "2023-03-11T12:29:28.86489-08:00",
      "UpdatedAt": "2023-03-11T13:11:45.371602-08:00",
      "DeletedAt": null,
      "title": "C Prog Edition 2",
      "price": "22 $"
    },
    {
      "ID": 4,
      "CreatedAt": "2023-03-11T12:45:40.632866-08:00",
      "UpdatedAt": "2023-03-11T12:45:40.632866-08:00",
      "DeletedAt": null,
      "title": "C++ Programming",
      "price": "25 $"
    },
    {
      "ID": 5,
      "CreatedAt": "2023-03-11T13:18:51.601652-08:00",
      "UpdatedAt": "2023-03-11T13:18:51.601652-08:00",
      "DeletedAt": null,
      "title": "Go Programming",
      "price": "30 $"
    },
    {
      "ID": 7,
      "CreatedAt": "2023-03-11T18:52:41.913989-08:00",
      "UpdatedAt": "2023-03-11T18:52:41.913989-08:00",
      "DeletedAt": null,
      "title": "Go Programming Adv",
      "price": "32 $"
    }
  ]
}
```

Get a specific book entry by ID
```
% curl -X GET http://localhost:8080/book/2 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   169  100   169    0     0   1077      0 --:--:-- --:--:-- --:--:--  1157
{
  "book": {
    "ID": 2,
    "CreatedAt": "2023-03-11T12:29:28.86489-08:00",
    "UpdatedAt": "2023-03-11T13:11:45.371602-08:00",
    "DeletedAt": null,
    "title": "C Prog Edition 2",
    "price": "22 $"
  }
}
```

Update an book entry by ID
```
 % curl -iX PATCH http://localhost:8080/book/2 -H "Content-Type: application/json" -d @book_update.json 
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 12 Mar 2023 02:55:55 GMT
Content-Length: 169

{"book":{"ID":2,"CreatedAt":"2023-03-11T12:29:28.86489-08:00","UpdatedAt":"2023-03-11T18:55:55.070844-08:00","DeletedAt":null,"title":"C Prog Edition 3","price":"23 $"}}

% curl -X GET http://localhost:8080/book/2 | jq                                                       
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   169  100   169    0     0   1295      0 --:--:-- --:--:-- --:--:--  1385
{
  "book": {
    "ID": 2,
    "CreatedAt": "2023-03-11T12:29:28.86489-08:00",
    "UpdatedAt": "2023-03-11T18:55:55.070844-08:00",
    "DeletedAt": null,
    "title": "C Prog Edition 3",
    "price": "23 $"
  }
}
```

Delete an book entry
```
% curl -X DELETE http://localhost:8080/book/5

% curl -X GET http://localhost:8080/book | jq                                                         
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   656  100   656    0     0   4976      0 --:--:-- --:--:-- --:--:--  5290
{
  "books": [
    {
      "ID": 1,
      "CreatedAt": "2023-03-11T12:28:14.132098-08:00",
      "UpdatedAt": "2023-03-11T12:28:14.132098-08:00",
      "DeletedAt": null,
      "title": "C Programming",
      "price": "20 $"
    },
    {
      "ID": 2,
      "CreatedAt": "2023-03-11T12:29:28.86489-08:00",
      "UpdatedAt": "2023-03-11T18:55:55.070844-08:00",
      "DeletedAt": null,
      "title": "C Prog Edition 3",
      "price": "23 $"
    },
    {
      "ID": 4,
      "CreatedAt": "2023-03-11T12:45:40.632866-08:00",
      "UpdatedAt": "2023-03-11T12:45:40.632866-08:00",
      "DeletedAt": null,
      "title": "C++ Programming",
      "price": "25 $"
    },
    {
      "ID": 7,
      "CreatedAt": "2023-03-11T18:52:41.913989-08:00",
      "UpdatedAt": "2023-03-11T18:52:41.913989-08:00",
      "DeletedAt": null,
      "title": "Go Programming Adv",
      "price": "32 $"
    }
  ]
}

```


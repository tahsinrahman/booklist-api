# Booklist-API #

## URLS ##

GET /books
----------
Shows list of all books

```
GET http://localhost:8080/books
HTTP/1.1 200 OK
Date: Thu, 22 Nov 2018 13:50:14 GMT
Content-Length: 410
Content-Type: text/plain; charset=utf-8
Request duration: 0.027841s

[
  {
    "id": 1,
    "name": "Linux in a Nutshell",
    "isbn": "0596009305",
    "authors": [
      {
        "first_name": "Ellen",
        "last_name": "Siever"
      },
      {
        "first_name": "Stephen",
        "last_name": "Figgins"
      }
    ]
  },
  {
    "id": 2,
    "name": "The Linux Command Line",
    "isbn": "1593273894",
    "authors": [
      {
        "first_name": "William",
        "last_name": "Scotts"
      }
    ]
  },
  {
    "id": 3,
    "name": "The Linux Programming Interface",
    "isbn": "1593272200",
    "authors": [
      {
        "first_name": "Michael",
        "last_name": "Kerrisk"
      }
    ]
  }
]
```

POST /books
-----------
Adds a new book
```
POST http://localhost:8080/books
HTTP/1.1 201 Created
Date: Thu, 22 Nov 2018 13:52:24 GMT
Content-Length: 168
Content-Type: text/plain; charset=utf-8
Request duration: 0.027163s

{
  "id": 4,
  "name": "Understanding the Linux Kernel",
  "isbn": "0596005652",
  "authors": [
    {
      "first_name": "Daniel",
      "last_name": "Bovet"
    },
    {
      "first_name": "Marco",
      "last_name": "cesati"
    }
  ]
}
```

GET /books/{id}
-----------------
Shows a single book
```
GET http://localhost:8080/books/10
HTTP/1.1 404 Not Found
Date: Thu, 22 Nov 2018 13:54:45 GMT
Content-Length: 25
Content-Type: text/plain; charset=utf-8
Request duration: 0.037078s

{
  "error": "book not found"
}
```

```
GET http://localhost:8080/books/1
HTTP/1.1 200 OK
Date: Thu, 22 Nov 2018 13:55:30 GMT
Content-Length: 160
Content-Type: text/plain; charset=utf-8
Request duration: 0.029872s

{
  "id": 1,
  "name": "Linux in a Nutshell",
  "isbn": "0596009305",
  "authors": [
    {
      "first_name": "Ellen",
      "last_name": "Siever"
    },
    {
      "first_name": "Stephen",
      "last_name": "Figgins"
    }
  ]
}
```


PUT /books/{id}
---------------
Updates a book

```
PUT http://localhost:8080/books/4
HTTP/1.1 200 OK
Date: Thu, 22 Nov 2018 13:56:40 GMT
Content-Length: 166
Content-Type: text/plain; charset=utf-8
Request duration: 0.133558s

{
  "id": 4,
  "name": "Understanding the Linux Kernel",
  "isbn": "new_isbn",
  "authors": [
    {
      "first_name": "Daniel",
      "last_name": "Bovet"
    },
    {
      "first_name": "Marco",
      "last_name": "cesati"
    }
  ]
}
```

DELETE /books/{id}
------------------
Deletes a book
```
DELETE http://localhost:8080/books/4
HTTP/1.1 200 OK
Date: Thu, 22 Nov 2018 13:57:09 GMT
Content-Length: 166
Content-Type: text/plain; charset=utf-8
Request duration: 0.029016s

{
  "id": 4,
  "name": "Understanding the Linux Kernel",
  "isbn": "new_isbn",
  "authors": [
    {
      "first_name": "Daniel",
      "last_name": "Bovet"
    },
    {
      "first_name": "Marco",
      "last_name": "cesati"
    }
  ]
}
```

```
GET http://localhost:8080/books/4
HTTP/1.1 404 Not Found
Date: Thu, 22 Nov 2018 13:57:27 GMT
Content-Length: 25
Content-Type: text/plain; charset=utf-8
Request duration: 0.011443s

{
  "error": "book not found"
}
```

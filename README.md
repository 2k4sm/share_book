# **_share_book_**

## A Web API which facilitates effortless book sharing.

## Tech Stack:-
`go(Programming Language.)`

`go-Fiber(go backend framework)`

`SQLite(Database)`

`GORM(Object Relation Mapper)`

### **_share_book_** uses the power of [go-Fiber](https://gofiber.io/) to build a clean API. For Storage **_share_book uses_** [SQLite](https://www.sqlite.org/index.html).

# Requirements:-
  `go > v1.16.4`
  
  `go-Fiber/v2 > v2.50.x`

  `SQLite-v3`

# API Design:-
![share_book](https://github.com/2k4sm/share_book/assets/101013814/95428636-859e-4679-aa2c-95eb0abd0af1)

# API Usage Guide:-
- To Add a book for sharing do:-
  - `PUT /api/v1/booky/`
- To Browse the shared books do:-
  - `GET /api/v1/booky/`
- To borrow a book for certain duration of time do:-
  - `PUT /api/v1/booky/<book_id>/borrow`
    - book_id -> ID of the Book that has been added for sharing with others
    - borrow_start_time -> Start Time from when the Book is being borrowed
    - borrow_end_time -> End Time till when the Book is being borrowed (a week).
- To view all the borrowed books:-
  - `GET /api/v1/booky/borrow`
- To return a borrowed book do:-
    - `POST /api/v1/booky/<book_id>/borrow/<borrow_id>`
      - book_id -> ID of the Book that has been added for sharing with others
      - borrow_id -> ID of the Borrow operation for Book with book_id

# Setting Up:-
- Start By cloning the repository:-
  ```
  git clone git@github.com:2k4sm/share_book.git
  ```
- Change directory to the cloned repository:-
    ```
    cd share_book
    ```
- Download dependencies using the go mod file:-
  ```
  go mod tidy
  ```
- To build the server do:-
    ```
    go build cmd/web/server.go
    ```
- To start the server do:-
    ```
    ./server
    ```
    From the root directory of _share_book_.

### Your server is started at port `:8000`.
-  Use curl/hoppscotch/postman to test the API.
    - To Add a book for sharing do:-
      - `PUT`
        ```
          localhost:8000/api/v1/booky/
        ```
    - To Browse the shared books do:-
      - `GET`
        ```
          localhost:8000/api/v1/booky/
        ```
    - To borrow a book for certain duration of time do:-
      - `PUT`
        ```
        localhost:8000/api/v1/booky/<book_id>/borrow
        ```
    - To view all the borrowed books:-
      - `GET`
        ```
        localhost:8000/api/v1/booky/borrow
        ```  
    - To return a borrowed book do:-
      - `POST`
        ```
        localhost:8000/api/v1/booky/<book_id>/borrow/<borrow_id>
        ```

           
TODO(Create these endpoints..):
- [x] Add a Book for sharing with others
- [x] Browse books that other users have put up for sharing 
- [x] Borrow a book from a user for a certain duration of time. Let's call these as borrow_start and borrow_end.
- [x] View all the borrowed Books.
- [x] Return a borrowed book


## **_Thanks for using share_book._**

# License

    MIT License

    Copyright (c) 2023 Shrinibas Mahanta
    
    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
    
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
    
    

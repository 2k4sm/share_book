# share_book

## A Web API which facilitates effortless book sharing.

### **_share_book_** uses the power of [go-Fiber](https://gofiber.io/) to build a clean API. For Storage share_book uses [SQLite](https://www.sqlite.org/index.html).

# Requirements:-
- go > v1.16.4
- go-Fiber/v2 > v2.50.x

# API Design:-
![share_book](https://github.com/2k4sm/share_book/assets/101013814/95428636-859e-4679-aa2c-95eb0abd0af1)

TODO:
- [x] Add a Book for sharing with others
- [x] Browse books that other users have put up for sharing
- [x] Borrow a book from a user for a certain duration of time. Let's call these as borrow_start_time and borrow_end_time
- [x] Return a borrowed book


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
## Your server is started at port `:8000`.
-  Use curl/hoppscotch/postman to test the API.

# API Usage Guide:-
- To Add a book for sharing do:-
  - PUT /api/v1/booky/
- To Browse the shared books do:-
  - GET /api/v1/booky/
- To borrow a book for certain duration of time do:-
  - PUT /api/v1/booky/<book_id>/borrow
    - book_id -> ID of the Book that has been added for sharing with others
    - borrow_start_time -> Start Time from when the Book is being borrowed
    - borrow_end_time -> End Time till when the Book is being borrowed
- To return a borrowed book do:-
    - POST /api/v1/booky/<book_id>/borrow/<borrow_id>
      - book_id -> ID of the Book that has been added for sharing with others
      - borrow_id -> ID of the Borrow operation for Book with book_id



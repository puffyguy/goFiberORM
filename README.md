# Go Book Service

```
The Go Book Service is a RESTful API for managing books. It is built using the Go programming language, the Fiber web framework, and supports both MongoDB and MySQL databases. This service allows you to perform various operations related to books, such as adding, updating, deleting, and retrieving book information.
```

# Endpoints

### /api/v1/books

- `GetBooks`: List all books details from database<br>
  **Method** - `GET`<br>

### /api/v1/books/:isbn

- `GetBook`: List deatils of single book from database using provided isbn<br>
  **Method** - `GET`<br>

### /api/v1/books

- `NewBook`: Creates a new book record in database<br>
  **Method** - `POST`<br>
  **BODY:** `JSON`<br>
  **Required fields:** <br> - `name` - string <br> - `isbn` - string <br> - `author` - string <br> - `price` - int <br>

### /api/v1/books/:isbn

- `UpdateBook`: Update single book details in the database using provided isbn<br>
  **Method** - `PUT`<br>
  **BODY:** `Params`<br>
  **Required fields:** <br> - `isbn` - string <br>
  **BODY:** `JSON`<br>
  **Required fields:** <br> - `name` - string <br> - `author` - string <br> - `price` - int <br>

### /api/v1/books/:isbn

- `DeleteBook`: Removes a single book from the database using provided isbn<br>
  **Method** - `DELETE`<br>
  **BODY:** `Params`<br>
  **Required fields:** <br> - `isbn` - string <br>

# Movies Server

Simple movies server that accepts CRUD requests using a different mutex from the default.

## Dependencies
```Bash
go get github.com/gorilla/mux
```

## API
| Operation | Address      | Description               |
|-----------|--------------|---------------------------|
| GET       | /movies      | Retrieves all movies      |
| POST      | /movies      | Creates a new movie       |
| GET       | /movies/{id} | Retrieve a specific movie |
| PUT       | /movies/{id} | Updates a movie           |
| DELETE    | /movies/{id} | Deletes a movie           |
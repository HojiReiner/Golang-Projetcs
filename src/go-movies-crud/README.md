# Movies Server

Simple movies server that accepts CRUD requests using a different mutex from the default.

Mutex: `go get github.com/gorilla/mux`


|Operation|Address|Description|
|---------|-------|-----------|
|GET| /movies|Retrieves all movies|
|GET| /movies/{id}|Retrieve a specific movie|
|POST| /movies|Creates a new movie|
|PUT| /movies/{id}|Updates a movie|
|DELETE| /movies/{id}|Deletes a movie|
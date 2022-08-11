# Book Store

Simple book store using a sql database and multiple files

## Dependencies

```bash
go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql     

go get -u github.com/gorilla/mux
```

## Useful Commands

```bash
#Initiate MySql server using docker container
docker-compose up -d
```

## API

| Operation | Address    | Description              |
|-----------|------------|--------------------------|
| GET       | /book      | Retrieves all books      |
| POST      | /book      | Creates a new book       |
| GET       | /book/{id} | Retrieve a specific book |
| PUT       | /book/{id} | Updates a book           |
| DELETE    | /book/{id} | Deletes a book           |
# Go CRM

Simple CRM created using go fiber framework

## Dependencies
```bash
go get -u github.com/gofiber/fiber/v2 
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## API
| Operation | Address           | Description              |
| --------- | ----------------- | ------------------------ |
| GET       | /api/v1/lead      | Retrieves all leads      |
| POST      | /api/v1/lead      | Creates a new lead       |
| GET       | /api/v1/lead/{id} | Retrieve a specific lead |
| DELETE    | /api/v1/lead/{id} | Deletes a lead           |

## Description
Sample CRUD API created in Go's Gin framework using the GORM ORM and PostgreSQL as the database

## Setup
Define environment variables in `.env` following the example provided in `.env.example`

## Routes
The following are the routes provided in this sample API:
### /products
```
GET / get list of all products
POST / create new product
GET /:id get the product registered under the specified id
PATCH /:id update the product registered under the specified id
DELETE /:id delete the product registered under the specified id
```
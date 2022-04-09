# Picus Security Golang Bootcamp Graduation Project
A RESTful API example for simple E-commerce application users can login and register.
Users can list categories,products and also they can add product their carts using JWT token
And also they can order the products which are added to their carts.

## Technologies 
* Golang
* GORM
* Gin-Gonic framework
* MySQL Database
* VIPER
* Testify
* Swagger Implementation
* Unit tests
* Pagination


## Installation & Run
```bash
# Download this project
go get github.com/FAdemoglu/GoBootcampGraduationProject
```

Before running API server, you should set the database config with yours or set the your database config with my values on [location.prod.yaml](https://github.com/FAdemoglu/GoBootcampGraduationProject/blob/main/config/location.prod.yaml)

```

```bash
# Build and Run
cd GoBootcampGraduationProject
go run main.go

# API Endpoint : http://localhost:8080

# Swagger Endpoint : http://localhost:8080/swagger/index.html
```

## Project Structure
```
├── config
│   ├── config.go              // Configuration Files contains yaml file
├── helper
│   ├── helper.go              //For reading csv files
├── internal
│   ├── api
│   │   ├── auth
│   │   │   ├── authcontroller.go
│   │   │   ├── types.go
│   │   ├── cart
│   │   │   ├── cartcontroller.go
│   │   │   ├── types.go
│   │   ├── category
│   │   │   ├── categorycontroller.go
│   │   │   ├── types.go
│   │   ├── order
│   │   │   ├── ordercontroller.go
│   │   ├── product
│   │   │   ├── productcontroller.go
│   │   │   ├── types.go
│   │   ├── router.go    //Configurations for database,migrations,groups for routing 
│   ├── config
│   │   ├── config.go
│   ├── domain
│   │   ├── cart
│   │   │   ├── entity.go
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   ├── category
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   ├── order
│   │   │   ├── entity.go
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   ├── products
│   │   │   ├── entity.go
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   ├── users
│   │   │   ├── entity.go
│   │   │   ├── errors.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
├── pkg
│   ├── database_handler
│   │   ├── mysql_handler.go
│   ├── gracefullyshut
│   │   ├── grafullyshutdown.go
│   ├── jwt
│   │   ├── jwt_helper.go
│   ├── middleware
│   │   ├── auth_middleware.go
│   │   ├── latency_logger.go
│   ├── pagination
│   │   ├── pagination.go
├── shared
│   ├── responses.go
└── main.go
```

## API

#### /user/login
* `POST` : Login with username and password

#### /user/register
* `POST` : Register with username, password and confirm password

#### /category/list
* `GET` : List categories 

#### /category/uploadcsv
* `POST` : Upload csv files which are contains categories  (This is for admin users)

#### /product/list
* `GET` : Get all products with pagination

#### /product/remove
* `DELETE` : Delete product from database (These just for admin users)

#### /product/create
* `POST` : Create product (For admin users)

#### /product/update
* `POST` : Update Product (For admin users)

#### /product/search?searched=...
* `GET` : Search product and also categories

#### /cart/list
* `GET` : Get products which are in users' cart

#### /cart/add
* `POST` : Add Product to cart

#### /cart/remove?Id=.
* `DELETE` : Delete product with Id

#### /cart/update?Id=.&ItemId=.&Count=?
* `PUT` : Update cart items

#### /order/list
* `GET` : Get all orders with pagination

#### /order/cancel
* `DELETE` : Cancel order if time no longer than 14 days

#### /order/create
* `POST` : Create Order
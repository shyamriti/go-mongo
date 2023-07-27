# Golang CRUD API Project - User Management with MongoDB and Gin

This is a simple CRUD API project implemented in Golang using the Gin framework and MongoDB for user management. The API allows users to perform basic operations such as register, login, log out, get user data, update user information, and delete their accounts.

## Requirements

- Go (Golang) version 1.16 or higher
- MongoDB

## Getting Started

1. Clone the repository to your local machine:

git clone https://github.com/your-username/golang-crud-api.git
cd golang-crud-api

## To install the dependancy

go mod tidy

## Run the application

go run main.go

## The following endpoints are available in this CRUD API

1. For register a new user

POST localhost:8080/user/register

2. Login

POST localhost:8080/user/login

3. Logout

POST localhost:8080/user/logout

4. Get a user

GET localhost:8080/user/find/:name

5. Update a user

PUT localhost:8080/user/update/:name

6. Delete a user

DELETE localhost:8080/user/delete/:name

## To run the test

go test ./...
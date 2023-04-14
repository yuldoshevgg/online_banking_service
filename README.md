# Online Banking Service

## Description


## Instructions
 - Create database with the name of "online_banking_service"
 - Migrate migrations. [migrate -path ./migrations -database 'postgres://{USER}:{PWD}@{HOST}:{PORT}/{DB}?sslmode=disable' up]
 - Run the application
   - go run cmd/main.go
 - Open http://localhost:8080/swagger/index.html#/

## How to use?
 - SignIn with username and password. Or Login if already have account. Both APIs return JWT Token and user_id which are needed for other actions. 
 - Copy JWT Token and authorize in Swagger to open APIs. 
 - Create bank account with user_id which is returned from SignIn/Login and account_number.
 - Pay for or withdraw from your account balance.
 - Transfer balance between accounts.
 - Get all accounts by username.
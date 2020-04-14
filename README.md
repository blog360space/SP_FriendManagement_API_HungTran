# SP_FriendManagement_API_HungTran
A simple test api with golang run in port ```8000```

## Required
- Mysql database mysql connection
- Golang installed

## Technical Stack
This Friends Management API service is a RESTful web API built on Golang, HTTP, GORM, Mysql and Unit Test 
by testify with 90% statement coverage.

### Project Structure
Project have 5 packages:
- main.go: Start server and routing.
- configs: Define constant.
- controllers: Handle request from client.
- models: Define structures models related to database table schema.
- repositories: Process logic business and actions query, update database.
- utils: Common functions: Db connect, Response json.
### UnitTest
There 2 databases wit the same schema 1 for main program and 1 for Unit Test. 

The testing mainly in the repositories package.

## Install
Create 2 databases ```friend_manager``` for the main program, ```friend_manager_test``` for unittest.
Use dump ```script db/friend_managerment.sql``` into database ```friend_manager``` and ```friend_manager_test```.
At the root of this project and run the following command:
```
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
```
To change db connection, please update file ```utils/Db.go```
To Start project:
```
cd rootOfThisProject
# Run run full test
go run main.go
# Run test with coverage
go test -cover
```

To run unitest:
```
cd rootOfThisProject/src/repositories
go test
```
## API Details

### 1. As a user, I need an API to create a friend connection between two email addresses.

Uri:
```
/api/user/create_friend
```
Method: POST

The API should receive the following JSON request:
```json
{
    "friends":
    [
        "andy@example.com",
        "john@example.com"
    ]
}
```
The Api response:
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "message": "Target lisa@example.com Blocked Requestor john@example.com",
    "success": false
}
```
### 2. As a user, I need an API to retrieve the friends list for an email address.
Uri:
```
/api/user/get_friend
```
The API should receive the following JSON request:
```json
{
    "email": "andy@example.com"
}
```
The Api response:
On success: HttpStatus = 200
```json
{
    "count": 2,
    "friends": [
        "john@example.com",
        "lisa@example.com"
    ],
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "message": "User blablabla@example.com not exits",
    "success": false
}
```
### 3. As a user, I need an API to retrieve the common friends list between two email addresses.
Uri:
```
/api/user/get_friend_common
```
The API should receive the following JSON request:
```json
{
    "count": 2,
    "friends": [
        "lisa@example.com",
        "andy@example.com"
    ],
    "success": true
}
```
The Api response:
On success: HttpStatus = 200
```json
{
    "success": true,
    "friends" :
    [
        "common@example.com"
    ],
    "count" : 1
}
```
On error: HttpStatus = 400
```json
{
    "message": "User blablabla@example.com not exits",
    "success": false
}
```
### 4. As a user, I need an API to subscribe to updates from an email address.
Uri:
```
/api/user/subscribe
```
The API should receive the following JSON request:
```json
{
    "requestor": "lisa@example.com",
    "target": "john@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "message": "User blablabla@example.com not exits",
    "success": false
}
```
### 5. As a user, I need an API to block updates from an email address.
Uri:
```
/api/user/block
```

Suppose "andy@example.com" blocks "john@example.com":
The API should receive the following JSON request:
```json
{
    "requestor": "andy@example.com",
    "target": "john@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true
}
```
On error: HttpStatus = 400
```json
{
    "message": "User blablabla@example.com not exits",
    "success": false
}
```
### 6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
Uri:
```
/api/post/create
```

The API should receive the following JSON request:
```json
{
    "sender": "john@example.com",
    "text": "Hello World! kate@example.com"
}
```
On success: HttpStatus = 200
```json
{
    "success": true,
    "recipients":
    [
        "lisa@example.com",
        "kate@example.com"
    ]
}
```
On error: HttpStatus = 400
```json
{
    "message": "User blablabla@example.com not exits",
    "success": false
}
```

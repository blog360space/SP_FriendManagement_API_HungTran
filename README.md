# SP_FriendManagement_API_HungTran
A simple test api with golang.
## Required
Mysql database mysql connection
## Install
Create 2 databases ```friend_manager``` for the main program, ```friend_manager_test``` for unittest.
Use dump ```script db/friend_managerment.sql``` into database ```friend_manager``` and ```friend_manager_test```.
At the root of this project and run the following command:
```
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
```
## Feature

### 1. As a user, I need an API to create a friend connection between two email addresses.

Uri:
```
/api/user/create_friend
```
Method: POST

The API should receive the following JSON request:
```
{
    "friends":
    [
        "andy@example.com"
        "john@example.com"
    ]
}
```
The API should return the following JSON response on success:
```
{
    "success": true
}
```
### 2. As a user, I need an API to retrieve the friends list for an email address.
Uri:
```
/api/user/get_friend
```
The API should receive the following JSON request:
```
    {
        "email": "andy@example.com"
    }
```
The API should return the following JSON response on success:
```
{
    "success": true,
    "friends" :
    [
        "john@example.com"
    ],
    "count" : 1
    }
```
### 3. As a user, I need an API to retrieve the common friends list between two email addresses.
Uri:
```
/api/user/get_friend_common
```
The API should receive the following JSON request:
```
{
    "friends":
    [
        "andy@example.com"
        "john@example.com"
    ]
}
```
The API should return the following JSON response on success:
```
    {
        "success": true,
        "friends" :
        [
            "common@example.com"
        ],
        "count" : 1
    }
```

### 4. As a user, I need an API to subscribe to updates from an email address.
Uri:
```
/api/user/subscribe
```
The API should receive the following JSON request:
```
    {
        "requestor": "lisa@example.com",
        "target": "john@example.com"
    }
```
The API should return the following JSON response on success:
```
    {
        "success": true
    }
```

### 5. As a user, I need an API to block updates from an email address.
Uri:
```
/api/user/block
```

Suppose "andy@example.com" blocks "john@example.com":
The API should receive the following JSON request:
```
    {
        "requestor": "andy@example.com",
        "target": "john@example.com"
    }
```
The API should return the following JSON response on success:
```
    {
        "success": true
    }
```

### 6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
Uri:
```
/api/post/create
```

The API should receive the following JSON request:
```
    {
        "sender": "john@example.com",
        "text": "Hello World! kate@example.com"
    }
```
The API should return the following JSON response on success:
```
    {
        "success": true
        "recipients":
        [
            "lisa@example.com",
            "kate@example.com"
        ]
    }
```

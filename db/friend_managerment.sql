Script started on 2020-04-13 15:42:13+0700
At the root of this project and run the following command:
```
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
```
## Feature

### 1. As a user, I need an API to create a friend connection between two email addresses.

Uri
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
Uri
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
Uri
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
Uri
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
Uri
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
Uri
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
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ At the root of this project and run the following command:
At: command not found
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ ```
> go get -u github.com/go-sql-driver/mysql
> go get -u github.com/jinzhu/gorm
> ```
^C
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ git stat^C
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ ^C
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ ^C
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ ^C
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhungtran@hungtranpc[00m:[01;34m~/coding/golang/friend-manager[00m$ ^C
]0;hungtran@hungtranpc: ~/coding/golang/friend-manager[01;32mhu
package main

import (
	c "controllers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go server running")

	http.HandleFunc("/api/user/create_friend", c.UserCreateFriend)

	http.HandleFunc("/api/user/get_friend", c.UserGetFriends)

	http.HandleFunc("/api/user/get_friend_common", c.UserGetFriendCommon)

	http.HandleFunc("/api/user/subscribe", c.UserSubscribe)

	http.HandleFunc("/api/user/block", c.UserBlock)

	http.HandleFunc("/api/post/create", c.PostCreate)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Print(err)
	}
}

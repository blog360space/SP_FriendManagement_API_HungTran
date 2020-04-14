package main

import (
	"configs"
	c "controllers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go server running")

	http.HandleFunc("/api/user/create_friend", c.UserCreateFriend)

	http.HandleFunc("/api/user/get_friend", c.UserGetFriends)

	http.HandleFunc("/api/user/get_friend_common", c.UserGetFriendsCommon)

	http.HandleFunc("/api/user/subscribe", c.UserSubscribe)

	http.HandleFunc("/api/user/block", c.UserBlock)

	http.HandleFunc("/api/user/register", c.UserRegister)

	http.HandleFunc("/api/post/create", c.PostCreate)

	err := http.ListenAndServe(configs.APP_PORT, nil)
	if err != nil {
		fmt.Print(err)
	}
}

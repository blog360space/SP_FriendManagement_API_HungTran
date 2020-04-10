package main

import (
	"fmt"
	"net/http"
	c "controllers"
)

func main() {
	fmt.Println("Go server running")

	http.HandleFunc("/api/user/create_friend", c.User_CreateFriend)

	http.HandleFunc("/api/user/get_friend", c.User_GetFriend)

	http.HandleFunc("/api/user/get_friend_common", c.User_GetFriendCommon)

	http.HandleFunc("/api/user/subscribe", c.User_Subscribe)

	http.HandleFunc("/api/user/block", c.User_Block)

	http.HandleFunc("/api/post/create", c.Post_Create)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Print(err)
	}
}

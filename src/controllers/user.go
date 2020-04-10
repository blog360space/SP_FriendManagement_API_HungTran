package controllers

import (
	"net/http"
	"utils"
)

// Method: POST
// Create Friend
func User_CreateFriend(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_CreateFriend")

	utils.Respond(w, resp)
}

func User_GetFriend(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_GetFriend")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

func User_GetFriendCommon(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_GetFriendCommon")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

func User_Subscribe(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_Subscribe")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

func User_Block(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_Block")
	resp["friends"] = nil
	utils.Respond(w, resp)
}
package controllers

import (
	"encoding/json"
	"net/http"
	repo "repositories"
	"utils"
)

// UserCreateFriendRequest Request struct Create Friend between 2 existing emails
type UserCreateFriendRequest struct {
	Friends [2]string `json:"friends"`
}

// UserCreateFriend Create Friend between 2 existing emails
// Method: POST
// Example request { "friends": ["andy@example.com","john@example.com"]}
func UserCreateFriend(w http.ResponseWriter, r *http.Request) {
	defer utils.DbClose()

	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp)
		return
	}

	requestData := &UserCreateFriendRequest{}
	err := json.NewDecoder(r.Body).Decode(requestData)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp)
		return
	}

	emails := requestData.Friends
	if len(emails) != 2 {
		resp := utils.Message(false, "Friends must have 2 emails")
		utils.Respond(w, resp)
		return
	}

	_, err1 := repo.UserCreateFriend(emails[0], emails[1])
	if err1 != nil {
		resp := utils.Message(false, err1.Error())
		utils.Respond(w, resp)
		return
	}

	resp := utils.Message(true, "")
	utils.Respond(w, resp)
}

// UserGetFriend Get Friend
func UserGetFriend(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_GetFriend")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

// UserGetFriendCommon Get common friend between 2 users
func UserGetFriendCommon(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_GetFriendCommon")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

// UserSubscribe Subsribe by email
func UserSubscribe(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_Subscribe")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

// UserBlock User Block
func UserBlock(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_Block")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

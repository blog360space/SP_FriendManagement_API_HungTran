package controllers

import (
	"encoding/json"
	"models"
	"net/http"
	repo "repositories"
	"utils"
)

// UserCreateFriendRequest Request struct Create Friend between 2 existing emails
type UserCreateFriendRequest struct {
	Friends []string `json:"friends"`
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

// UserGetFriendsRequest struct when request UserGetFriends
type UserGetFriendsRequest struct {
	Email string `json:"email"`
}

// UserGetFriends Get Friend
// Example request {"email": "andy@example.com"}
func UserGetFriends(w http.ResponseWriter, r *http.Request) {
	defer utils.DbClose()

	requestStruct := &UserGetFriendsRequest{}
	err := json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp)
		return
	}

	email := requestStruct.Email
	users, err := repo.UserGetFriendsByEmail(email)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	var emails []string
	i := 0
	count := len(users)
	for i = 0; i < count; i++ {
		emails = append(emails, users[i].Email)
	}

	resp := utils.Message(true, "")
	resp["friends"] = emails
	resp["count"] = count
	utils.Respond(w, resp)
}

// UserGetFriendsCommonRequest Request struct Get common friend between 2 emails
// Use at UserGetFriendsCommon
type UserGetFriendsCommonRequest struct {
	Friends []string `json:"friends"`
}

// UserGetFriendsCommon Get common friend between 2 users
// Example request { "friends": ["andy@example.com","john@example.com"]}
func UserGetFriendsCommon(w http.ResponseWriter, r *http.Request) {
	// Close DB
	defer utils.DbClose()
	var err error
	var count, i int

	// Decode request into struct UserGetFriendsCommonRequest
	requestStruct := &UserGetFriendsCommonRequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp)
		return
	}

	// Validate reuqest
	count = len(requestStruct.Friends)
	if count != 2 {
		resp := utils.Message(false, "Friends must have 2 emails.")
		utils.Respond(w, resp)
		return
	}

	friendEmails := requestStruct.Friends
	// Call repository to get data
	users, err := repo.UserGetFriendsCommon(friendEmails[0], friendEmails[1])

	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	var emails []string
	count = len(users)

	if count > 0 {
		for i = 0; i < count; i++ {
			emails = append(emails, users[i].Email)
		}
	}

	resp := utils.Message(true, "")
	resp["friends"] = emails
	resp["count"] = count
	utils.Respond(w, resp)
}

// UserSubscribeRequest use at UserSubscribe
type UserSubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// UserSubscribe Subsribe by email
// Example: {"requestor": "lisa@example.com","target": "john@example.com"}
func UserSubscribe(w http.ResponseWriter, r *http.Request) {
	// Close DB
	defer utils.DbClose()
	var err error
	var relationship models.Relationship

	// Decode request into struct UserSubscribeRequest
	requestStruct := &UserSubscribeRequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp)
		return
	}

	relationship, err = repo.UserSubscribe(requestStruct.Requestor, requestStruct.Target)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	if relationship.ID == 0 {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	resp := utils.Message(true, "")
	utils.Respond(w, resp)
}

// UserBlock User Block
func UserBlock(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "User_Block")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

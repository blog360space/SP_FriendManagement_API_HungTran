package controllers

import (
	"encoding/json"
	//"fmt"
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
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	requestData := &UserCreateFriendRequest{}
	err := json.NewDecoder(r.Body).Decode(requestData)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	emails := requestData.Friends
	if len(emails) != 2 {
		resp := utils.Message(false, "Friends must have 2 emails")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	var u1, u2 models.User
	u1, err = repo.UserGetByEmail(emails[0])
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	u2, err = repo.UserGetByEmail(emails[1])
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	_, err1 := repo.UserCreateFriend(u1, u2)
	if err1 != nil {
		resp := utils.Message(false, err1.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	resp := utils.Message(true, "")
	utils.Respond(w, resp, http.StatusOK)
}

// UserGetFriendsRequest struct when request UserGetFriends
type UserGetFriendsRequest struct {
	Email string `json:"email"`
}

// UserGetFriends Get Friend
// Example request {"email": "andy@example.com"}
func UserGetFriends(w http.ResponseWriter, r *http.Request) {
	defer utils.DbClose()
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	requestStruct := &UserGetFriendsRequest{}
	err := json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	email := requestStruct.Email
	users, err := repo.UserGetFriendsByEmail(email)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
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
	utils.Respond(w, resp, http.StatusOK)
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
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	var err error
	var count, i int

	// Decode request into struct UserGetFriendsCommonRequest
	requestStruct := &UserGetFriendsCommonRequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	// Validate reuqest
	count = len(requestStruct.Friends)
	if count != 2 {
		resp := utils.Message(false, "Friends must have 2 emails.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	friendEmails := requestStruct.Friends
	var u1, u2 models.User
	u1, err = repo.UserGetByEmail(friendEmails[0])
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	u2, err = repo.UserGetByEmail(friendEmails[1])
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	// Call repository to get data
	users, err := repo.UserGetFriendsCommon(u1, u2)

	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
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
	utils.Respond(w, resp, http.StatusOK)
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
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	var err error
	var relationship models.Relationship
	var userRequestor, userTarget models.User

	// Decode request into struct UserSubscribeRequest
	requestStruct := &UserSubscribeRequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	userRequestor, err = repo.UserGetByEmail(requestStruct.Requestor)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	userTarget, err = repo.UserGetByEmail(requestStruct.Target)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	relationship, err = repo.UserSubscribe(userRequestor, userTarget)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	if relationship.ID == 0 {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	resp := utils.Message(true, "")
	utils.Respond(w, resp, http.StatusOK)
}

// UserBlockequest use at UserBlock
type UserBlockequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// UserBlock As a user, I need an API to block updates from an email address.
func UserBlock(w http.ResponseWriter, r *http.Request) {
	// Close DB
	defer utils.DbClose()
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	var err error
	var userRequestor, userTarget models.User

	// Decode request into struct UserSubscribeRequest
	requestStruct := &UserBlockequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	userRequestor, err = repo.UserGetByEmail(requestStruct.Requestor)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	userTarget, err = repo.UserGetByEmail(requestStruct.Target)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	_, err = repo.UserBlock(userRequestor, userTarget)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	resp := utils.Message(true, "")
	utils.Respond(w, resp, http.StatusOK)
}

// UserRegisterRequest use at UserRegister
type UserRegisterRequest struct {
	Email string `json:"email"`
}
// UserRegister Register new user
func UserRegister(w http.ResponseWriter, r *http.Request)  {
	defer utils.DbClose()
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	var err error
	var user models.User

	requestStruct := &UserRegisterRequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	email := requestStruct.Email
	user, err = repo.UserRegister(email)
	if (err != nil) {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp, http.StatusBadRequest)
		return
	}

	resp := utils.Message(true, "")
	resp["email"] = user.Email
	utils.Respond(w, resp, http.StatusOK)
}

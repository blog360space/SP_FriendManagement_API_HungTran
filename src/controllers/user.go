package controllers

import (
	"net/http"
	"utils"
	"encoding/json"
	repo "repositories"
)

/**
 * Request struct Create Friend between 2 existing emails
 */
type User_CreateFriend_Request struct {
	Friends [2]string `json:"friends"`
}

/**
 * Create Friend between 2 existing emails
 * Method: POST
 * Example request { "friends": ["andy@example.com","john@example.com"]}
 */
func User_CreateFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp)
		return
	}

	requestData := &User_CreateFriend_Request{}
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

	_, err1 := repo.User_CreateFriend(emails[0], emails[1])
	if err1 != nil {
		resp := utils.Message(false, err1.Error())
		utils.Respond(w, resp)
		return
	}

	resp := utils.Message(true, "")
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
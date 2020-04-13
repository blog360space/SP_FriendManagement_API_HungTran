package controllers

import (
	"encoding/json"
	"models"
	"net/http"
	repo "repositories"
	"utils"
)

// PostCreateRequest struct Create post request
// Used for PostCreate
type PostCreateRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// PostCreate : Create Post
// Method: POST
// Create Post
func PostCreate(w http.ResponseWriter, r *http.Request) {
	defer utils.DbClose()
	if r.Method != "POST" {
		resp := utils.Message(false, "Invalid method")
		utils.Respond(w, resp)
		return
	}

	var err error
	var sender models.User
	var post models.Post
	var recipients []models.User

	requestStruct := &PostCreateRequest{}
	err = json.NewDecoder(r.Body).Decode(requestStruct)
	if err != nil {
		resp := utils.Message(false, "Error while decoding request body.")
		utils.Respond(w, resp)
		return
	}

	sender, err = repo.UserGetByEmail(requestStruct.Sender)

	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	if requestStruct.Text == "" {
		resp := utils.Message(false, "Post content is empty")
		utils.Respond(w, resp)
		return
	}

	post, err = repo.PostCreatePost(sender, requestStruct.Text)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	recipients, err = repo.PostGetPostRecipients(sender, post)
	if err != nil {
		resp := utils.Message(false, err.Error())
		utils.Respond(w, resp)
		return
	}

	var recipientEmails []string
	i := 0
	count := len(recipients)
	for i = 0; i < count; i++ {
		recipientEmails = append(recipientEmails, recipients[i].Email)
	}

	resp := utils.Message(true, "")
	resp["recipients"] = recipientEmails
	utils.Respond(w, resp)
}

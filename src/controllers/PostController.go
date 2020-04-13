package controllers

import (
	"net/http"
	"utils"
)

// PostCreate : Create Post
// Method: POST
// Create Friend
func PostCreate(w http.ResponseWriter, r *http.Request) {

	resp := utils.Message(true, "Post_Create")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

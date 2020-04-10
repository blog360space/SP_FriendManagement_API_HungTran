package controllers

import (
	"net/http"
	"utils"
)

// Method: POST
// Create Friend
func Post_Create(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "Post_Create")
	resp["friends"] = nil
	utils.Respond(w, resp)
}

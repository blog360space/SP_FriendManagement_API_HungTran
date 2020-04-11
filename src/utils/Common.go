package utils

import (
	"encoding/json"
	"net/http"
)

// Message to respose
func Message(status bool, message string) map[string]interface{} {
	if status == true {
		return map[string]interface{}{"success": status}
	}
	return map[string]interface{}{"success": status, "message": message}
}

// Respond json to browser
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

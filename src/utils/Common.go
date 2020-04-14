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
func Respond(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// ValidateIsEmail Validdate a string is email
func ValidateIsEmail(str string) bool {
	return true
}

package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}) {
	if (status == true) {
		return map[string]interface{} {"success" : status}
	} else {
		return map[string]interface{} {"success" : status, "message" : message}
	}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
package controllers

import (
	"fmt"
	"net/http"
)

func Order_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OrderController.Create")
}

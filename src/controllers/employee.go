package controllers

// import (
// 	"net/http"
// 	"repositories"
// 	"models"
// 	"utils"
// )

// func Employee_Index(w http.ResponseWriter, r *http.Request)  {

// 	type MyRes struct {
// 		Success string
// 		Count int
// 		Employees []models.Employee
// 	}

// 	arrEmployee, err := repositories.Employee_GetAll()

// 	if (err != nil) {
// 		resp := utils.Message(true, err.Error())
// 		utils.Respond(w, resp)
// 	} else {
// 		resp := utils.Message(true, "Success")
// 		resp["employee"] = arrEmployee
// 		utils.Respond(w, resp)
// 	}
// }

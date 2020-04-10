package main

import (
	"fmt"
	"net/http"
	controllers "controllers"
)

func main() {
	fmt.Println("Hello world");

	http.HandleFunc("/", controllers.Employee_Index)
	// repository.Order_GetAll();

	http.ListenAndServe(":8080", nil)
	fmt.Println("Hello world");


	// db := utils.DbConn();

	// var sql string = "SELECT * FROM employee ORDER BY id DESC"
	// selDb, err := db.Query(sql)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// emp := models.Employee{}
	// res := []models.Employee{}

	// for selDb.Next() {
	// 	var id int
	// 	var name, city string
	// 	err = selDb.Scan(&id, &name, &city)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	emp.Id = id
	// 	emp.Name = name
	// 	emp.City = city

	// 	res = append(res, emp)
	// }
	// // tmpl.ExecuteTemplate(w, "Index", res)
	// fmt.Println(res)
	// defer db.Close()



	// order := models.Order{1, "sssssss"}
	// fmt.Println(order.Name)


	// fmt.Printf("Type of order %T", order)
	// var c = controllers.OrderController{}
	// (&c).Create();
}

package repositories

import (
	"models"
	"utils"
	// "errors"
)

func Employee_GetAll() ([]models.Employee, error){

	db := utils.DbConn();

	var sql string = "SELECT * FROM employee ORDER BY id DESC"
	selDb, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	emp := models.Employee{}
	res := []models.Employee{}

	for selDb.Next() {
		var id int
		var name, city string
		err = selDb.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}
	// tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

	// err = errors.New("Test error")

	return res, err
}
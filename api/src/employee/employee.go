package employee

import (
	"database/sql"
    // "app/typefile"
	// "encoding/json"
)

type Employee struct{
    Id   int    `json:"id"`
    Name string `json:"name"`
}

type Employees []Employee

func ReadAll(db *sql.DB) Employees{
	var employees Employees
	rows, err := db.Query("select * from employee;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
        employee := Employee{}
        err = rows.Scan(&employee.Id, &employee.Name)
		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}

	rows.Close()
    return employees
}

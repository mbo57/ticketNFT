package employee

import (
	"database/sql"
    "app/typefile"
	// "encoding/json"
)

func ReadAll(db *sql.DB) typefile.Employees{
	var employees typefile.Employees
	rows, err := db.Query("select * from employee;")
	if err != nil {
		panic(err)
	}
    // fmt.Println(rows)
	for rows.Next() {
        employee := typefile.Employee{}
        err = rows.Scan(&employee.Id, &employee.Name)
		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}
    // jsonemployees, _ := json.Marshal(employees)
	rows.Close()
    return employees
}

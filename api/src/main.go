package main

import (
    "fmt"
    "log"
    "net/http"
	"database/sql"
	"app/employee"
    "app/typefile"
	"encoding/json"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = "root:password@tcp(db:3306)/sample?charset=utf8&parseTime=true"

	return open(path, 100)
}


func apiserver(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprint(w, "Welcome to the HomePage!")
    fmt.Println("Endopoint Hit: homePage")
	db := connectDB()
	defer db.Close()

	var employees typefile.Employees
    employees = employee.ReadAll(db)
    
	json.NewEncoder(w).Encode(employees)
}

func handleRequests() {
    http.HandleFunc("/", apiserver)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    handleRequests()
}

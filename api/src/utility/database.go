package utility

import (
    "fmt"
    "log"
    "time"
	"database/sql"
	
    _ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init(){
    var err error
	var path string = "root:password@tcp(db:3306)/sample?charset=utf8&parseTime=true"
    if Db, err = sql.Open("mysql", path); err != nil {
		log.Fatal("db open error:", err.Error())
	}
	checkConnect(100)

	fmt.Println("db connected!!")
}

func checkConnect(count uint) {
	var err error
	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		checkConnect(count)
	}
}

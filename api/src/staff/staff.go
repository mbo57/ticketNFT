package staff

import (
    "fmt"
    // "errors"
    "net/http"
	// "database/sql"
	"encoding/json"
    "app/utility"
)


type Staff struct{
    Id       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
}

type Staffs []Staff

func Show(w http.ResponseWriter, r *http.Request) {
	var staffs Staffs
    if r.Method == "GET" {
        query := "select * from staff"
        if err := r.ParseForm(); err != nil {
            fmt.Println("errorだよ")
        }
        length := len(r.Form)
        if length > 0 {
            query += " where "
            cnt := 0
            for k, v := range r.Form {
                cnt += 1
                query = query + fmt.Sprintf("%s='%s'", k, v[0])
                if length < cnt {
                    query += " and "
                }
            }
        }
        query += ";"
        fmt.Println(query)
	    rows, err := utility.Db.Query(query)
	    if err != nil {
	    	panic(err)
	    }

	    for rows.Next() {
            staff := Staff{}
            err = rows.Scan(&staff.Id,
                            &staff.Email,
                            &staff.Password,
                            &staff.Name)
	    	if err != nil {
	    		panic(err)
	    	}
            fmt.Println(staff)
	    	staffs = append(staffs, staff)
	    }

	    rows.Close()
    }
	json.NewEncoder(w).Encode(staffs)
}


func Create(w http.ResponseWriter, r *http.Request) {
    var ans Staffs
    if r.Method == "GET" {
        email := r.FormValue("email")
        password := r.FormValue("password")
        name := r.FormValue("name")

        in, err := utility.Db.Prepare("INSERT INTO staff(email, password, name) VALUES(?, ?, ?)")

        if err != nil {
        	fmt.Println("データベース接続失敗")
        	panic(err.Error())
        } else {
        	fmt.Println("データベース接続成功")
        }
        
        // defer utility.Db.Close()
        
        result, err := in.Exec(email, password, name)
        
        if err != nil {
        	panic(err.Error())
        }
        
        lastId, err := result.LastInsertId()
        
        if err != nil {
        	panic(err.Error())
        }
        
        fmt.Println(lastId)
    }
	json.NewEncoder(w).Encode(ans)
}

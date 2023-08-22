package staff

import (
    "fmt"
    "net/http"
	"encoding/json"
    "app/utility"
    "regexp"

    "golang.org/x/exp/slices"
)


type Staff struct{
    Id       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
}
var columns = []string{"id", "email", "password", "name"}

type Staffs []Staff

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func ResponseWriter(w http.ResponseWriter, r *http.Request, code int, mess string) {
    var res Response
    res.Code = code
    res.Message = mess
    json.NewEncoder(w).Encode(res)
}

func Entry(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        var res Response
        res.Code = 423
        res.Message = err.Error()
	    json.NewEncoder(w).Encode(res)
        return
    }
    switch r.Method {
    case "GET":
        Read(w, r)
    case "POST":
        Create(w, r)
    case "PATCH":
        Update(w, r)
    case "DELETE":
        Delete(w, r)
    default:
        ResponseWriter(w, r, 422, fmt.Sprintf("Methods error : un used '%s' method", r.Method))
    }
}


func Read(w http.ResponseWriter, r *http.Request) {
	var staffs Staffs
    re, err := regexp.Compile(`/staffs/([0-9]+).*` )
    if err != nil {
        ResponseWriter(w, r, 424, err.Error())
        return
    }
    
    matches := re.FindAllStringSubmatch(r.URL.Path, -1)
    
    if matches != nil {
        rows, err := utility.Db.Query("select * from staff where id = ?", matches[0][1])
	    if err != nil {
            ResponseWriter(w, r, 423, err.Error())
            return
	    }

	    for rows.Next() {
            staff := Staff{}
            err = rows.Scan(&staff.Id,
                            &staff.Email,
                            &staff.Password,
                            &staff.Name)
	    	if err != nil {
                ResponseWriter(w, r, 423, err.Error())
                return
	    	}
	    	staffs = append(staffs, staff)
	    }
	    rows.Close()
    } else {
        query := "select * from staff"
        
        length := len(r.Form)
        var list []any
        i := 0
        for k, v := range r.Form {
            if i == 0{
                query += " where"
            }
            if !slices.Contains(columns, k) {
                ResponseWriter(w, r, 422, fmt.Sprintf("'%s' column not found", k))
                return
            }
            list = append(list, v[0])
            query += fmt.Sprintf(" %s = ?", k)
            if i != length - 1 {
                query += " and"
            }
            i += 1
        }
        query += ";"
        rows, err := utility.Db.Query(query, list...)
	    if err != nil {
            ResponseWriter(w, r, 423,  err.Error())
            return
	    }

	    for rows.Next() {
            staff := Staff{}
            err = rows.Scan(&staff.Id,
                            &staff.Email,
                            &staff.Password,
                            &staff.Name)
	    	if err != nil {
                ResponseWriter(w, r, 423,  err.Error())
                return
	    	}
	    	staffs = append(staffs, staff)
	    }

	    rows.Close()
        
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
	json.NewEncoder(w).Encode(staffs)
}


func Create(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")
    password := r.FormValue("password")
    name := r.FormValue("name")

    result, err := utility.Db.Exec("INSERT INTO staff(email, password, name) VALUES(?, ?, ?)", email, password, name)
    
    if err != nil {
        ResponseWriter(w, r, 423,  err.Error())
        return
    }
    
    lastId, err := result.LastInsertId()
    
    if err != nil {
        ResponseWriter(w, r, 423,  err.Error())
        return
    }
    
    ResponseWriter(w, r, 200, fmt.Sprintf("sucsess : id = %d", lastId))
}

func Update(w http.ResponseWriter, r *http.Request) {
    re, err := regexp.Compile(`/staffs/([0-9]+).*` )
    if err != nil {
        ResponseWriter(w, r, 424, err.Error())
        return
    }
    
    matches := re.FindAllStringSubmatch(r.URL.Path, -1)

    if matches != nil {
        id := matches[0][1]
        query := "update staff set"
        
        length := len(r.Form)
        var list []any
        i := 0
        for k, v := range r.Form {
            if !slices.Contains(columns, k) {
                ResponseWriter(w, r, 422, fmt.Sprintf("'%s' column not found", k))
                return
            }
            list = append(list, v[0])
            query += fmt.Sprintf(" %s = ?", k)
            if i != length - 1 {
                query += " ,"
            }
            i += 1
        }

        list = append(list, id)

        query += " where id = ?;"
        result, err := utility.Db.Exec(query, list...)
	    if err != nil {
            ResponseWriter(w, r, 423, err.Error())
            return
	    }

        delete_num, err := result.RowsAffected()
        if err != nil {
            ResponseWriter(w, r, 423, err.Error())
            return
        }
        if delete_num == 1 {
            ResponseWriter(w, r, 200, "success update")
        } else if delete_num == 0 {
            ResponseWriter(w, r, 422, "fail update")
        }

    } else {
        ResponseWriter(w, r, 422, "Not found id")
    }
}

func Delete(w http.ResponseWriter, r *http.Request) {
    re, err := regexp.Compile(`/staffs/([0-9]+).*` )
    if err != nil {
        ResponseWriter(w, r, 424, err.Error())
        return
    }
    
    matches := re.FindAllStringSubmatch(r.URL.Path, -1)

    if matches != nil {
        id := matches[0][1]
        result, err := utility.Db.Exec("delete from staff where id = ?", id)
        if err != nil {
            ResponseWriter(w, r, 423, err.Error())
            return
        }
        delete_num, err := result.RowsAffected()
        if err != nil {
            ResponseWriter(w, r, 424, err.Error())
            return
        }
        if delete_num == 1 {
            ResponseWriter(w, r, 200, "success")
        } else if delete_num == 0 {
            ResponseWriter(w, r, 422, fmt.Sprintf("Not find id = %s", id))
        }
    } else {
        ResponseWriter(w, r, 422, "Not find id")
    }
}

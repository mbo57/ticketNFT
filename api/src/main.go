package main

import (
    "log"
    "net/http"
	"app/staff"
)


// func apiserver(w http.ResponseWriter, r *http.Request) {
//     staffs := staff.ReadAll(db)
//     
// 	json.NewEncoder(w).Encode(staffs)
// }

func main() {
    http.HandleFunc("/staff/show", staff.Show)
    http.HandleFunc("/staff/create", staff.Create)
    // http.HandleFunc("/staff/update", staff.Update)
    // http.HandleFunc("/staff/delete", staff.Delete)
    log.Fatal(http.ListenAndServe(":8000", nil))
    // handleRequests()
}

package main

import (
    "log"
    "net/http"
	"app/staff"
)

func main() {
    // http.HandleFunc("/staffs", staff.Entry)
    http.HandleFunc("/staffs/", staff.Entry)
    log.Fatal(http.ListenAndServe(":8000", nil))
    // handleRequests()
}

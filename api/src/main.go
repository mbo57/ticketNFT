package main

import (
    "log"
    "net/http"
	"app/staff"
    "app/event"
    "app/cast"
)

func main() {
    http.HandleFunc("/staff/", staff.Entry)
    http.HandleFunc("/event/", event.Entry)
    http.HandleFunc("/cast/", cast.Entry)
    log.Fatal(http.ListenAndServe(":8000", nil))
    // handleRequests()
}

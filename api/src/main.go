package main

import (
    "log"
    "net/http"
	// "app/staff"
    // "app/event"
    "app/crud"
)

func main() {
    http.HandleFunc("/staff/", crud.Entry)
    http.HandleFunc("/event/", crud.Entry)
    http.HandleFunc("/eventcategory/", crud.Entry)
    http.HandleFunc("/cast/", crud.Entry)
    // http.HandleFunc("/event/", event.Entry)
    log.Fatal(http.ListenAndServe(":8000", nil))
    // handleRequests()
}

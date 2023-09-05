package main

import (
	"app/crud"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/staff/", crud.Entry)
	http.HandleFunc("/event/", crud.Entry)
	http.HandleFunc("/eventcategory/", crud.Entry)
	http.HandleFunc("/cast/", crud.Entry)

	log.Fatal(http.ListenAndServe(":8000", nil))

	// ↓　共存むりかも　JWT
	// e := router.Init()
	// e.Logger.Fatal(e.Start(":8000"))

}

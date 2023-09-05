package main

import (
	"app/router"
)

func main() {

	// http.HandleFunc("/staff/", crud.Entry)
	// http.HandleFunc("/event/", crud.Entry)
	// http.HandleFunc("/eventcategory/", crud.Entry)
	// http.HandleFunc("/cast/", crud.Entry)

	e := router.Init()

	e.Logger.Fatal(e.Start(":8000"))

	// log.Fatal(http.ListenAndServe(":8000", nil))
}

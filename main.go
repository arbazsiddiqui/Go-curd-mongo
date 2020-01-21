package main

import (
	"net/http"
	"todoApp/routes"
)

func main() {
	r := routes.Router()
	http.ListenAndServe(":8080", r)
}

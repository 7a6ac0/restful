package main

import (
	"net/http"

	"github.com/7a6ac0/restful/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":8080", r)
}

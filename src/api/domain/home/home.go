package home

import (
	"fmt"
	"net/http"
)

// HomePage func to test if the API is up and running.
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

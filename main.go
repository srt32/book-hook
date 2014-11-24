package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/purchase", newPurchase)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + os.Getenv("PORT"))
}

func newPurchase(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "From the handler!")

	// parse the params

	// get the github username and product

	// add the user to the GH team
}

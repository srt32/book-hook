package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	n := negroni.Classic()
	n.UseHandler(Handlers())
	n.Run(":" + os.Getenv("PORT"))
}

func Handlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/purchase", newPurchase)

	return mux
}

func newPurchase(w http.ResponseWriter, req *http.Request) {
	data := json.NewDecoder(req.Body)
	var purchase Purchase

	if err := data.Decode(&purchase); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to decode json: %s", err),
			http.StatusBadRequest,
		)
		return
	}

	UserName := purchase.UserName
	productId := purchase.ProductId
	teamId := productId // requires mapping

	err := addUserToTeam(UserName, teamId)

	if err != nil {
		log.Fatal("Error adding user to team", err)
	}
}

type Purchase struct {
	UserName  string `json:"email"`
	ProductId int    `json:"product_id"`
}

func addUserToTeam(UserName string, teamId int) error {
	// add the user to the GH team
	fmt.Print("adding user")

	return nil
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

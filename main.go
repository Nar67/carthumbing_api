package main

import (
	"fmt"
	"hackupc2017w/carthumbing_api/api"
	"hackupc2017w/carthumbing_api/models"
	"hackupc2017w/carthumbing_api/alg"
	"net/http"
)

func main() {
	alg.TestArcGisApi()
	alg.TestRoute()
	return

	dbUser := "carthumbing_user"
	dbPass := "ohphahRohfohZoh6"
	dbName := "carthumbing_db"
	dbHost := "localhost"
	dbPort := "9432"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(dbUrl)

	models.InitDB(dbUrl)

	router := api.NewRouter()
	fmt.Println("Running server on: 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", router)
}

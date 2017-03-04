package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"hackupc2017w/carthumbing_api/models"
)

func CreateSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---Create Search")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//s := string(body[:])
	//fmt.Println(s)
	var rd models.Route
	err = json.Unmarshal(body, &rd)
	if err != nil {
		panic(err)
	}
	fmt.Println(rd)
	b, err := json.Marshal(rd)
	if err != nil {
		panic(err)
	}

	id, err := models.CreateRoute(&rd)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
	//query := db.QueryRow("INSERT INTO searches (VALUES{UserID | $1}, {Origin | $2}, {Destination | $3}, {Date | $4})", Userid, Origin, Destination, Date)
}

// func Publish(w http.ResponseWriter, r *http.Request) {
// 	Userid : = 123451
// 	Origin := "coord"
// 	Destination := "coord2"
// 	Date := "dayX"
// 	query := db.QueryRow("INSERT INTO routes
// 						   (VALUES{UserID | $1}, {Origin | $2}, {Destination | $3}, {Date | $4})", Userid, Origin, Destination, Date)

// }
// func GetRoute(w http.ResponseWriter, r *http.Request) {
// 	Userid : = 123451
// 	Origin := "coord"
// 	Destination := "coord2"
// 	Date := "dayX"
// 	query := db.QueryRow("SELECT r.UserID r.Date
// 						  FROM routes r searches s
// 						  WHERE r.Date = s.Date AND r.Origin = s.Origin AND r.Destination = s.Destination")

// }

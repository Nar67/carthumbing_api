package models

import "time"

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Route struct {
	UserUuid    string    `json:"user_id"`
	Origin      Coord     `json:"origin"`
	Destination Coord     `json:"destination"`
	Date        time.Time `json:"date"`
}

func CreateRoute(r *Route) (int64, error) {
	query := `INSERT INTO route as r (user_uuid, origin, destination, selected_date)
			  VALUES ($1, point($2,$3), point($4, $5), $6)
			  RETURNING id;`
	var id int64
	err := db.QueryRow(query, r.UserUuid, r.Origin.Lat, r.Origin.Lon, r.Destination.Lat, r.Destination.Lon, r.Date).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, err
}

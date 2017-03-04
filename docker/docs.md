

## Create Search Route

curl -XPOST localhost:8080/api/search \
-H "Content-Type: application/json" \
-d '{"user_id": "182379ryh23r7", "origin": {"lat":2.5, "lon":3.5}, "destination": {"lat":2.7, "lon":13.6}, "date": "2016-08-08T21:35:14.052975+02:00"}'

## Migrations

postgres://carthumbing_user:ohphahRohfohZoh6@localhost:9432/carthumbing_db?sslmode=disable
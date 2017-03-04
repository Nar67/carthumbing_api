# Image name: hackupc2017w/carthumbing_api
FROM golang:1.7

ADD docker/wait-for-it.sh /usr/local/bin/wait-for-it.sh

RUN mkdir -p /go/src/hackupc2017w/carthumbing_api
WORKDIR /go/src/hackupc2017w/carthumbing_api

COPY . /go/src/hackupc2017w/carthumbing_api


RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install

# get dependencies
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go get github.com/mattes/migrate
RUN go get github.com/gorilla/handlers
RUN go get github.com/twinj/uuid

ENV PORT 8080

EXPOSE 8080

CMD env && /usr/local/bin/wait-for-it.sh -t 300 carthumbing_db:5432 \
    # && migrate -url postgres://carthumbing_user:ohphahRohfohZoh6@carthumbing_db:5432/carthumbing?sslmode=disable -path ./models/migrations/ up \
    && go run main.go

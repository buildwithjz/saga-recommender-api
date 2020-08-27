FROM golang:1.15

RUN go get go.mongodb.org/mongo-driver/mongo
RUN mkdir -p /saga-recommender-api

ADD main.go query-db.go /saga-recommender-api/

WORKDIR /saga-recommender-api

RUN go build .
EXPOSE 8090

ENTRYPOINT ["./saga-recommender-api"]
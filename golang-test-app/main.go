package main

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hl-hw5/handlers"
	"net/http"
	"time"
)

const (
	// TODO move to env
	mongoConnection = "mongodb://mongodb:27017"
	httpPort        = ":8080"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// setup test mongo
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnection))
	err = mongoClient.Ping(ctx, readpref.Primary())

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	handler := handlers.NewHandlers(mongoClient)

	// setup test server and endpoint
	r := mux.NewRouter()
	r.HandleFunc("/", handler.InsertTestData)
	http.Handle("/", r)

	log.Infof("Starting HTTP server on %s", httpPort)
	srv := &http.Server{
		Addr:         httpPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type Handlers struct {
	mongoClient *mongo.Client
}

func NewHandlers(mongoClient *mongo.Client) Handlers {
	return Handlers{
		mongoClient,
	}
}

func (h *Handlers) InsertTestData(w http.ResponseWriter, r *http.Request) {
	if err := h.insertToMongo(r.Context(), r.URL.Query().Get("r")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status": false}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": true}`))
}

func (h *Handlers) insertToMongo(ctx context.Context, r string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	collection := h.mongoClient.Database("testing").Collection("numbers")
	_, err := collection.InsertOne(ctx, bson.D{{"value", r}})

	return err
}

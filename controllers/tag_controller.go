package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/thilina01/kb-api-go/config"
	"github.com/thilina01/kb-api-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tag.ID = primitive.NewObjectID()

	collection := config.DB.Collection("tags")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tag)
}

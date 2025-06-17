package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/thilina01/kb-api-go/config"
	"github.com/thilina01/kb-api-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateTagsExist(tagIDs []primitive.ObjectID) bool {
	collection := config.DB.Collection("tags")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$in": tagIDs}}
	count, err := collection.CountDocuments(ctx, filter)
	return err == nil && int(count) == len(tagIDs)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !validateTagsExist(article.Tags) {
		http.Error(w, "One or more tag IDs are invalid", http.StatusBadRequest)
		return
	}

	article.ID = primitive.NewObjectID()
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DB.Collection("articles").InsertOne(ctx, article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	tagFilter := r.URL.Query().Get("tag")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit

	matchStage := bson.M{}
	if tagFilter != "" {
		tagID, err := primitive.ObjectIDFromHex(tagFilter)
		if err == nil {
			matchStage = bson.M{"tags": tagID}
		}
	}

	pipeline := []bson.M{
		{"$match": matchStage},
		{"$sort": bson.M{"createdAt": -1}},
		{"$skip": int64(skip)},
		{"$limit": int64(limit)},
		{"$lookup": bson.M{
			"from":         "tags",
			"localField":   "tags",
			"foreignField": "_id",
			"as":           "tagDetails",
		}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var enriched []bson.M
	if err := cursor.All(ctx, &enriched); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(enriched)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	idHex := strings.TrimPrefix(r.URL.Path, "/articles/")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var updatedData models.Article
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !validateTagsExist(updatedData.Tags) {
		http.Error(w, "One or more tag IDs are invalid", http.StatusBadRequest)
		return
	}

	updatedData.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"title":     updatedData.Title,
			"content":   updatedData.Content,
			"tags":      updatedData.Tags,
			"updatedAt": updatedData.UpdatedAt,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.DB.Collection("articles").UpdateByID(ctx, id, update)
	if err != nil || res.MatchedCount == 0 {
		http.Error(w, "Article not found or update failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{"message": "Article updated"})
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	idHex := strings.TrimPrefix(r.URL.Path, "/articles/")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.DB.Collection("articles").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || res.DeletedCount == 0 {
		http.Error(w, "Article not found or delete failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bson.M{"message": "Article deleted"})
}

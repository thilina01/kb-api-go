package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/thilina01/kb-api-go/config"
	"github.com/thilina01/kb-api-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedTags() {
	tags := []string{"go", "mongodb", "docker", "api", "backend"}
	collection := config.DB.Collection("tags")

	for _, name := range tags {
		doc := models.Tag{
			ID:   primitive.NewObjectID(),
			Name: name,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := collection.InsertOne(ctx, doc)
		if err != nil {
			fmt.Println("⚠️ Tag already exists or failed:", name)
		} else {
			fmt.Println("✅ Seeded tag:", name)
		}
	}
}

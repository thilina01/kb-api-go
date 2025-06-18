package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/thilina01/kb-api-go/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnsureTextIndex() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	index := bson.D{
		{"title", "text"},
		{"content", "text"},
	}

	model := mongo.IndexModel{
		Keys:    index,
		Options: options.Index().SetName("text_index"),
	}

	_, err := config.DB.Collection("articles").Indexes().CreateOne(ctx, model)
	if err != nil {
		fmt.Println("❌ Failed to create text index:", err)
	} else {
		fmt.Println("✅ Text index created (if not exists)")
	}
}

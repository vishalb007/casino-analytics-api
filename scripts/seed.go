package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// connect mongo...

	users := make([]primitive.ObjectID, 500)
	for i := range users {
		users[i] = primitive.NewObjectID()
	}

	for i := 0; i < 2_000_000; i++ {
		_ = uuid.New().String()
		_ = users[rand.Intn(len(users))]
		// generate transactions (same as earlier)
	}
}
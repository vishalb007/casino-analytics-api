package main

import (
	"casino-analytics/internal/api"
	"casino-analytics/internal/cache"
	"casino-analytics/internal/db"
	"casino-analytics/internal/repository"
	"casino-analytics/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	mongo := db.NewMongo(os.Getenv("MONGO_URI"), os.Getenv("DB_NAME"))
	redis := cache.NewRedis(os.Getenv("REDIS_ADDR"))

	repo := repository.NewTransactionRepo(mongo)
	svc := service.NewAnalyticsService(repo, redis)

	handler := api.NewHandler(svc)

	r := gin.Default()
	api.RegisterRoutes(r, handler)

	r.Run(":" + os.Getenv("PORT"))
}
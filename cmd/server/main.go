package main

import (
	"github.com/sirupsen/logrus"

	"github.com/growthfolio/go-pricefeed/internal/config"
	"github.com/growthfolio/go-pricefeed/internal/postgres"
	"github.com/growthfolio/go-pricefeed/internal/redis"
)

func main() {
	logrus.Info("PriceFeed microservice started.")

	// Carrega configurações
	cfg := config.LoadConfig()

	// Inicializa Redis
	redisClient := redis.NewClient(cfg)

	if err := redis.PingRedis(redisClient.Redis); err != nil {
		logrus.WithError(err).Fatal("Redis connection failed")
	}

	// Inicializa Postgres
	_, err := postgres.NewDB(cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Postgres connection failed")
	}

	// TODO: Inicializar Binance listener, dispatcher, etc.
}

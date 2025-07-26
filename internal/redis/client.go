package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/growthfolio/go-pricefeed/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Redis *redis.Client
	TTL   time.Duration
}

func NewClient(cfg *config.Config) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		logrus.WithError(err).Fatal("Failed to connect to Redis")
	}
	return &Client{
		Redis: rdb,
		TTL:   GetTTL(cfg),
	}
}

func NewRedisClient(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,
		DB:       0,
	})
}

func GetTTL(cfg *config.Config) time.Duration {
	ttl, err := strconv.Atoi(cfg.CacheTTL)
	if err != nil {
		return 5 * time.Second // fallback
	}
	return time.Duration(ttl) * time.Second
}

func PingRedis(client *redis.Client) error {
	return client.Ping(context.Background()).Err()
}

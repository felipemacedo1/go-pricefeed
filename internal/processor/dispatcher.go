import (
    "github.com/growthfolio/go-pricefeed/internal/binance"
    "github.com/growthfolio/go-pricefeed/internal/redis"
    "github.com/growthfolio/go-pricefeed/internal/postgres"
    "github.com/sirupsen/logrus"
    "context"
    "fmt"


// StartDispatcher recebe atualizações de preço e salva no Redis e PostgreSQL
func StartDispatcher(
    updates <-chan binance.PriceUpdate,
    rds *redis.Client,
    db *postgres.DB,
    ttlSeconds int,
) {
    for update := range updates {
        // Salva no Redis
        key := fmt.Sprintf("price:%s", update.Symbol)
        err := rds.Redis.Set(context.Background(), key, update.Price, rds.TTL).Err()
        if err != nil {
            logrus.WithError(err).Error("Erro ao salvar preço no Redis")
        }

        // Salva no PostgreSQL
        _, err = db.Exec(
            "INSERT INTO price_feed_snapshot (symbol, price, ts) VALUES ($1, $2, NOW())",
            update.Symbol, update.Price,
        )
        if err != nil {
            logrus.WithError(err).Error("Erro ao salvar preço no PostgreSQL")
        }

        logrus.Infof("Processado %s: %.2f", update.Symbol, update.Price)
    }
}

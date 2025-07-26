package processor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/growthfolio/go-pricefeed/internal/binance"
	"github.com/growthfolio/go-pricefeed/internal/redis"
	"github.com/sirupsen/logrus"
)

// StartDispatcher recebe atualizações de preço e salva no Redis e PostgreSQL
func StartDispatcher(updates <-chan binance.PriceUpdate, rds *redis.Client, db *sql.DB, ttlSeconds int) {
	for update := range updates {
		// Salva no Redis
		key := fmt.Sprintf("price:%s", update.Symbol)
		err := rds.Redis.Set(context.Background(), key, update.ClosePrice, rds.TTL).Err()
		if err != nil {
			logrus.WithError(err).Error("Erro ao salvar preço no Redis")
		}

		// Salva no PostgreSQL
		_, err = db.Exec(
			`INSERT INTO price_history (
				symbol, timeframe, open_price, high_price, low_price, close_price, volume, timestamp
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			update.Symbol,
			update.Timeframe,
			update.OpenPrice,
			update.HighPrice,
			update.LowPrice,
			update.ClosePrice,
			update.Volume,
			update.Timestamp,
		)
		if err != nil {
			logrus.WithError(err).Error("Erro ao salvar histórico de preço no PostgreSQL")
		}

		logrus.Infof("Processado %s %s: O=%.2f H=%.2f L=%.2f C=%.2f V=%.2f T=%v",
			update.Symbol, update.Timeframe, update.OpenPrice, update.HighPrice, update.LowPrice, update.ClosePrice, update.Volume, update.Timestamp)
	}
}

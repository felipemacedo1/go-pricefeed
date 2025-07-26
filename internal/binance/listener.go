package binance

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type PriceUpdate struct {
	Symbol string  `json:"s"`
	Price  float64 `json:"c,string"`
}

// ListenPriceStream conecta ao websocket da Binance e envia atualizações de preço para o canal out
func ListenPriceStream(symbol string, out chan<- PriceUpdate) {
	url := "wss://stream.binance.com:9443/ws/" + symbol + "@ticker"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logrus.WithError(err).Error("Erro ao conectar ao websocket da Binance")
		return
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			logrus.WithError(err).Error("Erro ao ler mensagem do websocket")
			return
		}

		var update PriceUpdate
		if err := json.Unmarshal(message, &update); err != nil {
			logrus.WithError(err).Warn("Falha ao parsear mensagem de preço")
			continue
		}
		if update.Symbol != "" && update.Price != 0 {
			out <- update
		}
	}
}

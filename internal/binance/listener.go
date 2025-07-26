package binance

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type PriceUpdate struct {
	Symbol     string  `json:"s"`
	Timeframe  string  `json:"timeframe"`
	OpenPrice  float64 `json:"o,string"`
	HighPrice  float64 `json:"h,string"`
	LowPrice   float64 `json:"l,string"`
	ClosePrice float64 `json:"c,string"`
	Volume     float64 `json:"v,string"`
	Timestamp  int64   `json:"E"` // Event time (ms since epoch)
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
		// Valida todos os campos essenciais
		if update.Symbol != "" && update.ClosePrice != 0 && update.OpenPrice != 0 && update.HighPrice != 0 && update.LowPrice != 0 && update.Volume != 0 && update.Timestamp != 0 {
			// Preencher timeframe se necessário (exemplo: padrão 1m)
			if update.Timeframe == "" {
				update.Timeframe = "1m"
			}
			out <- update
		}
	}
}

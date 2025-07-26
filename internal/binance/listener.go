package binance

import (
	"encoding/json"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type PriceUpdate struct {
	Symbol string `json:"s"`
	Price  string `json:"c"`
}

// StartListener connects to Binance WebSocket and receives price updates
func StartListener(symbol string, onPrice func(PriceUpdate)) {
	u := url.URL{Scheme: "wss", Host: "stream.binance.com:9443", Path: "/ws/" + symbol + "@ticker"}
	logrus.Infof("Conectando ao WebSocket da Binance: %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.WithError(err).Fatal("Erro ao conectar ao WebSocket da Binance")
		return
	}
	defer c.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-interrupt:
			logrus.Info("Listener do Binance finalizado por sinal externo")
			return
		default:
			_, message, err := c.ReadMessage()
			if err != nil {
				logrus.WithError(err).Error("Erro ao ler mensagem do WebSocket")
				time.Sleep(2 * time.Second)
				continue
			}
			var update PriceUpdate
			if err := json.Unmarshal(message, &update); err != nil {
				logrus.WithError(err).Warn("Falha ao parsear mensagem de preço")
				continue
			}
			if update.Symbol != "" && update.Price != "" {
				onPrice(update)
			}
		}
	}
}

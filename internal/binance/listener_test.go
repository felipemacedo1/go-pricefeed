package binance

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPriceUpdateParser(t *testing.T) {
	raw := `{
		"s": "BTCUSDT",
		"o": "50000.00",
		"h": "50500.00",
		"l": "49500.00",
		"c": "50200.00",
		"v": "123.45",
		"E": 1620000000000
	}`

	var update PriceUpdate
	err := json.Unmarshal([]byte(raw), &update)
	assert.NoError(t, err)
	assert.Equal(t, "BTCUSDT", update.Symbol)
	assert.Equal(t, 50000.00, update.OpenPrice)
	assert.Equal(t, 50500.00, update.HighPrice)
	assert.Equal(t, 49500.00, update.LowPrice)
	assert.Equal(t, 50200.00, update.ClosePrice)
	assert.Equal(t, 123.45, update.Volume)
	assert.Equal(t, int64(1620000000000), update.Timestamp)
}

func TestTimeframeDefault(t *testing.T) {
	update := PriceUpdate{
		Symbol:     "BTCUSDT",
		OpenPrice:  50000,
		HighPrice:  50500,
		LowPrice:   49500,
		ClosePrice: 50200,
		Volume:     123.45,
		Timestamp:  time.Now().UnixMilli(),
	}
	if update.Timeframe == "" {
		update.Timeframe = "1m"
	}
	assert.Equal(t, "1m", update.Timeframe)
}

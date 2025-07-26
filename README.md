# PriceFeed Microservice

Microserviço Go para ingestão de preços em tempo real via WebSocket da Binance, com persistência em Redis/PostgreSQL e broadcast para WebSocket interno.

## Estrutura
```
price-feed/
├── cmd/
│   └── server/
│       └── main.go              # ponto de entrada
├── internal/
│   ├── binance/
│   │   └── listener.go          # conexão com WebSocket da Binance
│   ├── redis/
│   │   └── client.go            # conexão e cache de preços
│   ├── postgres/
│   │   └── db.go                # conexão com PostgreSQL (opcional)
│   ├── processor/
│   │   └── dispatcher.go        # lógica que recebe, processa e armazena preços
│   ├── config/
│   │   └── config.go            # parser de env vars ou arquivos
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
```

## Como rodar

1. Instale as dependências:
   ```sh
   go mod tidy
   ```
2. Configure variáveis de ambiente (ver `internal/config/config.go`).
3. Execute o serviço:
   ```sh
   go run cmd/server/main.go
   ```

## Funcionalidades
- Conexão com WebSocket da Binance para ingestão de preços em tempo real
- Cache de preços em Redis
- Persistência opcional em PostgreSQL
- Broadcast para WebSocket interno

## Dependências sugeridas
- [gorilla/websocket](https://github.com/gorilla/websocket)
- [go-redis/redis](https://github.com/go-redis/redis)
- [lib/pq](https://github.com/lib/pq) ou [pgx](https://github.com/jackc/pgx)

## Docker
Veja o `Dockerfile` para build e execução em container.

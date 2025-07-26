FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o pricefeed ./cmd/server
EXPOSE 8080
CMD ["./pricefeed"]

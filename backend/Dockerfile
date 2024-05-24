FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /mazmorras_y_losetas ./cmd/main.go

EXPOSE 8080

CMD ["/mazmorras_y_losetas"]
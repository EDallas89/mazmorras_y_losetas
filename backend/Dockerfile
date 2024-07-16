FROM golang:1.22.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /mazmorras_y_losetas ./src/main.go

EXPOSE 8080

CMD ["/mazmorras_y_losetas"]
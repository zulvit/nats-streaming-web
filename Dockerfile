FROM golang:1.17-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o /nats-publisher ./cmd

CMD ["/nats-publisher"]

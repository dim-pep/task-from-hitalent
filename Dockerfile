FROM golang:1.25.4-alpine

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o qna-api ./cmd

EXPOSE 8080

ENV APP_PORT=8080

CMD ["./qna-api"]
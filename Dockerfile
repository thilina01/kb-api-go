FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o kb-api main.go

EXPOSE 8080

CMD ["./kb-api"]

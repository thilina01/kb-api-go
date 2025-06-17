FROM golang:1.24-alpine

# Install tools
RUN apk add --no-cache git bash curl

WORKDIR /workspace

COPY . .

RUN go mod tidy

EXPOSE 8080
CMD ["go", "run", "main.go"]

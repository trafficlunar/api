# build
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

# copy
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8888

CMD [ "/app/api" ]

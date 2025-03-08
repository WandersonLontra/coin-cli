FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=darwin go build -o coin cmd/main.go

FROM scratch

COPY --from=builder /app/coin /coin

ENTRYPOINT ["/coin"]


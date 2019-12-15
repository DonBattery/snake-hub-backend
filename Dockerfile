FROM golang AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o snake-hub .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/snake-hub .

COPY snake-hub-config.yaml .

CMD ["./snake-hub"]
FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go .
COPY *.html .
RUN go build -o /app/go_pulse_tracker

# RUN CGO_ENABLED=0 GOOS=linux go build -o D:/opt/go-pulse-tracker
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/go_pulse_tracker .

EXPOSE 8080
# ENTRYPOINT ["/app/go_pulse_tracker"]
CMD ["/app/go_pulse_tracker"]


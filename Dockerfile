FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY app ./app
COPY templates ./templates
COPY *.go ./
RUN go build -o zadanie_remitly .

FROM golang:1.22
WORKDIR /app
COPY --from=builder /app/zadanie_remitly .
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./zadanie_remitly"]

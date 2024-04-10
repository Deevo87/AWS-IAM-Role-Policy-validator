# Użyj oficjalnego obrazu Go jako bazowego obrazu
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY app ./app
COPY templates ./templates

# Skopiuj tylko pliki Go
COPY *.go ./
RUN go build -o zadanie_remitly .
RUN CGO_ENABLED=0 GOOS=linux go build -o /json-verifier
EXPOSE 8080
CMD ["/json-verifier"]
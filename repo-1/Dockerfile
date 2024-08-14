FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 80
CMD ["/docker-gs-ping"]

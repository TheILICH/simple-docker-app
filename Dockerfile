FROM golang:1.24-alpine

WORKDIR /app

# For docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Add other things
COPY . .
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
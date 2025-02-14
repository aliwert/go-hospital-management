FROM golang:1.21-alpine

WORKDIR /app

# install required system packages
RUN apk add --no-cache gcc musl-dev

# copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the application
COPY . .

# build the application
RUN go build -o main cmd/main.go

# expose port
EXPOSE 8080

# run the application
CMD ["./main"]

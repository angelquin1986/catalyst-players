# Build stage
FROM golang:1.24

# Set working directory
WORKDIR /app


# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

EXPOSE 8082

# Comando para ejecutar la aplicación
CMD ["./main"]

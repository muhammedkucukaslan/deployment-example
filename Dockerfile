# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM scratch
COPY --from=builder /app/main /main
EXPOSE 3000
CMD ["/main"]
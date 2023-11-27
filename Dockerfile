FROM golang:1.21.3-alpine3.17 AS base
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:3.16 AS binary  

RUN apk --no-cache add ca-certificates
COPY .env .
COPY --from=base /app/main .

EXPOSE 8080
CMD ["./main"]
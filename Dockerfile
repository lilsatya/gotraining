# Start from golang base image
FROM golang:latest as builder

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /main .
COPY --from=builder /.env .

EXPOSE 8080

CMD ["./main"]
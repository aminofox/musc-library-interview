FROM golang:1.22.2 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go install ./vendor/...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -a -o app-bin ./cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app-bin .

EXPOSE 8000

RUN chmod +x /root/app-bin

CMD ["/root/app-bin"] 

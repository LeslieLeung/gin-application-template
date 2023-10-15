FROM golang:1.21 AS builder
WORKDIR $GOPATH/src/packages/app
COPY . .
RUN CGO_ENABLED=0 go build -o /go/myapp main.go

FROM alpine:latest
WORKDIR /app
COPY config /app/config
COPY --from=builder /go/myapp /app/myapp
RUN chmod +x /app/myapp
EXPOSE 8080

CMD ["/app/myapp", "serve"]
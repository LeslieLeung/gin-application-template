FROM alpine:latest
WORKDIR /app
COPY config /app/config
COPY app /app/app
EXPOSE 8080

CMD ["/app/app", "serve"]
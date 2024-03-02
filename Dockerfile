from ubuntu:latest

WORKDIR /app

COPY api /app

CMD ["./api"]


FROM golang:latest
RUN mkdir app
WORKDIR /app
COPY . .

RUN go build  ./cmd/serverqrpc
CMD ["./serverqrpc"]
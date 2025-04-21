FROM golang:1.23.2-alpine

WORKDIR /app
COPY . .

RUN go build -o go-demo

CMD ["./go-demo"]
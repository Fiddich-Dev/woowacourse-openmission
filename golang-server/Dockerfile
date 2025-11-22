FROM golang:1.25.4-alpine

WORKDIR /app

COPY . .
RUN go build -o openmission-server main.go

CMD ["./openmission-server"]

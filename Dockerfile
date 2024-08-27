FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p out/ \
    && go build -o out/reverse-image-generator ./cmd

EXPOSE 8080

CMD ["./out/reverse-image-generator", "start"]

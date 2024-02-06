FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/ilyakaznacheev/cleanenv
RUN go get github.com/nordew/GoChitChat-External/gen/go/user@v0.0.0-20240203224233-0d389f77d860

COPY . .

RUN go build -o gateway ./cmd

CMD ["./gateway"]

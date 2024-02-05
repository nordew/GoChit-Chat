FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get github.com/gin-gonic/gin/binding@v1.9.1
RUN go get github.com/nordew/GoChitChat-External/gen/go/user@v0.0.0-20240203224233-0d389f77d860

COPY . .

RUN go build -o chat ./cmd

CMD ["./chat"]

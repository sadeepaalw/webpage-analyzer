FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go clean
RUN go build


EXPOSE 8080

CMD ["./web-analyzer"]
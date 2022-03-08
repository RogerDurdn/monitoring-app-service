FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main ./

RUN go build -o /main

EXPOSE 9090

CMD [ "/main" ]

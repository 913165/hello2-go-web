FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /hello-go-web

EXPOSE 3000

CMD [ "/hello-go-web" ]
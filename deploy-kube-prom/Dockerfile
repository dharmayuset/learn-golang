FROM golang:1.18-alpine3.16
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

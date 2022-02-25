FROM golang:1.17
WORKDIR /usr/src/app
ADD . .
RUN go build -o /usr/local/bin/go-rest-api main.go
EXPOSE 8000
CMD [ "go-rest-api" ]
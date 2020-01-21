FROM golang:alpine
WORKDIR /app
ADD . /app
RUN apk add git && go get -u github.com/go-sql-driver/mysql && cd /app && go build -o server
CMD ["/app/server"]

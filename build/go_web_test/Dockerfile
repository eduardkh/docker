FROM golang:stretch
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod init example.com/m/v2
RUN go build -o main .
CMD ["/app/main"]

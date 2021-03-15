FROM golang:alpine3.13 AS BUILD
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go mod init example.com/m/v2
RUN go build -o app .


FROM alpine:3.13 AS FINAL
COPY --from=BUILD /src/app /app/app
CMD ["/app/app"]

FROM golang:alpine AS Builder
RUN mkdir /src
COPY . /src
WORKDIR /src
RUN go mod init example.com/m/v2
RUN go build -o app .
FROM alpine:latest AS Final
COPY --from=Builder /src/app /app/app
WORKDIR /app
CMD ["/app/app"]

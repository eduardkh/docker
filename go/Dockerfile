# Build stage
FROM golang:1.21 AS build-env
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp

# Final stage
FROM scratch
COPY --from=build-env /src/myapp /myapp
EXPOSE 8080
ENTRYPOINT ["/myapp"]

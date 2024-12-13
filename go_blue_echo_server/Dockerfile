# Build stage
FROM golang:1.23.3 AS build-env
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp

# Final stage
FROM scratch
COPY --from=build-env /src/myapp /myapp
ENV VERSION="1.0.0"
EXPOSE 8080
ENTRYPOINT ["/myapp"]

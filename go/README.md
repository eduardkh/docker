# sample dotnet app for production

> prerequisite

```bash
# go installed
go version
```

> create, run the app (local development)

```bash
go mod init myapp
go run main.go
```

> build and run the app (Docker)

```bash
# create Dockerfile
cat << EOF > Dockerfile
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
EOF

# use ".dockerignore" to tell docker not to pack unwanted files

# run hadolint (lint the Dockerfile)
hadolint Dockerfile

# build docker container
docker build --no-cache -t hellogoapp .


# save the container image to a file and scan it for vulnerabilities
docker image save hellogoapp:latest -o hellogoapp_latest.tar
syft hellogoapp_latest.tar -o cyclonedx-json=sbom.cyclonedx-json.json
grype sbom:sbom.cyclonedx-json.json

# run docker container
docker run -p 8080:8080 hellogoapp
```

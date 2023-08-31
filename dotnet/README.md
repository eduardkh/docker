# sample dotnet app for production

> prerequisite

```bash
# install SDK (for local development)
sudo apt-get update && sudo apt-get install -y dotnet-sdk-7.0

# install runtime (for local development)
sudo apt-get update && sudo apt-get install -y aspnetcore-runtime-7.0

# check version
dotnet --version
```

> create, build and run the app (local development)

```bash
dotnet new web -o HelloWorldWebApp
cd HelloWorldWebApp

# build local hello world app
dotnet build

# run local hello world app
dotnet run --urls "http://192.168.1.165:8080"
```

> build and run the app (Docker)

```bash
# create Dockerfile
cat << EOF > Dockerfile
# Use SDK image to build the app
FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build-env
WORKDIR /app

# Copy and restore
COPY *.csproj ./
RUN dotnet restore

# Copy everything else and build
COPY . ./
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:7.0
WORKDIR /app
COPY --from=build-env /app/out .
ENTRYPOINT ["dotnet", "HelloWorldWebApp.dll"]
EOF

# use ".dockerignore" to tell docker not to pack unwanted files

# run hadolint (lint the Dockerfile)
hadolint Dockerfile

# build docker container
docker build --no-cache -t helloworldwebapp .

# save the container image to a file and scan it for vulnerabilities
docker image save helloworldwebapp:latest -o helloworldwebapp_latest.tar
syft helloworldwebapp_latest.tar -o cyclonedx-json=sbom.cyclonedx-json.json
grype sbom:sbom.cyclonedx-json.json

# run docker container
docker run -p 8080:80 helloworldwebapp
```

# Blue Server

> commands

```bash
# build docker container
docker build -t myapp .
docker build --no-cache -t myapp .

# save the container image to a file and scan it for vulnerabilities
docker run -p 8080:8080 myapp
docker run -e VERSION=2.0.0 -e CUSTOM_VAR=custom_value -p 8080:8080 myapp
```

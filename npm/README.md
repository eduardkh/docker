# sample node app for production

> prerequisite

```bash
# use ".dockerignore" to tell docker not to pack unwanted files

# run hadolint (lint the Dockerfile)
hadolint Dockerfile

# save the container image to a file and scan it for vulnerabilities
docker image save my-js-app:latest -o my-js-app_latest.tar
syft my-js-app_latest.tar -o cyclonedx-json=sbom.cyclonedx-json.json
grype sbom:sbom.cyclonedx-json.json
```

> build and run the app

```bash
docker build --no-cache -t my-js-app .
docker run -p 3000:3000 my-js-app
```

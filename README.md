# docker build and config example files

> run a stack from yaml file (has to bee in swarm mode)

```bash
docker stack deploy --compose-file docker-compose.yml jenkins
```

> build an image (from Dockerfile)

```bash
docker build -t eduardkh/go_web_test:001 .
```

> run the newly built container (multi-stage)

```bash
docker build --target Builder -t eduardkh/go_web_test_builder:001 .
```

```bash
docker build -t eduardkh/go_web_test_final:001 .
```

> Scan Docker Files with Hadolint

```bash
# Install hadolint
sudo wget -O /bin/hadolint https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && sudo chmod +x /bin/hadolint

# Run hadolint
hadolint Dockerfile
```

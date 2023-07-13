# Syft and Grype

> Generate an SBOM

```bash
syft minio.tar -o cyclonedx-json=sbom.cyclonedx-json.json
```

> Analyze an SBOM

```bash
grype sbom:sbom.cyclonedx-json.json
```

> Generate an SBOM from project a directory

```bash
syft packages dir:/home/vagrant/dj-basic-auth  -o cyclonedx-json=sbom.cyclonedx-json.json
```

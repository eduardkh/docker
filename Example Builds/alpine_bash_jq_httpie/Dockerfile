FROM alpine:3.18.2
LABEL maintainer="Eduard Khiaev"
# get packages here - https://pkgs.alpinelinux.org/packages
RUN apk add --no-cache bash=5.2.15-r5 httpie=3.2.1-r4 jq=1.6-r3 && mkdir /workdir
WORKDIR /workdir
COPY entrypoint.sh /workdir/entrypoint.sh
RUN chmod +x /workdir/entrypoint.sh
ENTRYPOINT ["/workdir/entrypoint.sh"]

# hadolint "Example Builds/alpine_bash_jq_httpie/Dockerfile"
# docker build -f "Example Builds/alpine_bash_jq_httpie/Dockerfile" . -t alpine_bash_jq_httpie:01

# CAUTION: remote builds are not good when you "COPY" stuff :)
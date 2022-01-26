FROM alpine:latest
LABEL maintainer="Eduard Khiaev"
RUN apk add --no-cache \
    bash \
    httpie \
    jq && \
    which bash && \
    which http && \
    which jq
COPY entrypoint.sh usr/local/bin/entrypoint.sh
RUN chmod +x usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
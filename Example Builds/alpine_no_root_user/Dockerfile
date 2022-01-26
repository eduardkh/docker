FROM alpine:latest
LABEL maintainer="Eduard Khiaev"
# RUN apk update && apk upgrade
RUN apk add --no-cache \
    bash && \
    which bash
RUN addgroup -S suka && adduser -S suka -G suka
USER suka
WORKDIR /home/suka
CMD ["bash"]
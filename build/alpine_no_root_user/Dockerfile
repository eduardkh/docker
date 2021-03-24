FROM alpine:latest
LABEL maintainer="Eduard Khiaev"
RUN apk update && apk upgrade
RUN addgroup -S sukagroup && adduser -S suka -G sukagroup
USER suka
WORKDIR /home/suka
CMD [ "sh" ]
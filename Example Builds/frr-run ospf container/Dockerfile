FROM frrouting/frr:latest
LABEL maintainer="Eduard Khiaev"
# run watchfrr zebra ospfd at startup
COPY entrypoint.sh usr/local/bin/entrypoint.sh
RUN chmod +x usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
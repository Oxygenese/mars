FROM golang:1.16 AS builder

COPY . /src
WORKDIR /src
RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y \
        && apt-get install librados-dev \
        && apt-get install librbd-dev

COPY --from=builder /src/bin /app
WORKDIR /app
VOLUME /etc/ceph
VOLUME /var/log/ceph
VOLUME /var/lib/ceph
RUN ["mkdir","-p","/app/cache/naming"]
EXPOSE 8003
EXPOSE 9003
VOLUME /data/chief/conf
CMD sleep 10;./chief -conf /data/chief/conf

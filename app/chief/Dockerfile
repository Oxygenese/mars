FROM debian:stable-slim AS builder

COPY . /src
WORKDIR /src
RUN apt-get update
RUN apt-get install golang
RUN apt-get install librados-dev
RUN apt-get install librbd-dev
RUN make build

FROM debian:stable-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y
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

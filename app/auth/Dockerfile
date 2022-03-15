#FROM golang:1.16 AS builder
#
#COPY . /src
#WORKDIR /src
#
#RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y


COPY ./bin /app
WORKDIR /app
RUN ["mkdir","-p","/app/cache/naming"]
EXPOSE 8001
EXPOSE 9001
VOLUME /data/auth/conf

CMD sleep 10;./auth -conf /data/auth/conf

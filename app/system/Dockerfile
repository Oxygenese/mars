FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y
RUN ["mkdir","-p","/app/cache/naming"]
COPY ./bin/ /app
WORKDIR /app
EXPOSE 8004
EXPOSE 9004
VOLUME /data/system/conf

CMD sleep 10;./system -conf /data/system/conf

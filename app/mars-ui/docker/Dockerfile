FROM registry.cn-shanghai.aliyuncs.com/lwmeng/nginx

COPY ./dist /data

RUN rm /etc/nginx/conf.d/default.conf

ADD mars-ui.conf /etc/nginx/conf.d/

EXPOSE 80

RUN /bin/bash -c 'echo init ok'

FROM alpine:3.2

RUN apk add --update openssh-client && rm -rf /var/cache/apk/*

CMD rm -rf /root/.ssh && mkdir /root/.ssh && cp -R /root/ssh/* /root/.ssh/ && chmod -R 600 /root/.ssh/* && \
ls -l /root/.ssh/ && \
ssh \
-o StrictHostKeyChecking=no \
-N $TUNNEL_HOST \
-L *:$LOCAL_PORT:$REMOTE_HOST:$REMOTE_PORT \
&& while true; do sleep 30; done;

EXPOSE 5432
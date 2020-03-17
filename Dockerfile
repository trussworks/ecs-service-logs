FROM alpine:3
COPY ecs-service-logs /usr/local/bin/ecs-service-logs
ENTRYPOINT [ "ecs-service-logs" ]

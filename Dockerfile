FROM alpine:3
COPY ecs-service-logs /bin/ecs-service-logs
ENTRYPOINT [ "ecs-service-logs" ]

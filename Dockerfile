FROM alpine:20221110
COPY ecs-service-logs /bin/ecs-service-logs
ENTRYPOINT [ "ecs-service-logs" ]

FROM alpine:3
COPY dist/ecs-service-logs_linux_amd64/ecs-service-logs /bin/ecs-service-logs
ENTRYPOINT [ "ecs-service-logs" ]

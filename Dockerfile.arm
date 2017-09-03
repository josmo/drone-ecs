FROM plugins/base:multiarch

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/josmo/drone-ecs.git"
LABEL org.label-schema.name="Drone ECS"
LABEL org.label-schema.vendor="Josmo"
LABEL org.label-schema.schema-version="1.0"

ADD release/linux/arm/drone-ecs /bin/
ENTRYPOINT ["/bin/drone-ecs"]
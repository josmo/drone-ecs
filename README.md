[![Build Status](https://build.andyet.com/api/badges/andyet/drone-ecs/status.svg)](https://build.andyet.com/andyet/drone-ecs)
# drone-ecs

Drone plugin to deploy or update a project on AWS ECS. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build docker
```

### Example

```
docker run --rm                             \
  -e PLUGIN_AWS_SECRET_ACCESS_KEY=<secret>  \
  -e PLUGIN_AWS_ACCESS_KEY_ID=<key>         \
  -e PLUGIN_SERVICE=<service>               \  
  -e PLUGIN_DOCKER_IMAGE=<image>            \
  -v $(pwd):$(pwd)                          \
  -w $(pwd)                                 \
  < your-image >/drone-ecs
```

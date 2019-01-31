[![Build Status](https://cloud.drone.io/api/badges/josmo/drone-ecs/status.svg)](https://cloud.drone.io/josmo/drone-ecs)
[![Go Doc](https://godoc.org/github.com/josmo/drone-ecs?status.svg)](http://godoc.org/github.com/josmo/drone-ecs)
[![Go Report](https://goreportcard.com/badge/github.com/josmo/drone-ecs)](https://goreportcard.com/report/github.com/josmo/drone-ecs)
[![](https://images.microbadger.com/badges/image/peloton/drone-ecs.svg)](https://microbadger.com/images/peloton/drone-ecs "Get your own image badge on microbadger.com")

# drone-ecs


Drone plugin to deploy or update a project on AWS ECS. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `drone cli`:

```
drone exec
```

### Example

```
docker run --rm                          \
  -e PLUGIN_ACCESS_KEY=<key>             \
  -e PLUGIN_SECRET_KEY=<secret>          \
  -e PLUGIN_SERVICE=<service>            \  
  -e PLUGIN_DOCKER_IMAGE=<image>         \
  -v $(pwd):$(pwd)                       \
  -w $(pwd)                              \
  peloton/drone-ecs
```

### Contribution

This repo is setup in a way that if you enable a personal drone server to build your fork it will
 build and publish your image (makes it easier to test PRs and use the image till the contributions get merged)
 
* Build local ```DRONE_REPO_OWNER=josmo DRONE_REPO_NAME=drone-ecs drone exec```
* on your server just make sure you have DOCKER_USERNAME, DOCKER_PASSWORD, and PLUGIN_REPO set as secrets

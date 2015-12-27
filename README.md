# drone-ecs

Drone plugin for deploying to AWS ECS

## Getting Started

Sample .drone.yml file

```
deploy:
  ecs:
    image: plugins/drone-ecs

    region: eu-west-1
    access_key_id: $$ACCESS_KEY_ID
    secret_access_key: $$SECRET_ACCESS_KEY
    task_name: my-ecs-task
    image_name: namespace/repo
    image_tag: latest
    service_name: my-ecs-service
    environment_variables:
      - DATABASE_URI=$$MY_DATABASE_URI
    port_mappings:
      - 80 9000
    memory: 128
```

### Settings

* access_key_id
	* Your AWS access key
	* This MUST be an IAM user with the AmazonEC2ContainerRegistryPowerUser policy attached
* secret_access_key
	* Your AWS secret access key matching the access_key_id
* service_name
	* The service name in your cluster. This MUST be created already
* task_name
	* The name of the task you want to create or update with a new revision
* port_mappings
	* format is "hostPort containerPort"
	* protocol is automatically set to TransportProtocol
* memory
	* amount of memory to assign to the container. Default is 128


## Usage

```sh
./drone-ecs <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-ecs <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
    }
}
EOF
```

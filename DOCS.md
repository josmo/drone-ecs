Use this plugin for deploying a docker container application to AWS EC2 Container Service (ECS).

### Settings

* `access_key` - AWS access key ID, MUST be an IAM user with the AmazonEC2ContainerServiceFullAccess policy attached
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `family` - Family name of the task definition to create or update with a new revision
* `task_role_arn` - ECS task IAM role
* `service` - Name of the service in the cluster, **MUST** be created already in ECS
* `container_name` - Name of the container, defaults to ${family}-container
* `docker_image` - Container image to use, do not include the tag here
* `tag` - Tag of the image to use, defaults to latest
* `cluster` - Name of the cluster. Optional. Default cluster is used if not specified
* `log_driver` - The log driver to use for the container
* `log_options` - The configuration options to send to the log driver
* `port_mappings` - Port mappings from host to container, format is `hostPort containerPort`, protocol is automatically set to TransportProtocol
* `docker_labels` - A key/value map of labels to add to the container
* `environment-variables` - List of Environment Variables to be passed to the container, format is `NAME=VALUE`
* `cpu`, The number of cpu units to reserve for the container
* `memory`, The hard limit (in MiB) of memory to present to the container
* `memory_reservation`, The soft limit (in MiB) of memory to reserve for the container. Defaults to 128
* `deployment_configuration` - Deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks, format is `minimumHealthyPercent maximumPercent`
* `desired_count` - The number of instantiations of the specified task definition to place and keep running on your cluster

## Example

```yaml
deploy:
  ecs:
    image: andyet/drone-ecs
    region: eu-west-1
    access_key: ${ACCESS_KEY_ID}
    secret_key: ${SECRET_ACCESS_KEY}
    family: my-ecs-task
    task_role_arn: arn:aws:iam::012345678901:role/rolename
    service: my-ecs-service
    container-name: my-awesome-container
    docker_image: namespace/repo
    tag: latest
    cluster: my-cluster
    log_driver: awslogs
    log_options:
      - awslogs-group=my-ecs-group
      - awslogs-region=us-east-1
    port_mappings:
      - 80 9000
    environment_variables:
      - DATABASE_URI=${MY_DATABASE_URI}
    docker_labels:
      - traefik.frontend.rule=Host:my.host.gov
      - traefik.backend=pirates
    memory: 256
    memoryReservation: 128
    cpu: 1024
    desired_count: 1
    deployment_configuration: 50 200
```

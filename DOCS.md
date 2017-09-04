Use this plugin for deploying a docker container application to AWS EC2 Container Service (ECS).

### Settings

* `access_key` - AWS access key ID, MUST be an IAM user with the AmazonEC2ContainerServiceFullAccess policy attached
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `service` - Name of the service in the cluster, **MUST** be created already in ECS
* `container_name` - Name of the container, defaults to ${family}-container
* `cluster` - Name of the cluster. Optional. Default cluster is used if not specified
* `family` - Family name of the task definition to create or update with a new revision
* `task_role_arn` - ECS task IAM role
* `docker_image`, Container image to use, do not include the tag here
* `tag` - Tag of the image to use, defaults to latest
* `port_mappings` - Port mappings from host to container, format is `hostPort containerPort`, protocol is automatically set to TransportProtocol
* `cpu`, The number of cpu units to reserve for the container
* `memory`, The hard limit (in MiB) of memory to present to the container
* `memory_reservation`, The soft limit (in MiB) of memory to reserve for the container. Defaults to 128
* `environment_variables` - List of Environment Variables to be passed to the container, format is `NAME=VALUE`
* `deployment_configuration` - Deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks, format is `minimumHealthyPercent maximumPercent`
* `desired_count` - The number of instantiations of the specified task definition to place and keep running on your cluster
* `log_driver` - The log driver to use for the container
* `log_options` - The configuration options to send to the log driver
* `labels` - A key/value map of labels to add to the container
* `secret_environment_variables` - List of Environment Variables to be injected into the container from drone secrets. You can use the name of the secret itself or set a custom name to be used within the container. Syntax is `NAME` (must match the name of one of your secrets) or `CUSTOM_NAME=NAME`


## Example

```yaml
deploy:
  ecs:
    image: peloton/drone-ecs

    region: eu-west-1
    family: my-ecs-task
    docker_image: namespace/repo
    tag: latest
    service: my-ecs-service
    task_role_arn: arn:aws:iam::012345678901:role/rolename
    log_driver: awslogs
    log_options:
      - awslogs-group=my-ecs-group
      - awslogs-region=us-east-1
    environment_variables:
      - DATABASE_URI=$$MY_DATABASE_URI
    secret_environment_variables:
      - MY_SECRET=MY_SANDBOX_SECRET
      - MY_ACCESS_KEY
    labels:
      - traefik.frontend.rule=Host:my.host.gov
      - traefik.backend=pirates
    port_mappings:
      - 80 9000
    memoryReservation: 128
    cpu: 1024
    desired_count: 1
    deployment_configuration: 50 200
    secrets: [AWS_SECRET_KEY, AWS_ACCESS_KEY]
```

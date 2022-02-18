Use this plugin for deploying a docker container application to AWS EC2 Container Service (ECS).

### Settings

* `access_key` - AWS access key ID, MUST be an IAM user with the AmazonEC2ContainerServiceFullAccess policy attached
* `secret_key` - AWS secret access key
* `user_role_arn` - AWS user role. Optional. Switch to different role after initial authentication
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
* `desired_count` - The number of instantiations of the specified task definition to place and keep running on your cluster. Set it to a negative number to not modify current desired_count in the service.
* `log_driver` - The log driver to use for the container
* `log_options` - The configuration options to send to the log driver
* `labels` - A key/value map of labels to add to the container
* `entry_point` - A list of strings to build the container entry point configuration
* `secret_environment_variables` - List of Environment Variables to be injected into the container from drone secrets. You can use the name of the secret itself or set a custom name to be used within the container. Syntax is `NAME` (must match the name of one of your secrets) or `CUSTOM_NAME=NAME`
* `secrets_manager_variables` - List of Environment Variables to be injected into the container from AWS secrets manager, format is `NAME=VALUE`. `VALUE` must match the resource name in AWS secrets manager. I.e. `arn:aws:secretsmanager:region:aws_account_id:secret:password-xxxx`
* `task_cpu` - The number of CPU units used by the task. It can be expressed as an integer using CPU units, for example 1024, or as a string using vCPUs, for example 1 vCPU or 1 vcpu
* `task_memory` - The amount of memory (in MiB) used by the task.It can be expressed as an integer using MiB, for example 1024, or as a string using GB. Required if using Fargate launch type
* `task_execution_role_arn` - The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
* `compatibilities` - Space-delimited list of launch types supported by the task, defaults to EC2 if not specified
* `task_network_mode` - If compatibilities includes FARGATE, this must be set to awsvpc.
* `service_network_assign_public_ip` - Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.
* `service_network_security_groups` - The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.
* `service_network_subnets` - The subnets associated with the task or service. There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.
* `ulimits` - The Ulimit property specifies the ulimit settings to pass to the container. This is an array of strings in the format: `name softLimit hardLimit` where name is one of: core, cpu, data, fsize, locks, memlock, msgqueue, nice, nofile, nproc, rss, rtprio, rttime, sigpending, stack and soft/hard limits are integers.
* `mount_points` - Mount points from host to container, format is `sourceVolume containerPath readOnly` where `sourceVolume`, `containerPath` are strings, `readOnly` is string [`true`, `false`]
* `volumes` - Bind Mount Volumes, format is `name sourcePath` both values are strings. Note with FARGATE launch type, you only provide the name of the volume, not the `sourcePath`
* `efs_volumes` - Define EFS volume, format: `name efs-id root-directory`. Current configuration doesn't support encryption in transit.
* `placement_constraints` - Ecs task definition placement constraints. Specify an array of constraints as a single string. Note that "distinctInstance" type can only be specified during run task or in service. Not inside a task definition. 
* `privileged` - Container will run in privileged mode (applicable only for EC2 launch type)


## Example

```yaml
steps:
  - name: Deploy to ECS
    image: pelotech/drone-ecs
    settings:
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
      placement_constraints: [{"type": "memberOf","expression": "attribute:test == true"}]
      cpu: 1024
      desired_count: 1
      deployment_configuration: 50 200
      ulimits:
        - nofile 2048 4096
      # this mount_point and volumes config will give drone_runner_docker access to docker.sock
      mount_points:
        - dockersock /var/run/docker.sock false
        - myefs /mounted-efs false
      volumes:
        - dockersock /var/run/docker.sock
      efs_volumes:
        - myefs efs-efsid /
      secrets: [AWS_SECRET_KEY, AWS_ACCESS_KEY]
    # declaring the environment is necessary to get secret_environment_variables to work  
    environment:
      MY_SANDBOX_SECRET:
        from_secret: MY_SANDBOX_SECRET
      MY_ACCESS_KEY:
        from_secret: access_key

```

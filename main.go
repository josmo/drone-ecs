package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "AWS ECS Deploy"
	app.Usage = "AWS ECS Deploy"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "AWS access key",
			EnvVar: "PLUGIN_ACCESS_KEY,ECS_ACCESS_KEY,AWS_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "AWS secret key",
			EnvVar: "PLUGIN_SECRET_KEY,ECS_SECRET_KEY,AWS_SECRET_KEY",
		},
		cli.StringFlag{
			Name:   "user-role-arn",
			Usage:  "AWS user role",
			EnvVar: "PLUGIN_USER_ROLE_ARN,ECS_USER_ROLE_ARN,AWS_USER_ROLE_ARN",
		},
		cli.StringFlag{
			Name:   "region",
			Usage:  "aws region",
			EnvVar: "PLUGIN_REGION",
		},
		cli.StringFlag{
			Name:   "family",
			Usage:  "ECS family",
			EnvVar: "PLUGIN_FAMILY",
		},
		cli.StringFlag{
			Name:   "task-role-arn",
			Usage:  "ECS task IAM role",
			EnvVar: "PLUGIN_TASK_ROLE_ARN",
		},
		cli.StringFlag{
			Name:   "service",
			Usage:  "Service to act on",
			EnvVar: "PLUGIN_SERVICE",
		},
		cli.StringFlag{
			Name:   "container-name",
			Usage:  "Container name",
			EnvVar: "PLUGIN_CONTAINER_NAME",
		},
		cli.StringFlag{
			Name:   "docker-image",
			Usage:  "image to use",
			EnvVar: "PLUGIN_DOCKER_IMAGE",
		},
		cli.StringFlag{
			Name:   "tag",
			Usage:  "AWS tag",
			EnvVar: "PLUGIN_TAG",
		},
		cli.StringFlag{
			Name:   "cluster",
			Usage:  "AWS ECS cluster",
			EnvVar: "PLUGIN_CLUSTER",
		},
		cli.StringFlag{
			Name:   "log-driver",
			Usage:  "The log driver to use for the container",
			EnvVar: "PLUGIN_LOG_DRIVER",
		},
		cli.StringSliceFlag{
			Name:   "log-options",
			Usage:  "The configuration options to send to the log driver",
			EnvVar: "PLUGIN_LOG_OPTIONS",
		},
		cli.StringSliceFlag{
			Name:   "port-mappings",
			Usage:  "ECS port maps",
			EnvVar: "PLUGIN_PORT_MAPPINGS",
		},
		cli.StringSliceFlag{
			Name:   "labels",
			Usage:  "A key/value map of labels to add to the container",
			EnvVar: "PLUGIN_LABELS",
		},
		cli.StringSliceFlag{
			Name:   "entry-point",
			Usage:  "A list of values to build the container entry point argument",
			EnvVar: "PLUGIN_ENTRY_POINT",
		},
		cli.StringSliceFlag{
			Name:   "environment-variables",
			Usage:  "ECS environment-variables",
			EnvVar: "PLUGIN_ENVIRONMENT_VARIABLES",
		},
		cli.StringSliceFlag{
			Name:   "secret-environment-variables",
			Usage:  "Secret ECS environment-variables",
			EnvVar: "PLUGIN_SECRET_ENVIRONMENT_VARIABLES",
		},
		cli.StringSliceFlag{
			Name:   "secrets-manager-variables",
			Usage:  "Environment-variables from AWS Secrets manager",
			EnvVar: "PLUGIN_SECRETS_MANAGER_VARIABLES",
		},
		cli.Int64Flag{
			Name:   "cpu",
			Usage:  "The number of cpu units to reserve for the container",
			EnvVar: "PLUGIN_CPU",
		},
		cli.Int64Flag{
			Name:   "memory",
			Usage:  "The hard limit (in MiB) of memory to present to the container",
			EnvVar: "PLUGIN_MEMORY",
		},
		cli.Int64Flag{
			Name:   "memory-reservation",
			Usage:  "The soft limit (in MiB) of memory to reserve for the container. Defaults to 128",
			Value:  128,
			EnvVar: "PLUGIN_MEMORY_RESERVATION",
		},
		cli.StringFlag{
			Name:   "network-mode",
			Usage:  "The Docker networking mode to use for the containers in the task. Defaults to bridge if unspecified",
			EnvVar: "PLUGIN_TASK_NETWORK_MODE",
		},
		cli.StringFlag{
			Name:   "deployment-configuration",
			Usage:  "Deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks",
			EnvVar: "PLUGIN_DEPLOYMENT_CONFIGURATION",
		},
		cli.Int64Flag{
			Name:   "desired-count",
			Usage:  "The number of instantiations of the specified task definition to place and keep running on your cluster",
			EnvVar: "PLUGIN_DESIRED_COUNT",
		},
		cli.BoolTFlag{
			Name:   "yaml-verified",
			Usage:  "Ensure the yaml was signed",
			EnvVar: "DRONE_YAML_VERIFIED",
		},
		cli.StringFlag{
			Name:   "task-cpu",
			Usage:  "The number of CPU units used by the task. It can be expressed as an integer using CPU units, for example 1024, or as a string using vCPUs, for example 1 vCPU or 1 vcpu",
			EnvVar: "PLUGIN_TASK_CPU",
		},
		cli.StringFlag{
			Name:   "task-memory",
			Usage:  "The amount of memory (in MiB) used by the task.It can be expressed as an integer using MiB, for example 1024, or as a string using GB. Required if using Fargate launch type",
			EnvVar: "PLUGIN_TASK_MEMORY",
		},
		cli.StringFlag{
			Name:   "task-execution-role-arn",
			Usage:  "The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.",
			EnvVar: "PLUGIN_TASK_EXECUTION_ROLE_ARN",
		},
		cli.StringFlag{
			Name:   "compatibilities",
			Usage:  "List of launch types supported by the task",
			EnvVar: "PLUGIN_COMPATIBILITIES",
		},
		cli.StringSliceFlag{
			Name:   "healthcheck-command",
			Usage:  "List representing the command that the container runs to determine if it is healthy. Must start with CMD to execute the command arguments directly, or CMD-SHELL to run the command with the container's default shell",
			EnvVar: "PLUGIN_HEALTHCHECK_COMMAND",
		},
		cli.Int64Flag{
			Name:   "healthcheck-interval",
			Usage:  "The time period in seconds between each health check execution. You may specify between 5 and 300 seconds. Defaults to 30 seconds",
			Value:  30,
			EnvVar: "PLUGIN_HEALTHCHECK_INTERVAL",
		},
		cli.Int64Flag{
			Name:   "healthcheck-retries",
			Usage:  "The number of times to retry a failed health check before the container is considered unhealthy. You may specify between 1 and 10 retries. Defaults to 3",
			Value:  3,
			EnvVar: "PLUGIN_HEALTHCHECK_RETRIES",
		},
		cli.Int64Flag{
			Name:   "healthcheck-start-period",
			Usage:  "The grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries. You may specify between 0 and 300 seconds. The startPeriod is disabled by default",
			Value:  0,
			EnvVar: "PLUGIN_HEALTHCHECK_START_PERIOD",
		},
		cli.Int64Flag{
			Name:   "healthcheck-timeout",
			Usage:  "The time period in seconds to wait for a health check to succeed before it is considered a failure. You may specify between 2 and 60 seconds. Defaults to 5 seconds",
			Value:  5,
			EnvVar: "PLUGIN_HEALTHCHECK_TIMEOUT",
		},
		cli.StringFlag{
			Name:   "service-network-assign-public-ip",
			Usage:  "Assign public IP (ENABLED|DISABLED), defaults to DISABLED",
			Value:  "DISABLED",
			EnvVar: "PLUGIN_SERVICE_NETWORK_ASSIGN_PUBLIC_IP",
		},
		cli.StringSliceFlag{
			Name:   "service-network-security-groups",
			Usage:  "The security groups to associate with the service",
			EnvVar: "PLUGIN_SERVICE_NETWORK_SECURITY_GROUPS",
		},
		cli.StringSliceFlag{
			Name:   "service-network-subnets",
			Usage:  "The subnets to associate with the service",
			EnvVar: "PLUGIN_SERVICE_NETWORK_SUBNETS",
		},
		cli.StringSliceFlag{
			Name:   "ulimits",
			Usage:  "ECS ulimits",
			EnvVar: "PLUGIN_ULIMITS",
		},
		cli.StringSliceFlag{
			Name:   "mount-points",
			Usage:  "ECS mount points",
			EnvVar: "PLUGIN_MOUNT_POINTS",
		},
		cli.StringSliceFlag{
			Name:   "volumes",
			Usage:  "ECS volume definitions",
			EnvVar: "PLUGIN_VOLUMES",
		},
		cli.StringSliceFlag{
			Name:   "efs-volumes",
			Usage:  "EFS volume definitions",
			EnvVar: "PLUGIN_EFS_VOLUMES",
		},
		cli.StringFlag{
			Name:   "placement-constraints",
			Usage:  "json array of placement constraints",
			EnvVar: "PLUGIN_PLACEMENT_CONSTRAINTS",
		},
		cli.BoolFlag{
			Name:   "privileged",
			Usage:  "Container will run in privileged mode (applicable only for EC2 launch type)",
			EnvVar: "PLUGIN_PRIVILEGED",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Key:                          c.String("access-key"),
		Secret:                       c.String("secret-key"),
		UserRoleArn:                  c.String("user-role-arn"),
		Region:                       c.String("region"),
		Family:                       c.String("family"),
		TaskRoleArn:                  c.String("task-role-arn"),
		Service:                      c.String("service"),
		ContainerName:                c.String("container-name"),
		DockerImage:                  c.String("docker-image"),
		Tag:                          c.String("tag"),
		Cluster:                      c.String("cluster"),
		LogDriver:                    c.String("log-driver"),
		LogOptions:                   c.StringSlice("log-options"),
		PortMappings:                 c.StringSlice("port-mappings"),
		Environment:                  c.StringSlice("environment-variables"),
		SecretEnvironment:            c.StringSlice("secret-environment-variables"),
		SecretsManagerEnvironment:    c.StringSlice("secrets-manager-variables"),
		EntryPoint:                   c.StringSlice("entry-point"),
		Labels:                       c.StringSlice("labels"),
		CPU:                          c.Int64("cpu"),
		Memory:                       c.Int64("memory"),
		MemoryReservation:            c.Int64("memory-reservation"),
		NetworkMode:                  c.String("network-mode"),
		DeploymentConfiguration:      c.String("deployment-configuration"),
		DesiredCount:                 c.Int64("desired-count"),
		YamlVerified:                 c.BoolT("yaml-verified"),
		TaskCPU:                      c.String("task-cpu"),
		TaskMemory:                   c.String("task-memory"),
		TaskExecutionRoleArn:         c.String("task-execution-role-arn"),
		Compatibilities:              c.String("compatibilities"),
		HealthCheckCommand:           c.StringSlice("healthcheck-command"),
		HealthCheckInterval:          c.Int64("healthcheck-interval"),
		HealthCheckRetries:           c.Int64("healthcheck-retries"),
		HealthCheckStartPeriod:       c.Int64("healthcheck-start-period"),
		HealthCheckTimeout:           c.Int64("healthcheck-timeout"),
		ServiceNetworkAssignPublicIp: c.String("service-network-assign-public-ip"),
		ServiceNetworkSecurityGroups: c.StringSlice("service-network-security-groups"),
		ServiceNetworkSubnets:        c.StringSlice("service-network-subnets"),
		Ulimits:                      c.StringSlice("ulimits"),
		MountPoints:                  c.StringSlice("mount-points"),
		Volumes:                      c.StringSlice("volumes"),
		EfsVolumes:                   c.StringSlice("efs-volumes"),
		PlacementConstraints:         c.String("placement-constraints"),
		Privileged:                   c.Bool("privileged"),
	}
	return plugin.Exec()
}

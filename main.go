package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var version string

func main() {
	app := cli.NewApp()
	app.Name = "rancher publish"
	app.Usage = "rancher publish"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "AWS access key",
			EnvVar: "PLUGIN_ACCESS_KEY, ECS_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "AWS secret key",
			EnvVar: "PLUGIN_SECRET_KEY, ECS_SECRET_KEY",
		},
		cli.StringFlag{
			Name:   "region",
			Usage:  "aws region",
			EnvVar: "PLUGIN_REGION",
		},
		cli.StringFlag{
			Name:   "family",
			Usage:  "ECS family",
			EnvVar: "PLUGIN_FAMILY,",
		},
		cli.StringFlag{
			Name:   "service",
			Usage:  "Service to act on",
			EnvVar: "PLUGIN_SERVICE",
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
		cli.StringSliceFlag{
			Name:	"port-mappings",
			Usage:	"ECS port maps",
			EnvVar: "PLUGIN_PORT_MAPPINGS",
		},
		cli.StringSliceFlag{
			Name:	"environment-variables",
			Usage:	"ECS environment-variables",
			EnvVar: "PLUGIN_ENVIRONMENT_VARIABLES",
		},
		cli.Int64Flag{
			Name:   "memory",
			Usage:  "Amount of memory to the conatiner defaults to 128",
			Value:  128,
			EnvVar: "PLUGIN_MEMORY",
		},
		cli.BoolTFlag{
			Name:   "yaml-verified",
			Usage:  "Ensure the yaml was signed",
			EnvVar: "DRONE_YAML_VERIFIED",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Key:            c.String("access-key"),
		Secret:         c.String("secret-key"),
		Region:         c.String("region"),
		Family:         c.String("family"),
		Service:        c.String("service"),
		DockerImage:    c.String("docker-image"),
		Tag:            c.String("tag"),
		Cluster:        c.String("cluster"),
		PortMappings:   c.StringSlice("port-mappings"),
		Environment:    c.StringSlice("environment-variables"),
		Memory:         c.Int64("memory"),
		YamlVerified:   c.BoolT("yaml-verified"),
	}
	return plugin.Exec()
}

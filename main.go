package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone ECS Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.AccessKeyID) == 0 {
		fmt.Println("Please provide an access key")

		os.Exit(1)
		return
	}

	if len(vargs.SecretAccessKey) == 0 {
		fmt.Println("Please provide a secret key")

		os.Exit(1)
		return
	}

	if len(vargs.Region) == 0 {
		fmt.Println("Please provide a region")

		os.Exit(1)
		return
	}

	if len(vargs.TaskName) == 0 {
		fmt.Println("Please provide a task name")

		os.Exit(1)
		return
	}
	fmt.Println(vargs.TaskName)

	if len(vargs.Image) == 0 {
		fmt.Println("Please provide an image name")

		os.Exit(1)
		return
	}

	if len(vargs.Service) == 0 {
		fmt.Println("Please provide a service name")

		os.Exit(1)
		return
	}

	if len(vargs.Memory) == 0 {
		fmt.Println("Memory not specified. Defaulting to 128")
		vargs.Memory = "128"
	}

	svc := ecs.New(
		session.New(&aws.Config{
			Region:      aws.String(vargs.Region),
			Credentials: credentials.NewStaticCredentials(vargs.AccessKeyID, vargs.SecretAccessKey, ""),
		}))

	Image := vargs.Image + ":" + vargs.Tag
	memory, memoryErr := strconv.ParseInt(vargs.Memory, 10, 64)
	if memoryErr != nil {
		fmt.Println(memoryErr.Error())
		os.Exit(1)
		return
	}

	definition := ecs.ContainerDefinition{ // Required
		Command: []*string{},

		DnsSearchDomains:      []*string{},
		DnsServers:            []*string{},
		DockerLabels:          map[string]*string{},
		DockerSecurityOptions: []*string{},
		EntryPoint:            []*string{},
		Environment:           []*ecs.KeyValuePair{},
		Essential:             aws.Bool(true),
		ExtraHosts:            []*ecs.HostEntry{},

		Image:        aws.String(Image),
		Links:        []*string{},
		Memory:       aws.Int64(memory),
		MountPoints:  []*ecs.MountPoint{},
		Name:         aws.String(vargs.TaskName),
		PortMappings: []*ecs.PortMapping{},

		Ulimits: []*ecs.Ulimit{},
		//User: aws.String("String"),
		VolumesFrom: []*ecs.VolumeFrom{},
		//WorkingDirectory: aws.String("String"),
	}

	// Port mappings
	for _, portMapping := range vargs.PortMappings.Slice() {
		parts := strings.Split(portMapping, " ")
		hostPort, hostPortErr := strconv.ParseInt(parts[0], 10, 64)
		if hostPortErr != nil {
			fmt.Println(hostPortErr.Error())
			os.Exit(1)
			return
		}
		containerPort, containerPortError := strconv.ParseInt(parts[1], 10, 64)
		if containerPortError != nil {
			fmt.Println(containerPortError.Error())
			os.Exit(1)
			return
		}

		pair := ecs.PortMapping{
			ContainerPort: aws.Int64(containerPort),
			HostPort:      aws.Int64(hostPort),
			Protocol:      aws.String("TransportProtocol"),
		}

		definition.PortMappings = append(definition.PortMappings, &pair)
	}

	// Environment variables
	for _, envVar := range vargs.EnvironmentVariables.Slice() {
		parts := strings.Split(envVar, "=")
		pair := ecs.KeyValuePair{
			Name:  aws.String(strings.Trim(parts[0], " ")),
			Value: aws.String(strings.Trim(parts[1], " ")),
		}
		definition.Environment = append(definition.Environment, &pair)
	}
	params := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: []*ecs.ContainerDefinition{ // Required
			&definition,
			// More values...
		},
		Family:  aws.String(vargs.TaskName), // Required
		Volumes: []*ecs.Volume{},
	}
	resp, err := svc.RegisterTaskDefinition(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())

		os.Exit(1)
		return
	}

	val := *(resp.TaskDefinition.TaskDefinitionArn)
	sparams := &ecs.UpdateServiceInput{
		Service:        aws.String(vargs.Service), // Required
		TaskDefinition: aws.String(val),
	}
	sresp, serr := svc.UpdateService(sparams)

	if serr != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(serr.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(sresp)

	fmt.Println(resp)
}

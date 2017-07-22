package main

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ecs"
)

type Plugin struct {
    Key                     string
    Secret                  string
    Region                  string
    Family                  string
    TaskRoleArn             string
    Service                 string
    ContainerName           string
    DockerImage             string
    Tag                     string
    Cluster                 string
    DeploymentConfiguration string
    LogDriver               string
    LogOptions              []string
    PortMappings            []string
    DockerLabels            []string
    Environment             []string
    DesiredCount            int64
    CPU                     int64
    Memory                  int64
    MemoryReservation       int64
    YamlVerified            bool
}

func (p *Plugin) Exec() error {
    fmt.Println("Drone AWS ECS Plugin built")
    awsConfig := aws.Config{}

    if len(p.Key) != 0 && len(p.Secret) != 0 {
        awsConfig.Credentials = credentials.NewStaticCredentials(p.Key, p.Secret, "")
    }
    awsConfig.Region = aws.String(p.Region)
    svc := ecs.New(session.New(&awsConfig))

    Image := p.DockerImage + ":" + p.Tag
    if len(p.ContainerName) == 0 {
        p.ContainerName = p.Family + "-container"
    }

    definition := ecs.ContainerDefinition{
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
        Memory:       aws.Int64(p.Memory),
        MountPoints:  []*ecs.MountPoint{},
        Name:         aws.String(p.ContainerName),
        PortMappings: []*ecs.PortMapping{},

        Ulimits: []*ecs.Ulimit{},
        //User: aws.String("String"),
        VolumesFrom: []*ecs.VolumeFrom{},
        //WorkingDirectory: aws.String("String"),
    }

    if p.CPU != 0 {
        definition.Cpu = aws.Int64(p.CPU)
    }

    if p.Memory == 0 && p.MemoryReservation == 0 {
        definition.MemoryReservation = aws.Int64(128)
    } else {
        if p.Memory != 0 {
            definition.Memory = aws.Int64(p.Memory)
        }
        if p.MemoryReservation != 0 {
            definition.MemoryReservation = aws.Int64(p.MemoryReservation)
        }
    }

    // Port mappings
    for _, portMapping := range p.PortMappings {
        cleanedPortMapping := strings.Trim(portMapping, " ")
        parts := strings.SplitN(cleanedPortMapping, " ", 2)
        hostPort, hostPortErr := strconv.ParseInt(parts[0], 10, 64)
        if hostPortErr != nil {
            fmt.Println(hostPortErr.Error())
            return hostPortErr
        }
        containerPort, containerPortError := strconv.ParseInt(parts[1], 10, 64)
        if containerPortError != nil {
            fmt.Println(containerPortError.Error())
            return containerPortError
        }

        pair := ecs.PortMapping{
            ContainerPort: aws.Int64(containerPort),
            HostPort:      aws.Int64(hostPort),
            Protocol:      aws.String("TransportProtocol"),
        }

        definition.PortMappings = append(definition.PortMappings, &pair)
    }

    // Environment variables
    for _, envVar := range p.Environment {
        parts := strings.SplitN(envVar, "=", 2)
        pair := ecs.KeyValuePair{
            Name:  aws.String(strings.Trim(parts[0], " ")),
            Value: aws.String(strings.Trim(parts[1], " ")),
        }
        definition.Environment = append(definition.Environment, &pair)
    }

    // DockerLabels
    for _, label := range p.DockerLabels {
        parts := strings.SplitN(label, "=", 2)
        definition.DockerLabels[strings.Trim(parts[0], " ")] = aws.String(strings.Trim(parts[1], " "))
    }

    // LogOptions
    if len(p.LogDriver) > 0 {
        definition.LogConfiguration = new(ecs.LogConfiguration)
        definition.LogConfiguration.LogDriver = &p.LogDriver
        if len(p.LogOptions) > 0 {
            definition.LogConfiguration.Options = make(map[string]*string)
            for _, logOption := range p.LogOptions {
                parts := strings.SplitN(logOption, "=", 2)
                logOptionKey := strings.Trim(parts[0], " ")
                logOptionValue := aws.String(strings.Trim(parts[1], " "))
                definition.LogConfiguration.Options[logOptionKey] = logOptionValue
            }
        }
    }

    params := &ecs.RegisterTaskDefinitionInput{
        ContainerDefinitions: []*ecs.ContainerDefinition{
            &definition,
        },
        Family:      aws.String(p.Family),
        Volumes:     []*ecs.Volume{},
        TaskRoleArn: aws.String(p.TaskRoleArn),
    }
    resp, err := svc.RegisterTaskDefinition(params)

    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    val := *(resp.TaskDefinition.TaskDefinitionArn)
    sparams := &ecs.UpdateServiceInput{
        Cluster:        aws.String(p.Cluster),
        Service:        aws.String(p.Service),
        TaskDefinition: aws.String(val),
    }

    if p.DesiredCount != 0 {
        sparams.DesiredCount = aws.Int64(p.DesiredCount)
    }

    // DeploymentConfiguration
    cleanedDeploymentConfiguration := strings.Trim(p.DeploymentConfiguration, " ")
    parts := strings.SplitN(cleanedDeploymentConfiguration, " ", 2)
    if len(parts) > 1 {
        minimumHealthyPercent, minimumHealthyPercentError := strconv.ParseInt(parts[0], 10, 64)
        if minimumHealthyPercentError != nil {
            fmt.Println(minimumHealthyPercentError.Error())
            return minimumHealthyPercentError
        }
        maximumPercent, maximumPercentErr := strconv.ParseInt(parts[1], 10, 64)
        if maximumPercentErr != nil {
            fmt.Println(maximumPercentErr.Error())
            return maximumPercentErr
        }

        sparams.DeploymentConfiguration = &ecs.DeploymentConfiguration{
            MaximumPercent:        aws.Int64(maximumPercent),
            MinimumHealthyPercent: aws.Int64(minimumHealthyPercent),
        }
    }

    sresp, serr := svc.UpdateService(sparams)

    if serr != nil {
        fmt.Println(serr.Error())
        return serr
    }

    fmt.Println(sresp)

    fmt.Println(resp)
    return nil

}

package main

import (
	"github.com/drone/drone-go/drone"

)
type Params struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Region          string `json:"region"`
	TaskName          string `json:"task_name"`
	Image          string `json:"image_name"`
	Tag          string `json:"image_tag"`
	Service          string `json:"service_name"`
	Memory          string `json:"memory"`
	EnvironmentVariables          drone.StringSlice `json:"environment_variables"`
	PortMappings          drone.StringSlice `json:"port_mappings"`
}
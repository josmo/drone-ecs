package main

import (
	"github.com/drone/drone-go/drone"
)

type Params struct {
	AccessKey    string            `json:"access_key"`
	SecretKey    string            `json:"secret_key"`
	Region       string            `json:"region"`
	Family       string            `json:"family"`
	Image        string            `json:"image_name"`
	Tag          string            `json:"image_tag"`
	Service      string            `json:"service"`
	Memory       int64             `json:"memory"`
	Environment  drone.StringSlice `json:"environment_variables"`
	PortMappings drone.StringSlice `json:"port_mappings"`
}

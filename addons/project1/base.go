package project1

import (
	"github.com/arutselvan15/pulumi-lab/addons/k8s"
)

type P1ResourceBuilder struct {
	k8s.ResourceBuilder
}

func NewResourceBuilder(ro *k8s.ResourceOptions) *P1ResourceBuilder {
	ro.BuilderName = "project-1"
	return &P1ResourceBuilder{k8s.ResourceBuilder{ResOpts: ro}}
}

func (r *P1ResourceBuilder) AdjustDeploymentOptions(do *k8s.DeploymentOptions) *k8s.DeploymentOptions {
	do.ImagePullPolicy = "Always"
	return do
}

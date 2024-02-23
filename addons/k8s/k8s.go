package k8s

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

//type ResourceBuilder interface {
//	Deployment(*pulumi.Context, *DeploymentOptions) (*appsv1.Deployment, error)
//}

type ResourceOptions struct {
	Metadata
	BuilderName string
	Provider    pulumi.ResourceOrInvokeOption
}

type ResourceBuilder struct {
	ResOpts *ResourceOptions
}

func NewResourceBuilder(ro *ResourceOptions) *ResourceBuilder {
	ro.BuilderName = "k8s"
	return &ResourceBuilder{ResOpts: ro}
}

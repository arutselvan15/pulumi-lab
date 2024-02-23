package k8s

import (
	"fmt"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	defaultReplicas        = 1
	defaultImagePullPolicy = "IfNotPresent"
)

type DeploymentOptions struct {
	Metadata
	Replicas        *int
	Image           string
	ImagePullPolicy string
}

func (r *ResourceBuilder) Deployment(ctx *pulumi.Context, do *DeploymentOptions) (*appsv1.Deployment, error) {
	fmt.Printf("##### %s deployment: %s\n", r.ResOpts.BuilderName, do.Name)
	do = r.applyDeploymentDefaultOptions(do)
	do = r.AdjustDeploymentOptions(do)
	return appsv1.NewDeployment(ctx, do.Name, r.deploymentArgs(do), r.ResOpts.Provider)
}

func (r *ResourceBuilder) AdjustDeploymentOptions(do *DeploymentOptions) *DeploymentOptions {
	return do
}

func (r *ResourceBuilder) AdjustDeploymentObjectMetaArgs(dom *metav1.ObjectMetaArgs) *metav1.ObjectMetaArgs {
	return dom
}

func (r *ResourceBuilder) AdjustDeploymentSpecArgs(ds *appsv1.DeploymentSpecArgs) *appsv1.DeploymentSpecArgs {
	return ds
}

func (r *ResourceBuilder) applyDeploymentDefaultOptions(do *DeploymentOptions) *DeploymentOptions {
	if do.Replicas == nil {
		do.Replicas = &defaultReplicas
	}
	if do.ImagePullPolicy == "" {
		do.ImagePullPolicy = defaultImagePullPolicy
	}
	if do.Labels == nil {
		do.Labels = r.ResOpts.Labels
	}
	if do.Annotations == nil {
		do.Annotations = r.ResOpts.Annotations
	}
	if do.Namespace == "" {
		do.Namespace = r.ResOpts.Namespace
	}
	return do
}

func (r *ResourceBuilder) deploymentObjectMetaArgs(do *DeploymentOptions) *metav1.ObjectMetaArgs {
	metadata := &metav1.ObjectMetaArgs{
		Name:        pulumi.String(do.Name),
		Labels:      pulumi.ToStringMap(do.Labels),
		Annotations: pulumi.ToStringMap(do.Annotations),
	}

	if do.Namespace != "" {
		metadata.Namespace = pulumi.String(do.Namespace)
	}

	fmt.Printf("-->%s\n", do.ImagePullPolicy)
	return metadata
}

func (r *ResourceBuilder) deploymentSpecArgs(do *DeploymentOptions) *appsv1.DeploymentSpecArgs {
	spec := &appsv1.DeploymentSpecArgs{
		Selector: &metav1.LabelSelectorArgs{
			MatchLabels: pulumi.ToStringMap(do.Labels),
		},
		Replicas: pulumi.Int(*do.Replicas),
		Template: &corev1.PodTemplateSpecArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Labels: pulumi.ToStringMap(do.Labels),
			},
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{
					corev1.ContainerArgs{
						Name:            pulumi.String(do.Name),
						Image:           pulumi.String(do.Image),
						ImagePullPolicy: pulumi.String(do.ImagePullPolicy),
					},
				},
			},
		},
	}

	return spec
}

func (r *ResourceBuilder) deploymentArgs(do *DeploymentOptions) *appsv1.DeploymentArgs {
	return &appsv1.DeploymentArgs{
		Metadata: r.AdjustDeploymentObjectMetaArgs(r.deploymentObjectMetaArgs(do)),
		Spec:     r.AdjustDeploymentSpecArgs(r.deploymentSpecArgs(do)),
	}
}

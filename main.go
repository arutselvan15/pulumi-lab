package main

import (
	"github.com/arutselvan15/pulumi-lab/addons/k8s"
	"github.com/arutselvan15/pulumi-lab/addons/project1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumiK8s()
}

func pulumiK8s() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return pulumiRun(ctx)
	})
}

func pulumiRun(ctx *pulumi.Context) error {
	renderProvider, err := kubernetes.NewProvider(ctx, "k8s-yaml-renderer", &kubernetes.ProviderArgs{
		RenderYamlToDirectory: pulumi.String("yaml"),
	})
	if err != nil {
		return err
	}

	rs := k8s.NewResourceBuilder(&k8s.ResourceOptions{Provider: pulumi.Provider(renderProvider)})
	_, err = rs.Deployment(
		ctx,
		&k8s.DeploymentOptions{
			Metadata: k8s.Metadata{
				Name:        "nginx",
				Namespace:   "default",
				Labels:      map[string]string{"app": "nginx"},
				Annotations: map[string]string{"app": "nginx"},
			},
		})
	if err != nil {
		return err
	}

	prs := project1.NewResourceBuilder(&k8s.ResourceOptions{Provider: pulumi.Provider(renderProvider)})
	_, err = prs.Deployment(
		ctx,
		&k8s.DeploymentOptions{
			Metadata: k8s.Metadata{
				Name:        "nginx-project-1",
				Namespace:   "default",
				Labels:      map[string]string{"app": "nginx-project-1"},
				Annotations: map[string]string{"app": "nginx-project-1"},
			},
		})
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"github.com/construkt-dev/construkt"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type Grafana struct {
	construkt.Meta
	Spec GrafanaSpec `json:"spec"`
}

type GrafanaSpec struct {
	GrafanaVersion string
}

func main() {
	ctx := construkt.NewContext()
	grafanaResource := construkt.NewResource("examples.construkt.dev", "Grafana", "v1", &Grafana{})
	go grafanaResource.RunReconciler(ctx, ReconcileFunc)
	ctx.Wait()
}

func ReconcileFunc(obj *Grafana) ([]construkt.HasMeta, error) {
	depl := apps.Deployment{
		Spec: apps.DeploymentSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "grafana",
							Image: fmt.Sprintf("grafana/grafana:%s", obj.Spec.GrafanaVersion),
						},
					},
				},
			},
		},
	}
	return []construkt.HasMeta{&depl}, nil
}

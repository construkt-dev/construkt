package main

import (
	"github.com/construkt-dev/construkt"
)

type Grafana = GrafanaV1

const currentVersion = "v1"

type GrafanaV1 struct {
	GrafanaVersion string
}

func main() {
	// The construkt Context holds the kubernetes config and other shared state
	ctx := construkt.NewContext()

	// Create a new resource definition for the Grafana CRD - this does not do anything in kubernetes yet
	grafanaResource := construkt.NewResource("examples.construkt.dev", "Grafana", currentVersion, &Grafana{})

	// Start the reconciler loop, calling the Reconcile function for each object
	go grafanaResource.RunManagedReconciler(ctx, ReconcileFunc)

	// Wait for an interrupt signal or fatal errors and then try to shut down gracefully
	ctx.Wait()
}

func ReconcileFunc(obj *construkt.Object[*Grafana]) ([]construkt.Object[any], error) {
	return nil, nil
}

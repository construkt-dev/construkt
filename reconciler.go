package construkt

import (
	gocontext "context"
	"fmt"
	"k8s.io/apimachinery/pkg/watch"
)

type Reconciler[T HasMeta] interface {
	Reconcile(obj T) error
}

type ReconcileFunc[T HasMeta] func(obj T) ([]HasMeta, error)

func RunManagedReconciler[T HasMeta](ctx *Context, cr ClientProvider[T], fn ReconcileFunc[T]) {
	RunReconciler[T](ctx, cr, &managedReconciler[T]{
		fn: fn,
	})
}

type managedReconciler[T HasMeta] struct {
	fn ReconcileFunc[T]
}

func (m *managedReconciler[T]) Reconcile(obj T) error {
	results, err := m.fn(obj)
	for _, result := range results {
		fmt.Printf("Creating %s/%s of type %s\n", result.GetNamespace(), result.GetName(), result.GroupVersionKind())
	}
	return err
}

type ClientProvider[T HasMeta] interface {
	ResourceId() string
	Client(ctx *Context) (*Client[T], error)
}

func RunReconciler[T HasMeta](ctx *Context, cr ClientProvider[T], reconciler Reconciler[T]) {
	id := fmt.Sprintf("reconciler-%s", cr.ResourceId())
	done := ctx.RegisterComponent(id)
	defer done()

	cli, err := cr.Client(ctx)
	if err != nil {
		panic(err)
	}

	w, err := cli.Watch(gocontext.Background())
	if err != nil {
		ctx.ShutdownErr(err)
		return
	}
	defer w.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopping reconciler for %s\n", cr.ResourceId())
			return
		case event := <-w.ResultChan():
			if event.Type == watch.Error {
				panic(event.Object)
			}
			err := reconciler.Reconcile(event.Object.Inner)
			if err != nil {
				panic(err)
			}
		}
	}
}

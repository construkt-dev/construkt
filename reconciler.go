package construkt

import (
	gocontext "context"
	"fmt"
	"k8s.io/apimachinery/pkg/watch"
)

type Reconciler[T any] interface {
	Reconcile(obj *Object[T]) error
}

type ReconcileFunc[T any] func(obj *Object[T]) ([]Object[any], error)

func RunManagedReconciler[T any](ctx *Context, cr ClientProvider[T], fn ReconcileFunc[T]) {
	RunReconciler[T](ctx, cr, &managedReconciler[T]{
		fn: fn,
	})
}

type managedReconciler[T any] struct {
	fn ReconcileFunc[T]
}

func (m *managedReconciler[T]) Reconcile(obj *Object[T]) error {
	results, err := m.fn(obj)
	for _, result := range results {
		fmt.Printf("Creating %s/%s of type %s\n", result.GetNamespace(), result.GetName(), result.GetObjectKind())
	}
	return err
}

type ClientProvider[T any] interface {
	ResourceId() string
	Client(ctx *Context) (*Client[T], error)
}

func RunReconciler[T any](ctx *Context, cr ClientProvider[T], reconciler Reconciler[T]) {
	done := ctx.RegisterComponent(fmt.Sprintf("reconciler-%s", cr.ResourceId()))
	defer done()

	cli, err := cr.Client(ctx)
	if err != nil {
		panic(err)
	}

	w, err := cli.Watch(gocontext.Background())
	if err != nil {
		panic(err)
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
			err := reconciler.Reconcile(event.Object)
			if err != nil {
				panic(err)
			}
		}
	}
}

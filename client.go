package construkt

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

type Client[T HasMeta] struct {
	client client.WithWatch
	scheme *runtime.Scheme
}

type RestConfigProvider interface {
	Config() *rest.Config
}

func NewClient[T HasMeta](ctx *Context, res ResourceMeta) (*Client[T], error) {
	runtimeScheme := runtime.NewScheme()
	builder := scheme.Builder{GroupVersion: schema.GroupVersion{
		Group:   res.gvk.Group,
		Version: res.gvk.Version,
	}}

	builder.SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypeWithName(res.gvk, &object[T]{})
		s.AddKnownTypeWithName(schema.GroupVersionKind{
			Group:   res.gvk.Group,
			Version: res.gvk.Version,
			Kind:    res.names.listKind,
		}, &ObjectList[T]{})
		v1.AddToGroupVersion(s, builder.GroupVersion)
		return nil
	})
	err := builder.AddToScheme(runtimeScheme)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewWithWatch(ctx.Config(), client.Options{
		Scheme: runtimeScheme,
	})
	if err != nil {
		return nil, err
	}

	return &Client[T]{
		client: cli,
		scheme: runtimeScheme,
	}, nil
}

func (c *Client[T]) Get(ctx context.Context, key client.ObjectKey) (object[T], error) {
	t := object[T]{}

	err := c.client.Get(ctx, key, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (c *Client[T]) Watch(ctx context.Context) (Watch[T], error) {
	t := ObjectList[T]{}

	w, err := c.client.Watch(ctx, &t)
	if err != nil {
		return nil, err
	}

	return newWatch[T](w), nil
}

func (c *Client[T]) List(ctx context.Context) (ObjectList[T], error) {
	t := ObjectList[T]{}

	err := c.client.List(ctx, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (c *Client[T]) Create(ctx context.Context, obj *object[T]) error {
	return c.client.Create(ctx, obj)
}

func (c *Client[T]) Update(ctx context.Context, obj *object[T]) error {
	return c.client.Update(ctx, obj)
}

package construkt

import (
	"fmt"
	"github.com/construkt-dev/construkt/internal/jsonschema"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

type Resource[T HasMeta] struct {
	ResourceMeta
	versions map[string]ResourceVersion[T]
}

type ResourceMeta struct {
	gvk        schema.GroupVersionKind
	names      ResourceNames
	namespaced bool
}

type ResourceNames struct {
	plural     string
	singular   string
	shortNames []string
	listKind   string
}

type ResourceVersion[T HasMeta] struct {
	ResourceVersionMeta
}

type ResourceVersionMeta struct {
	served  bool
	stored  bool
	version string
	schema  v1.JSONSchemaProps
}

func NewResource[T HasMeta](api, kind, version string, resource T, opts ...CustomResourceOpt) *Resource[T] {
	singular := strings.ToLower(kind)

	meta := ResourceMeta{
		gvk: schema.GroupVersionKind{
			Group:   api,
			Version: version,
			Kind:    kind,
		},
		names: ResourceNames{
			plural:   fmt.Sprintf("%ss", singular),
			singular: singular,
			listKind: fmt.Sprintf("%sList", kind),
		},
	}

	versionMeta := ResourceVersionMeta{
		version: version,
		served:  true,
		stored:  true,
		schema:  jsonschema.ToJsonSchema(resource),
	}

	for _, opt := range opts {
		meta, versionMeta = opt.Apply(meta, versionMeta)
	}

	return &Resource[T]{
		ResourceMeta: meta,
		versions: map[string]ResourceVersion[T]{
			version: {
				ResourceVersionMeta: versionMeta,
			},
		},
	}
}

type Convertable[A any] interface {
	From(A)
	Into() A
}

func (r *Resource[T]) Version(version string, resource Convertable[T], opts ...CustomResourceVersionOpt) *Resource[T] {
	versionMeta := ResourceVersionMeta{
		version: version,
		served:  true,
		stored:  true,
		schema:  jsonschema.ToJsonSchema(resource),
	}

	for _, opt := range opts {
		versionMeta = opt.Apply(versionMeta)
	}

	r.versions[version] = ResourceVersion[T]{
		ResourceVersionMeta: versionMeta,
	}

	return r
}

func (r *Resource[T]) Client(ctx *Context) (*Client[T], error) {
	return NewClient[T](ctx, r.ResourceMeta)
}

func (r *Resource[T]) ResourceId() string {
	return fmt.Sprintf("%s.%s/%s", r.names.plural, r.gvk.Group, r.gvk.Version)
}

func (r *Resource[T]) RunReconciler(ctx *Context, p ReconcileFunc[T]) {
	RunManagedReconciler[T](ctx, r, p)
}

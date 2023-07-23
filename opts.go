package construkt

import v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type CustomResourceOpt interface {
	Apply(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta)
}

type CustomResourceOptFunc func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta)

func (f CustomResourceOptFunc) Apply(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
	return f(meta, version)
}

func Namespaced() CustomResourceOptFunc {
	return func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
		meta.namespaced = true
		return meta, version
	}
}

func ClusterScoped() CustomResourceOptFunc {
	return func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
		meta.namespaced = false
		return meta, version
	}
}

func Plural(plural string) CustomResourceOptFunc {
	return func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
		meta.names.plural = plural
		return meta, version
	}
}

func Singular(singular string) CustomResourceOptFunc {
	return func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
		meta.names.singular = singular
		return meta, version
	}
}

func ShortNames(shortNames ...string) CustomResourceOptFunc {
	return func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
		meta.names.shortNames = shortNames
		return meta, version
	}
}

func ListKind(listKind string) CustomResourceOptFunc {
	return func(meta ResourceMeta, version ResourceVersionMeta) (ResourceMeta, ResourceVersionMeta) {
		meta.names.listKind = listKind
		return meta, version
	}
}

type CustomResourceVersionOpt interface {
	Apply(version ResourceVersionMeta) ResourceVersionMeta
}

type CustomResourceVersionOptFunc func(version ResourceVersionMeta) ResourceVersionMeta

func (f CustomResourceVersionOptFunc) Apply(r ResourceVersionMeta) ResourceVersionMeta {
	return f(r)
}

func Served() CustomResourceVersionOptFunc {
	return func(r ResourceVersionMeta) ResourceVersionMeta {
		r.served = true
		return r
	}
}

func NotServed() CustomResourceVersionOptFunc {
	return func(r ResourceVersionMeta) ResourceVersionMeta {
		r.served = false
		return r
	}
}

func ModifySchema(f func(schema v1.JSONSchemaProps) v1.JSONSchemaProps) CustomResourceVersionOptFunc {
	return func(r ResourceVersionMeta) ResourceVersionMeta {
		r.schema = f(r.schema)
		return r
	}
}

package construkt

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8sObject = client.Object
type TypeMeta = metav1.TypeMeta
type ObjectMeta = metav1.ObjectMeta

type Object[T any] struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	Spec       T `json:"spec"`
}

func NewObject[T any](meta ObjectMeta, spec T) *Object[T] {
	return &Object[T]{
		TypeMeta:   TypeMeta{},
		ObjectMeta: meta,
		Spec:       spec,
	}
}

func (r *Object[T]) DeepCopyObject() runtime.Object {
	panic("implement me")
}

type ObjectList[T any] struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object[T] `json:"items"`
}

func (r *ObjectList[T]) DeepCopyObject() runtime.Object {
	panic("implement me")
}

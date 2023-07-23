package construkt

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8sObject = client.Object
type TypeMeta = metav1.TypeMeta
type ObjectMeta = metav1.ObjectMeta

type Meta struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
}

type HasMeta interface {
	schema.ObjectKind
	metav1.Object
}

type object[T HasMeta] struct {
	Inner T
}

func (o *object[T]) GetNamespace() string {
	return o.Inner.GetNamespace()
}

func (o *object[T]) SetNamespace(namespace string) {
	o.Inner.SetNamespace(namespace)
}

func (o *object[T]) GetName() string {
	return o.Inner.GetName()
}

func (o *object[T]) SetName(name string) {
	o.Inner.SetName(name)
}

func (o *object[T]) GetGenerateName() string {
	return o.Inner.GetGenerateName()
}

func (o *object[T]) SetGenerateName(name string) {
	o.Inner.SetGenerateName(name)
}

func (o *object[T]) GetUID() types.UID {
	return o.Inner.GetUID()
}

func (o *object[T]) SetUID(uid types.UID) {
	o.Inner.SetUID(uid)
}

func (o *object[T]) GetResourceVersion() string {
	return o.Inner.GetResourceVersion()
}

func (o *object[T]) SetResourceVersion(version string) {
	o.Inner.SetResourceVersion(version)
}

func (o *object[T]) GetGeneration() int64 {
	return o.Inner.GetGeneration()
}

func (o *object[T]) SetGeneration(generation int64) {
	o.Inner.SetGeneration(generation)
}

func (o *object[T]) GetSelfLink() string {
	return o.Inner.GetSelfLink()
}

func (o *object[T]) SetSelfLink(selfLink string) {
	o.Inner.SetSelfLink(selfLink)
}

func (o *object[T]) GetCreationTimestamp() metav1.Time {
	return o.Inner.GetCreationTimestamp()
}

func (o *object[T]) SetCreationTimestamp(timestamp metav1.Time) {
	o.Inner.SetCreationTimestamp(timestamp)
}

func (o *object[T]) GetDeletionTimestamp() *metav1.Time {
	return o.Inner.GetDeletionTimestamp()
}

func (o *object[T]) SetDeletionTimestamp(timestamp *metav1.Time) {
	o.Inner.SetDeletionTimestamp(timestamp)
}

func (o *object[T]) GetDeletionGracePeriodSeconds() *int64 {
	return o.Inner.GetDeletionGracePeriodSeconds()
}

func (o *object[T]) SetDeletionGracePeriodSeconds(i *int64) {
	o.Inner.SetDeletionGracePeriodSeconds(i)
}

func (o *object[T]) GetLabels() map[string]string {
	return o.Inner.GetLabels()
}

func (o *object[T]) SetLabels(labels map[string]string) {
	o.Inner.SetLabels(labels)
}

func (o *object[T]) GetAnnotations() map[string]string {
	return o.Inner.GetAnnotations()
}

func (o *object[T]) SetAnnotations(annotations map[string]string) {
	o.Inner.SetAnnotations(annotations)
}

func (o *object[T]) GetFinalizers() []string {
	return o.Inner.GetFinalizers()
}

func (o *object[T]) SetFinalizers(finalizers []string) {
	o.Inner.SetFinalizers(finalizers)
}

func (o *object[T]) GetOwnerReferences() []metav1.OwnerReference {
	return o.Inner.GetOwnerReferences()
}

func (o *object[T]) SetOwnerReferences(references []metav1.OwnerReference) {
	o.Inner.SetOwnerReferences(references)
}

func (o *object[T]) GetManagedFields() []metav1.ManagedFieldsEntry {
	return o.Inner.GetManagedFields()
}

func (o *object[T]) SetManagedFields(managedFields []metav1.ManagedFieldsEntry) {
	o.Inner.SetManagedFields(managedFields)
}

func NewObject[T HasMeta](resource T) *object[T] {
	return &object[T]{
		Inner: resource,
	}
}

func (o *object[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Inner)
}

func (o *object[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.Inner)
}

func (o *object[T]) DeepCopyObject() runtime.Object {
	panic("implement me")
}

func (o *object[T]) GetObjectKind() schema.ObjectKind {
	//TODO implement me
	panic("implement me")
}

type ObjectList[T HasMeta] struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []object[T] `json:"items"`
}

func (r *ObjectList[T]) DeepCopyObject() runtime.Object {
	panic("implement me")
}

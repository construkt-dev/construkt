package main

import (
	"context"
	"fmt"
	"github.com/construkt-dev/construkt"
)

func main() {
	ctx := construkt.NewContext()
	cr := construkt.NewResource("examples.construkt.dev", "MyCustomResource", "v1", &MyCustomResource{}).
		Version("v1alpha1", &MyCustomResourceV1alpha1{})

	cli, _ := cr.Client(ctx)

	list, err := cli.List(context.Background())
	if err != nil {
		panic(err)
	}

	for _, item := range list.Items {
		fmt.Printf("%s/%s\n", item.GetNamespace(), item.GetName())
	}
	ctx.Wait()
}

type MyCustomResource = MyCustomResourceV1
type MyCustomResourceV1 struct {
	construkt.Meta
	Spec MyCustomResourceV1Spec `json:"spec"`
}

type MyCustomResourceV1Spec struct {
	New int `json:"new"`
}

type MyCustomResourceV1alpha1 struct {
	construkt.Meta
	Spec MyCustomResourceV1alpha1Spec `json:"spec"`
}

type MyCustomResourceV1alpha1Spec struct {
	Old int `json:"old"`
}

func (m *MyCustomResourceV1alpha1) From(a *MyCustomResource) {
	m.Meta = a.Meta
	m.Spec.Old = a.Spec.New
}

func (m *MyCustomResourceV1alpha1) Into() *MyCustomResource {
	return &MyCustomResourceV1{
		Meta: m.Meta,
		Spec: MyCustomResourceV1Spec{
			New: m.Spec.Old,
		},
	}
}

package jsonschema

import (
	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"testing"
)

type SomeStruct struct {
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		expected v1.JSONSchemaProps
	}{
		{"TestConvert", struct {
			SomeField int
		}{}, v1.JSONSchemaProps{
			Properties: map[string]v1.JSONSchemaProps{
				"someField": {
					Type: "integer",
				},
			},
		}},
		{"Nested struct", struct {
			SomeField struct{ NestedField int }
		}{}, v1.JSONSchemaProps{
			Properties: map[string]v1.JSONSchemaProps{
				"someField": {
					Type: "object",
				},
			},
		}},
		{"Nested named struct", struct {
			SomeField SomeStruct
		}{}, v1.JSONSchemaProps{
			Properties: map[string]v1.JSONSchemaProps{
				"someField": {
					Type: "object",
				},
			},
		}},
	}

	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := ToJsonSchema(test.obj)
			if !cmp.Equal(result, test.expected) {
				tt.Errorf("result differs from expected: %s", cmp.Diff(test.expected, result))
			}
		})
	}
}

package jsonschema

import (
	"github.com/invopop/jsonschema"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"unicode"
	"unicode/utf8"
)

func ToJsonSchema(obj interface{}) v1.JSONSchemaProps {
	reflector := &jsonschema.Reflector{}
	reflector.KeyNamer = firstToLower
	reflector.DoNotReference = true
	schema, err := reflector.Reflect(obj).MarshalJSON()
	if err != nil {
		panic(err)
	}

	result := v1.JSONSchemaProps{}
	err = json.Unmarshal(schema, &result)
	if err != nil {
		panic(err)
	}

	result.Schema = ""
	delete(result.Properties, "metadata")
	result = removeAdditionalProperties(result)
	result = removePatternProperties(result)

	return result
}

func removeAdditionalProperties(schema v1.JSONSchemaProps) v1.JSONSchemaProps {
	schema.AdditionalProperties = nil
	for key, props := range schema.Properties {
		schema.Properties[key] = removeAdditionalProperties(props)
	}
	return schema
}

func removePatternProperties(schema v1.JSONSchemaProps) v1.JSONSchemaProps {
	schema.PatternProperties = nil
	for key, props := range schema.Properties {
		schema.Properties[key] = removePatternProperties(props)
	}
	return schema
}

func firstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

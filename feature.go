package twins

import "encoding/json"

type Feature struct {
	Name              string     `json:"name"`
	Properties        Properties `json:"properties,omitempty"`
	DesiredProperties Properties `json:"desired,omitempty"`
	Attributes        Attributes `json:"attributes,omitempty"`
}

func (feature *Feature) WithName(name string) *Feature {
	feature.Name = name
	return feature
}

func (feature *Feature) WithAttributes(attrs Attributes) *Feature {
	feature.Attributes = attrs
	return feature
}

func (feature *Feature) WithAttribute(id string, value string) *Feature {
	if feature.Attributes == nil {
		feature.Attributes = make(Attributes)
	}
	feature.Attributes[id] = value
	return feature
}

func (feature *Feature) WithDesiredProperties(properties map[string]interface{}) *Feature {
	feature.DesiredProperties = properties
	return feature
}

func (feature *Feature) WithDesiredProperty(id string, value interface{}) *Feature {
	if feature.DesiredProperties == nil {
		feature.DesiredProperties = make(map[string]interface{})
	}
	feature.DesiredProperties[id] = value
	return feature
}

func (feature *Feature) WithProperties(properties map[string]interface{}) *Feature {
	feature.Properties = properties
	return feature
}

func (feature *Feature) WithProperty(id string, value interface{}) *Feature {
	if feature.Properties == nil {
		feature.Properties = make(map[string]interface{})
	}
	feature.Properties[id] = value
	return feature
}

func (t *Feature) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type FeatureList map[string]*Feature

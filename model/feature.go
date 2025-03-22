package model

import "encoding/json"

type Feature struct {
	Name       string     `json:"name"`
	Metrics    Metrics    `json:"metrics,omitempty"`
	Dimensions Dimensions `json:"dimensions,omitempty"`
}

func UnmarshalFeature(buf []byte, msg *Feature) error {
	return json.Unmarshal(buf, msg)
}

func MarshalFeature(msg *Feature) ([]byte, error) {
	return json.Marshal(msg)
}

func (feature *Feature) WithName(name string) *Feature {
	feature.Name = name
	return feature
}

func (feature *Feature) WithAttributes(dims Dimensions) *Feature {
	feature.Dimensions = dims
	return feature
}

func (feature *Feature) WithAttribute(id string, value string) *Feature {
	if feature.Dimensions == nil {
		feature.Dimensions = make(Dimensions)
	}
	feature.Dimensions[id] = value
	return feature
}

func (feature *Feature) WithMetrics(metrics map[string]interface{}) *Feature {
	feature.Metrics = metrics
	return feature
}

func (feature *Feature) WithMetric(id string, value interface{}) *Feature {
	if feature.Metrics == nil {
		feature.Metrics = make(map[string]interface{})
	}
	feature.Metrics[id] = value
	return feature
}

func (t *Feature) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type FeatureList map[string]*Feature

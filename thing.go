package twins

import "encoding/json"

type Thing struct {
	Name       string      `json:"name"`
	Attributes Attributes  `json:"attributes,omitempty"`
	Features   FeatureList `json:"features,omitempty"`
	Revision   int64       `json:"revision,omitempty"`
	Time       string      `json:"time,omitempty"`
}

func (thing *Thing) WithName(name string) *Thing {
	thing.Name = name
	return thing
}

func (thing *Thing) WithRevision(revision int64) *Thing {
	thing.Revision = revision
	return thing
}

func (thing *Thing) WithTime(time string) *Thing {
	thing.Time = time
	return thing
}

func (thing *Thing) WithAttributes(attrs Attributes) *Thing {
	thing.Attributes = attrs
	return thing
}

func (thing *Thing) WithAttribute(id string, value string) *Thing {
	if thing.Attributes == nil {
		thing.Attributes = make(Attributes)
	}
	thing.Attributes[id] = value
	return thing
}

func (thing *Thing) WithFeatures(features FeatureList) *Thing {
	thing.Features = features
	return thing
}

func (thing *Thing) WithFeature(id string, value *Feature) *Thing {
	if thing.Features == nil {
		thing.Features = make(FeatureList)
	}
	thing.Features[id] = value
	return thing
}

func (t *Thing) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

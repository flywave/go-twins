package twins

import "encoding/json"

type StrategyList map[string]*Strategy

type Strategy struct {
	Name       string     `json:"name"`
	Attributes Attributes `json:"attributes,omitempty"`
	Indicators Indicators `json:"indicators,omitempty"`
}

func (st *Strategy) WithName(name string) *Strategy {
	st.Name = name
	return st
}

func (st *Strategy) WithAttributes(attrs Attributes) *Strategy {
	st.Attributes = attrs
	return st
}

func (st *Strategy) WithAttribute(id string, value string) *Strategy {
	if st.Attributes == nil {
		st.Attributes = make(Attributes)
	}
	st.Attributes[id] = value
	return st
}

func (st *Strategy) WithIndicators(indicators map[string]interface{}) *Strategy {
	st.Indicators = indicators
	return st
}

func (st *Strategy) WithIndicator(id string, value interface{}) *Strategy {
	if st.Indicators == nil {
		st.Indicators = make(map[string]interface{})
	}
	st.Indicators[id] = value
	return st
}

func (t *Strategy) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

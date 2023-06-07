package signals

import (
	"encoding/json"
)

type ErrorPayload struct {
	Status      int64                  `json:"status,omitempty"`
	Error       string                 `json:"error,omitempty"`
	Description string                 `json:"description"`
	Props       map[string]interface{} `json:"props,omitempty"`
}

func (p *ErrorPayload) UnmarshalJSON(b []byte) error {
	var kvp map[string]interface{}
	err := json.Unmarshal(b, &kvp)
	if err != nil {
		return err
	}
	for key, value := range kvp {
		switch key {
		case "status":
			p.Status = int64(value.(int))
		case "error":
			p.Error = value.(string)
		default:
			p.Props[key] = value
		}
	}
	return nil
}

func (p ErrorPayload) MarshalJSON() ([]byte, error) {
	kvp := make(map[string]interface{})

	for key, value := range p.Props {
		kvp[key] = value
	}

	kvp["status"] = p.Status
	kvp["error"] = p.Error

	return json.Marshal(kvp)
}

type EventType string

const (
	EVENT_TYPE_THING      EventType = "thing"
	EVENT_TYPE_STREAM     EventType = "stream"
	EVENT_TYPE_CONNECTION EventType = "connection"
	EVENT_TYPE_DEVICE     EventType = "device"
	EVENT_TYPE_TIMESERIES EventType = "timeseries"
)

type EventPayload struct {
	Type        EventType              `json:"type"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Content     string                 `json:"content,omitempty"`
	Props       map[string]interface{} `json:"props,omitempty"`
}

func NewEventPayload() *EventPayload {
	return &EventPayload{Props: make(map[string]interface{})}
}

func (p *EventPayload) UnmarshalJSON(b []byte) error {
	var kvp map[string]interface{}
	err := json.Unmarshal(b, &kvp)
	if err != nil {
		return err
	}
	for key, value := range kvp {
		switch key {
		case "type":
			p.Type = EventType(value.(string))
		case "name":
			p.Name = value.(string)
		case "description":
			p.Description = value.(string)
		case "content":
			p.Content = value.(string)
		default:
			p.Props[key] = value
		}
	}
	return nil
}

func (p EventPayload) MarshalJSON() ([]byte, error) {
	kvp := make(map[string]interface{})

	for key, value := range p.Props {
		kvp[key] = value
	}

	kvp["type"] = p.Type
	kvp["name"] = p.Name
	if p.Description != "" {
		kvp["description"] = p.Description
	}
	if p.Content != "" {
		kvp["content"] = p.Content
	}

	return json.Marshal(kvp)
}

type AlarmSeverity string

const (
	ALARM_SEVERITY_CRITICAL      AlarmSeverity = "critical"
	ALARM_SEVERITY_MAJOR         AlarmSeverity = "major"
	ALARM_SEVERITY_MINOR         AlarmSeverity = "minor"
	ALARM_SEVERITY_WARNING       AlarmSeverity = "warning"
	ALARM_SEVERITY_INDETERMINATE AlarmSeverity = "indeterminate"
)

type AlarmPayload struct {
	Severity    AlarmSeverity          `json:"severity"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Content     string                 `json:"content,omitempty"`
	Props       map[string]interface{} `json:"props,omitempty"`
}

func (p *AlarmPayload) UnmarshalJSON(b []byte) error {
	var kvp map[string]interface{}
	err := json.Unmarshal(b, &kvp)
	if err != nil {
		return err
	}
	for key, value := range kvp {
		switch key {
		case "severity":
			p.Severity = AlarmSeverity(value.(string))
		case "name":
			p.Name = value.(string)
		case "description":
			p.Description = value.(string)
		case "content":
			p.Content = value.(string)
		default:
			p.Props[key] = value
		}
	}
	return nil
}

func (p AlarmPayload) MarshalJSON() ([]byte, error) {
	kvp := make(map[string]interface{})

	for key, value := range p.Props {
		kvp[key] = value
	}

	kvp["severity"] = p.Severity
	kvp["name"] = p.Name
	if p.Description != "" {
		kvp["description"] = p.Description
	}
	if p.Content != "" {
		kvp["content"] = p.Content
	}

	return json.Marshal(kvp)
}

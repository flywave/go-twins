package signals

import (
	"encoding/json"
	"errors"

	"github.com/flywave/go-twins/protocol"
)

type Alarm struct {
	Topic   *protocol.Topic `json:"topic"`
	Path    *protocol.Path  `json:"path,omitempty"`
	Payload interface{}     `json:"payload,omitempty"`
}

func NewAlarmWithEnvelope(en *protocol.Envelope) (*Alarm, error) {
	if en.Topic.IsAlarm() {
		var ap AlarmPayload
		if buf, err := json.Marshal(en.Value); err != nil {
			return nil, err
		} else {
			err := json.Unmarshal(buf, &ap)
			if err != nil {
				return nil, err
			}
		}
		return &Alarm{Topic: en.Topic, Path: en.Path, Payload: ap}, nil
	}
	return nil, errors.New("Envelope is not alarm!")
}

func NewAlarm(ns string, channel string) *Alarm {
	return &Alarm{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityThings).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionAlarms),
		Path: protocol.NewRootPath(),
	}
}

func UnmarshalAlarm(buf []byte, msg *Alarm) error {
	return json.Unmarshal(buf, msg)
}

func MarshalAlarm(msg *Alarm) ([]byte, error) {
	return json.Marshal(msg)
}

func (c *Alarm) GetTopic() string {
	return c.Topic.String()
}

func (c Alarm) GetType() SignalType {
	return SignalTypeAlarm
}

func (c *Alarm) GetPayload() interface{} {
	return c.Payload
}

func (c *Alarm) GetPath() string {
	return c.Path.String()
}

func (c *Alarm) Create(alarm *AlarmPayload) *Alarm {
	c.Payload = alarm
	return c
}

func (c *Alarm) Thing(thingName string) *Alarm {
	c.Path.WithThing(thingName)
	return c
}

func (c *Alarm) Features(thingName string) *Alarm {
	c.Path.WithThingFeatures(thingName)
	return c
}

func (c *Alarm) Feature(thingName, featureName string) *Alarm {
	c.Path.WithThingFeature(thingName, featureName)
	return c
}

func (c *Alarm) FeatureProperty(thingName, featureName, propertyPath string) *Alarm {
	c.Path.WithThingFeaturePropertie(thingName, featureName, propertyPath)
	return c
}

func (c *Alarm) Devices(deviceName string) *Alarm {
	c.Path.WithDevice(deviceName)
	return c
}

func (c *Alarm) Connections(connName string) *Alarm {
	c.Path.WithConnection(connName)
	return c
}

func (c *Alarm) Streams(streamName string) *Alarm {
	c.Path.WithStream(streamName)
	return c
}

func (c *Alarm) Channel(channel string) *Alarm {
	c.Topic.WithChannelName(channel)
	return c
}

func (c *Alarm) Envelope(headerOpts ...HeaderOpt) *protocol.Envelope {
	msg := &protocol.Envelope{
		Topic: c.Topic,
		Path:  c.Path,
		Value: c.Payload,
	}
	if headerOpts != nil {
		msg.Headers = NewHeaders(headerOpts...)
	}
	return msg
}

package signals

import (
	"encoding/json"
	"errors"

	"github.com/flywave/go-twins/protocol"
)

type Event struct {
	Topic   *protocol.Topic `json:"topic"`
	Path    *protocol.Path  `json:"path,omitempty"`
	Payload interface{}     `json:"payload,omitempty"`
}

func NewEventWithEnvelope(en *protocol.Envelope) (*Event, error) {
	if en.Topic.IsEvent() {
		ap := NewEventPayload()
		if buf, err := json.Marshal(en.Value); err != nil {
			return nil, err
		} else {
			err := json.Unmarshal(buf, ap)
			if err != nil {
				return nil, err
			}
		}
		return &Event{Topic: en.Topic, Path: en.Path, Payload: ap}, nil
	}
	return nil, errors.New("Envelope is not event!")
}

func NewEvent(ns string, channel string, entity protocol.EntityType) *Event {
	return &Event{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(entity).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionEvents),
		Path: protocol.NewRootPath(),
	}
}

func NewEventForThing(ns string, channel string) *Event {
	return &Event{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityThings).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionEvents),
		Path: protocol.NewRootPath(),
	}
}

func NewEventForDevice(ns string, channel string) *Event {
	return &Event{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityDevices).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionEvents),
		Path: protocol.NewRootPath(),
	}
}

func NewEventForConnection(ns string, channel string) *Event {
	return &Event{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityConnections).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionEvents),
		Path: protocol.NewRootPath(),
	}
}

func NewEventForStream(ns string, channel string) *Event {
	return &Event{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityStreams).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionEvents),
		Path: protocol.NewRootPath(),
	}
}

func UnmarshalEvent(buf []byte, msg *Event) error {
	return json.Unmarshal(buf, msg)
}

func MarshalEvent(msg *Event) ([]byte, error) {
	return json.Marshal(msg)
}

func (c Event) GetType() SignalType {
	return SignalTypeEvent
}

func (c *Event) GetTopic() string {
	return c.Topic.String()
}

func (c *Event) GetPayload() interface{} {
	return c.Payload
}

func (c *Event) GetPath() string {
	return c.Path.String()
}

func (event *Event) Created(e *EventPayload) *Event {
	event.Topic.WithAction(protocol.ActionCreated)
	event.Payload = e
	return event
}

func (event *Event) Modified(e *EventPayload) *Event {
	event.Topic.WithAction(protocol.ActionModified)
	event.Payload = e
	return event
}

func (event *Event) Deleted(e *EventPayload) *Event {
	event.Topic.WithAction(protocol.ActionDeleted)
	event.Payload = e
	return event
}

func (event *Event) Cleared(e *EventPayload) *Event {
	event.Topic.WithAction(protocol.ActionCleared)
	event.Payload = e
	return event
}

func (event *Event) Thing(thingName string) *Event {
	event.Path.WithThing(thingName)
	return event
}

func (event *Event) Attributes(thingName string) *Event {
	event.Path.WithThingAttributes(thingName)
	return event
}

func (event *Event) Attribute(thingName, attributePath string) *Event {
	event.Path.WithThingAttribute(thingName, attributePath)
	return event
}

func (event *Event) Features(thingName string) *Event {
	event.Path.WithThingFeatures(thingName)
	return event
}

func (event *Event) Feature(thingName, featureName string) *Event {
	event.Path.WithThingFeature(thingName, featureName)
	return event
}

func (event *Event) FeatureProperties(thingName, featureName string) *Event {
	event.Path.WithThingFeatureProperties(thingName, featureName)
	return event
}

func (event *Event) FeatureProperty(thingName, featureName, propertyPath string) *Event {
	event.Path.WithThingFeaturePropertie(thingName, featureName, propertyPath)
	return event
}

func (event *Event) FeaturePropertyTimeseries(thingName, featureName, propertyPath string) *Event {
	event.Path.WithThingFeaturePropertiesTimeSeries(thingName, featureName, propertyPath)
	return event
}

func (event *Event) FeatureDesiredProperties(thingName, featureName string) *Event {
	event.Path.WithThingFeatureDesireds(thingName, featureName)
	return event
}

func (event *Event) FeatureDesiredProperty(thingName, featureName, propertyPath string) *Event {
	event.Path.WithThingFeatureDesired(thingName, featureName, propertyPath)
	return event
}

func (event *Event) Devices(deviceName string) *Event {
	event.Path.WithDevice(deviceName)
	return event
}

func (event *Event) DeviceAttributes(deviceName string) *Event {
	event.Path.WithDeviceAttributes(deviceName)
	return event
}

func (event *Event) DeviceAttribute(deviceName, propertyPath string) *Event {
	event.Path.WithDeviceAttribute(deviceName, propertyPath)
	return event
}

func (event *Event) DeviceHealthStatus(deviceName string) *Event {
	event.Path.WithDeviceStatus(deviceName)
	return event
}

func (event *Event) DeviceStrategys(deviceName string) *Event {
	event.Path.WithDeviceStrategys(deviceName)
	return event
}

func (event *Event) DeviceStrategy(deviceName, strategyName string) *Event {
	event.Path.WithDeviceStrategy(deviceName, strategyName)
	return event
}

func (event *Event) DeviceIndicators(deviceName, strategyName string) *Event {
	event.Path.WithDeviceStrategyIndicators(deviceName, strategyName)
	return event
}

func (event *Event) DeviceIndicator(deviceName, strategyName, indicatorName string) *Event {
	event.Path.WithDeviceStrategyIndicator(deviceName, strategyName, indicatorName)
	return event
}

func (event *Event) DeviceProfiles(deviceName string) *Event {
	event.Path.WithDeviceProfiles(deviceName)
	return event
}

func (event *Event) DeviceProfile(deviceName, profileName string) *Event {
	event.Path.WithDeviceProfile(deviceName, profileName)
	return event
}

func (event *Event) Connections(connName string) *Event {
	event.Path.WithConnection(connName)
	return event
}

func (event *Event) ConnectionStatus(connName string) *Event {
	event.Path.WithConnectionStatus(connName)
	return event
}

func (event *Event) Streams(streamName string) *Event {
	event.Path.WithStream(streamName)
	return event
}

func (event *Event) StreamStatus(streamName string) *Event {
	event.Path.WithStreamStatus(streamName)
	return event
}

func (event *Event) StreamVideos(streamName string) *Event {
	event.Path.WithStreamVideos(streamName)
	return event
}

func (event *Event) StreamAudios(streamName string) *Event {
	event.Path.WithStreamAudios(streamName)
	return event
}

func (event *Event) StreamSubscribers(streamName string) *Event {
	event.Path.WithStreamSubscribers(streamName)
	return event
}

func (event *Event) Channel(channel string) *Event {
	event.Topic.WithChannelName(channel)
	return event
}

func (event *Event) Envelope(headerOpts ...HeaderOpt) *protocol.Envelope {
	msg := &protocol.Envelope{
		Topic: event.Topic,
		Path:  event.Path,
		Value: event.Payload,
	}
	if headerOpts != nil {
		msg.Headers = NewHeaders(headerOpts...)
	}
	return msg
}

package signals

import (
	"encoding/json"
	"errors"

	"github.com/flywave/go-twins/protocol"
)

type Errors struct {
	Topic   *protocol.Topic `json:"topic"`
	Path    *protocol.Path  `json:"path,omitempty"`
	Payload interface{}     `json:"payload,omitempty"`
	Status  int             `json:"status,omitempty"`
}

func NewErrorsWithEnvelope(en *protocol.Envelope) (*Errors, error) {
	if en.Topic.IsError() {
		var ep ErrorPayload
		if buf, err := json.Marshal(en.Value); err != nil {
			return nil, err
		} else {
			err := json.Unmarshal(buf, &ep)
			if err != nil {
				return nil, err
			}
		}
		return &Errors{Topic: en.Topic, Path: en.Path, Payload: ep, Status: en.Status}, nil
	}
	return nil, errors.New("Envelope is not errors!")
}

func NewErrorsForThing(ns string, channel string) *Errors {
	return &Errors{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityThings).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionErrors),
		Path: protocol.NewRootPath(),
	}
}

func NewErrorsForDevice(ns string, channel string) *Errors {
	return &Errors{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityDevices).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionErrors),
		Path: protocol.NewRootPath(),
	}
}

func NewErrorsForConnection(ns string, channel string) *Errors {
	return &Errors{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityConnections).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionErrors),
		Path: protocol.NewRootPath(),
	}
}

func NewErrorsForStream(ns string, channel string) *Errors {
	return &Errors{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityStreams).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionErrors),
		Path: protocol.NewRootPath(),
	}
}

func UnmarshalErrors(buf []byte, msg *Alarm) error {
	return json.Unmarshal(buf, msg)
}

func MarshalErrors(msg *Alarm) ([]byte, error) {
	return json.Marshal(msg)
}

func (err *Errors) GetTopic() string {
	return err.Topic.String()
}

func (err Errors) GetType() SignalType {
	return SignalTypeErrors
}

func (err *Errors) GetPayload() interface{} {
	return err.Payload
}

func (err *Errors) GetPath() string {
	return err.Path.String()
}

func (err *Errors) Created(payload *ErrorPayload) *Errors {
	err.Topic.WithAction(protocol.ActionCreated)
	err.Payload = payload
	return err
}

func (err *Errors) WithStatus(status int) *Errors {
	err.Status = status
	return err
}

func (err *Errors) Modified(payload *ErrorPayload) *Errors {
	err.Payload = payload
	return err
}

func (err *Errors) Deleted() *Errors {
	err.Topic.WithAction(protocol.ActionDeleted)
	return err
}

func (err *Errors) Cleared() *Errors {
	err.Topic.WithAction(protocol.ActionCleared)
	return err
}

func (err *Errors) Thing(thingName string) *Errors {
	err.Path.WithThing(thingName)
	return err
}

func (err *Errors) ThingAttributes(thingName string) *Errors {
	err.Path.WithThingAttributes(thingName)
	return err
}

func (err *Errors) ThingAttribute(thingName, attributePath string) *Errors {
	err.Path.WithThingAttribute(thingName, attributePath)
	return err
}

func (err *Errors) Features(thingName string) *Errors {
	err.Path.WithThingFeatures(thingName)
	return err
}

func (err *Errors) Feature(thingName, featureName string) *Errors {
	err.Path.WithThingFeature(thingName, featureName)
	return err
}

func (err *Errors) FeatureProperties(thingName, featureName string) *Errors {
	err.Path.WithThingFeatureProperties(thingName, featureName)
	return err
}

func (err *Errors) FeatureProperty(thingName, featureName, propertyPath string) *Errors {
	err.Path.WithThingFeaturePropertie(thingName, featureName, propertyPath)
	return err
}

func (err *Errors) FeaturePropertyTimeseries(thingName, featureName, propertyPath string) *Errors {
	err.Path.WithThingFeaturePropertiesTimeSeries(thingName, featureName, propertyPath)
	return err
}

func (err *Errors) FeatureDesiredProperties(thingName, featureName string) *Errors {
	err.Path.WithThingFeatureDesireds(thingName, featureName)
	return err
}

func (err *Errors) FeatureDesiredProperty(thingName, featureName, propertyPath string) *Errors {
	err.Path.WithThingFeatureDesired(thingName, featureName, propertyPath)
	return err
}

func (err *Errors) Devices(deviceName string) *Errors {
	err.Path.WithDevice(deviceName)
	return err
}

func (err *Errors) DeviceAttributes(deviceName string) *Errors {
	err.Path.WithDeviceAttributes(deviceName)
	return err
}

func (err *Errors) DeviceAttribute(deviceName, propertyPath string) *Errors {
	err.Path.WithDeviceAttribute(deviceName, propertyPath)
	return err
}

func (err *Errors) DeviceHealthStatus(deviceName string) *Errors {
	err.Path.WithDeviceStatus(deviceName)
	return err
}

func (err *Errors) DeviceStrategys(deviceName string) *Errors {
	err.Path.WithDeviceStrategys(deviceName)
	return err
}

func (err *Errors) DeviceStrategy(deviceName, strategyName string) *Errors {
	err.Path.WithDeviceStrategy(deviceName, strategyName)
	return err
}

func (err *Errors) DeviceIndicators(deviceName, strategyName string) *Errors {
	err.Path.WithDeviceStrategyIndicators(deviceName, strategyName)
	return err
}

func (err *Errors) DeviceIndicator(deviceName, strategyName, indicatorName string) *Errors {
	err.Path.WithDeviceStrategyIndicator(deviceName, strategyName, indicatorName)
	return err
}

func (err *Errors) DeviceProfiles(deviceName string) *Errors {
	err.Path.WithDeviceProfiles(deviceName)
	return err
}

func (err *Errors) DeviceProfile(deviceName, profileName string) *Errors {
	err.Path.WithDeviceProfile(deviceName, profileName)
	return err
}

func (err *Errors) Connections(connName string) *Errors {
	err.Path.WithConnection(connName)
	return err
}

func (err *Errors) ConnectionStatus(connName string) *Errors {
	err.Path.WithConnectionStatus(connName)
	return err
}

func (err *Errors) Streams(streamName string) *Errors {
	err.Path.WithStream(streamName)
	return err
}

func (err *Errors) StreamStatus(streamName string) *Errors {
	err.Path.WithStreamStatus(streamName)
	return err
}

func (err *Errors) StreamVideos(streamName string) *Errors {
	err.Path.WithStreamVideos(streamName)
	return err
}

func (err *Errors) StreamAudios(streamName string) *Errors {
	err.Path.WithStreamAudios(streamName)
	return err
}

func (err *Errors) StreamSubscribers(streamName string) *Errors {
	err.Path.WithStreamSubscribers(streamName)
	return err
}

func (err *Errors) Channel(channel string) *Errors {
	err.Topic.WithChannelName(channel)
	return err
}

func (err *Errors) Envelope(headerOpts ...HeaderOpt) *protocol.Envelope {
	msg := &protocol.Envelope{
		Topic: err.Topic,
		Path:  err.Path,
		Value: err.Payload,
	}
	if headerOpts != nil {
		msg.Headers = NewHeaders(headerOpts...)
	}
	return msg
}

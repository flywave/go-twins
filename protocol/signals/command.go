package signals

import (
	"encoding/json"
	"errors"

	"github.com/flywave/go-twins/protocol"
)

type Command struct {
	Topic   *protocol.Topic `json:"topic"`
	Path    *protocol.Path  `json:"path,omitempty"`
	Payload interface{}     `json:"payload,omitempty"`
}

func NewCommandWithEnvelope(en *protocol.Envelope) (*Command, error) {
	if en.Topic.IsCommand() {
		return &Command{Topic: en.Topic, Path: en.Path, Payload: en.Value}, nil
	}
	return nil, errors.New("Envelope is not command!")
}

func NewCommand(ns string, channel string, entity protocol.EntityType) *Command {
	return &Command{
		Topic: (&protocol.Topic{}).
			WithEntity(entity).
			WithTenantName(ns).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionCommands),
		Path: protocol.NewRootPath(),
	}
}

func NewCommandForThing(ns string, channel string) *Command {
	return &Command{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityThings).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionCommands),
		Path: protocol.NewRootPath(),
	}
}

func NewCommandForDevice(ns string, channel string) *Command {
	return &Command{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityDevices).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionCommands),
		Path: protocol.NewRootPath(),
	}
}

func NewCommandForConnection(ns string, channel string) *Command {
	return &Command{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityConnections).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionCommands),
		Path: protocol.NewRootPath(),
	}
}

func NewCommandForStream(ns string, channel string) *Command {
	return &Command{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityStreams).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionCommands),
		Path: protocol.NewRootPath(),
	}
}

func UnmarshalCommand(buf []byte, msg *Command) error {
	return json.Unmarshal(buf, msg)
}

func MarshalCommand(msg *Command) ([]byte, error) {
	return json.Marshal(msg)
}

func (c Command) GetType() SignalType {
	return SignalTypeCommand
}

func (c *Command) GetTopic() string {
	return c.Topic.String()
}

func (c *Command) GetPayload() interface{} {
	return c.Payload
}

func (c *Command) GetPath() string {
	return c.Path.String()
}

func (cmd *Command) CreateOrModify(payload interface{}) *Command {
	cmd.Topic.WithAction(protocol.ActionCreateOrModify)
	cmd.Payload = payload
	return cmd
}

func (cmd *Command) Delete() *Command {
	cmd.Topic.WithAction(protocol.ActionDelete)
	return cmd
}

func (event *Event) Clear() *Event {
	event.Topic.WithAction(protocol.ActionClear)
	return event
}

func (cmd *Command) Thing(thingName string) *Command {
	cmd.Path.WithThing(thingName)
	return cmd
}

func (cmd *Command) ThingAttributes(thingName string) *Command {
	cmd.Path.WithThingAttributes(thingName)
	return cmd
}

func (cmd *Command) ThingAttribute(thingName, attributePath string) *Command {
	cmd.Path.WithThingAttribute(thingName, attributePath)
	return cmd
}

func (cmd *Command) Features(thingName string) *Command {
	cmd.Path.WithThingFeatures(thingName)
	return cmd
}

func (cmd *Command) Feature(thingName, featureName string) *Command {
	cmd.Path.WithThingFeature(thingName, featureName)
	return cmd
}

func (cmd *Command) FeatureProperties(thingName, featureName string) *Command {
	cmd.Path.WithThingFeatureProperties(thingName, featureName)
	return cmd
}

func (cmd *Command) FeatureProperty(thingName, featureName, propertyPath string) *Command {
	cmd.Path.WithThingFeaturePropertie(thingName, featureName, propertyPath)
	return cmd
}

func (cmd *Command) FeaturePropertyTimeseries(thingName, featureName, propertyPath string) *Command {
	cmd.Path.WithThingFeaturePropertiesTimeSeries(thingName, featureName, propertyPath)
	return cmd
}

func (cmd *Command) FeatureDesiredProperties(thingName, featureName string) *Command {
	cmd.Path.WithThingFeatureDesireds(thingName, featureName)
	return cmd
}

func (cmd *Command) FeatureDesiredProperty(thingName, featureName, propertyPath string) *Command {
	cmd.Path.WithThingFeatureDesired(thingName, featureName, propertyPath)
	return cmd
}

func (cmd *Command) Devices(deviceName string) *Command {
	cmd.Path.WithDevice(deviceName)
	return cmd
}

func (cmd *Command) DeviceAttributes(deviceName string) *Command {
	cmd.Path.WithDeviceAttributes(deviceName)
	return cmd
}

func (cmd *Command) DeviceAttribute(deviceName string, propertyPath string) *Command {
	cmd.Path.WithDeviceAttribute(deviceName, propertyPath)
	return cmd
}

func (cmd *Command) DeviceHealthStatus(deviceName string) *Command {
	cmd.Path.WithDeviceStatus(deviceName)
	return cmd
}

func (cmd *Command) DeviceStrategys(deviceName string) *Command {
	cmd.Path.WithDeviceStrategys(deviceName)
	return cmd
}

func (cmd *Command) DeviceStrategy(deviceName, strategyName string) *Command {
	cmd.Path.WithDeviceStrategy(deviceName, strategyName)
	return cmd
}

func (cmd *Command) DeviceIndicators(deviceName, strategyName string) *Command {
	cmd.Path.WithDeviceStrategyIndicators(deviceName, strategyName)
	return cmd
}

func (cmd *Command) DeviceIndicator(deviceName, strategyName, indicatorName string) *Command {
	cmd.Path.WithDeviceStrategyIndicator(deviceName, strategyName, indicatorName)
	return cmd
}

func (cmd *Command) DeviceProfiles(deviceName string) *Command {
	cmd.Path.WithDeviceProfiles(deviceName)
	return cmd
}

func (cmd *Command) DeviceProfile(deviceName, profileName string) *Command {
	cmd.Path.WithDeviceProfile(deviceName, profileName)
	return cmd
}

func (cmd *Command) Connections(connName string) *Command {
	cmd.Path.WithConnection(connName)
	return cmd
}

func (cmd *Command) ConnectionStatus(connName string) *Command {
	cmd.Path.WithConnectionStatus(connName)
	return cmd
}

func (cmd *Command) Streams(streamName string) *Command {
	cmd.Path.WithStream(streamName)
	return cmd
}

func (cmd *Command) StreamStatus(streamName string) *Command {
	cmd.Path.WithStreamStatus(streamName)
	return cmd
}

func (cmd *Command) StreamVideos(streamName string) *Command {
	cmd.Path.WithStreamVideos(streamName)
	return cmd
}

func (cmd *Command) StreamAudios(streamName string) *Command {
	cmd.Path.WithStreamAudios(streamName)
	return cmd
}

func (cmd *Command) StreamSubscribers(streamName string) *Command {
	cmd.Path.WithStreamSubscribers(streamName)
	return cmd
}

func (cmd *Command) Channel(channel string) *Command {
	cmd.Topic.WithChannelName(channel)
	return cmd
}

func (cmd *Command) Envelope(headerOpts ...HeaderOpt) *protocol.Envelope {
	msg := &protocol.Envelope{
		Topic: cmd.Topic,
		Path:  cmd.Path,
		Value: cmd.Payload,
	}
	if headerOpts != nil {
		msg.Headers = NewHeaders(headerOpts...)
	}
	return msg
}

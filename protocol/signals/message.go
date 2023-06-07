package signals

import (
	"encoding/json"
	"errors"

	"github.com/flywave/go-twins/protocol"
)

type Message struct {
	Topic     *protocol.Topic        `json:"topic"`
	Subject   string                 `json:"subject,omitempty"`
	Direction protocol.DirectionType `json:"direction,omitempty"`
	Path      *protocol.Path         `json:"path,omitempty"`
	Payload   interface{}            `json:"payload,omitempty"`
}

func NewMessageWithEnvelope(en *protocol.Envelope) (*Message, error) {
	if en.Topic.IsMessage() && (en.Path.Type() == protocol.PathTypeThingFeatureMessages || en.Path.Type() == protocol.PathTypeThingMessages) {
		var subject string
		var direction protocol.DirectionType
		switch en.Path.Type() {
		case protocol.PathTypeThingFeatureMessages:
			ps := en.Path.GetThingFeatureMessages()
			subject = ps.Subject
			direction = ps.Direction
		case protocol.PathTypeThingMessages:
			ps := en.Path.GetThingMessages()
			subject = ps.Subject
			direction = ps.Direction
		}
		return &Message{Topic: en.Topic, Path: en.Path, Payload: en.Value, Subject: subject, Direction: direction}, nil
	}
	return nil, errors.New("Envelope is not event!")
}

func NewMessage(ns string, channel string) *Message {
	return &Message{
		Topic: (&protocol.Topic{}).
			WithTenantName(ns).
			WithEntity(protocol.EntityThings).
			WithChannelName(channel).
			WithCriterion(protocol.CriterionMessages),
		Path: protocol.NewRootPath(),
	}
}

func UnmarshalMessage(buf []byte, msg *Message) error {
	return json.Unmarshal(buf, msg)
}

func MarshalMessage(msg *Message) ([]byte, error) {
	return json.Marshal(msg)
}

func (c Message) GetType() SignalType {
	return SignalTypeMessage
}

func (c *Message) GetTopic() string {
	return c.Topic.String()
}

func (c *Message) GetPayload() interface{} {
	return c.Payload
}

func (msg *Message) Incoming(subject string) *Message {
	msg.Topic.WithAction(protocol.TopicAction(subject))
	msg.Subject = subject
	msg.Direction = protocol.DirectionIncoming
	return msg
}

func (msg *Message) Outgoing(subject string) *Message {
	msg.Topic.WithAction(protocol.TopicAction(subject))
	msg.Subject = subject
	msg.Direction = protocol.DirectionOutgoing
	return msg
}

func (msg *Message) WithPayload(payload interface{}) *Message {
	msg.Payload = payload
	return msg
}

func (msg *Message) Thing(thingName string) *Message {
	msg.Path.WithThing(thingName)
	return msg
}

func (msg *Message) Feature(thingName, featureName string) *Message {
	msg.Path.WithThingFeature(thingName, featureName)
	return msg
}

func (msg *Message) Envelope(headerOpts ...HeaderOpt) *protocol.Envelope {
	switch en := msg.Path.Entity.(type) {
	case *protocol.ThingPath:
		msg.Path.Entity = &protocol.ThingMessagesPath{ThingPath: *en, Direction: msg.Direction, Subject: msg.Subject}
	case *protocol.ThingFeaturesPath:
		msg.Path.Entity = &protocol.ThingFeatureMessagesPath{ThingFeaturesPath: *en, Direction: msg.Direction, Subject: msg.Subject}
	default:
	}
	res := &protocol.Envelope{
		Topic: msg.Topic,
		Path:  msg.Path,
		Value: msg.Payload,
	}
	if headerOpts != nil {
		res.Headers = NewHeaders(headerOpts...)
	}
	return res
}

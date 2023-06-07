package model

import "encoding/json"

type StreamStatus string

const (
	STREAM_STATUS_OPEN    StreamStatus = "open"
	STREAM_STATUS_CLOSED  StreamStatus = "closed"
	STREAM_STATUS_IDLE    StreamStatus = "idle"
	STREAM_STATUS_FAILED  StreamStatus = "failed"
	STREAM_STATUS_UNKNOWN StreamStatus = "unknown"
)

type Stream struct {
	Name        string       `json:"name"`
	Status      StreamStatus `json:"status"`
	Videos      Metadata     `json:"videos"`
	Audios      Metadata     `json:"audios"`
	Subscribers Subscribers  `json:"subscribers"`
}

func UnmarshalStream(buf []byte, msg *Stream) error {
	return json.Unmarshal(buf, msg)
}

func MarshalStream(msg *Stream) ([]byte, error) {
	return json.Marshal(msg)
}

func (stream *Stream) WithName(name string) *Stream {
	stream.Name = name
	return stream
}

func (stream *Stream) WithStatus(st StreamStatus) *Stream {
	stream.Status = st
	return stream
}

func (stream *Stream) WithVideoMetadatas(attrs Metadata) *Stream {
	stream.Videos = attrs
	return stream
}

func (stream *Stream) WithVideoMetadata(id string, value interface{}) *Stream {
	if stream.Videos == nil {
		stream.Videos = make(Metadata)
	}
	stream.Videos[id] = value
	return stream
}

func (stream *Stream) WithAudioMetadatas(attrs Metadata) *Stream {
	stream.Audios = attrs
	return stream
}

func (stream *Stream) WithAudiosAttribute(id string, value interface{}) *Stream {
	if stream.Audios == nil {
		stream.Audios = make(Metadata)
	}
	stream.Audios[id] = value
	return stream
}

func (stream *Stream) WithSubscribers(subs Subscribers) *Stream {
	stream.Subscribers = subs
	return stream
}

func (stream *Stream) WithSubscriber(id string, value Subscriber) *Stream {
	if stream.Subscribers == nil {
		stream.Subscribers = make(Subscribers)
	}
	stream.Subscribers[id] = value
	return stream
}

func (t *Stream) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

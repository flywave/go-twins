package protocol

import (
	"encoding/json"
	"time"

	"github.com/flywave/go-twins"
)

type Envelope struct {
	Topic    *Topic      `json:"topic"`
	Headers  *Headers    `json:"headers,omitempty"`
	Path     *Path       `json:"path"`
	Value    interface{} `json:"value,omitempty"`
	Status   int         `json:"status,omitempty"`
	Revision int64       `json:"revision,omitempty"`
	Time     time.Time   `json:"time,omitempty"`
}

func (msg *Envelope) WithTopic(topic *Topic) *Envelope {
	msg.Topic = topic
	return msg
}

func (msg *Envelope) WithHeaders(headers *Headers) *Envelope {
	msg.Headers = headers
	return msg
}

func (msg *Envelope) WithPath(path *Path) *Envelope {
	msg.Path = path
	return msg
}

func (msg *Envelope) WithValue(value interface{}) *Envelope {
	msg.Value = value
	return msg
}

func (msg *Envelope) WithStatus(status int) *Envelope {
	msg.Status = status
	return msg
}

func (msg *Envelope) WithRevision(revision int64) *Envelope {
	msg.Revision = revision
	return msg
}

func (msg *Envelope) WithTime(t time.Time) *Envelope {
	msg.Time = t
	return msg
}

func (msg *Envelope) UnmarshalJSON(d []byte) error {
	ps := &struct {
		Time     string      `json:"time,omitempty"`
		Topic    *Topic      `json:"topic"`
		Headers  *Headers    `json:"headers,omitempty"`
		Path     *Path       `json:"path"`
		Value    interface{} `json:"value,omitempty"`
		Status   int         `json:"status,omitempty"`
		Revision int64       `json:"revision,omitempty"`
	}{}

	err := json.Unmarshal(d, ps)
	if err != nil {
		return err
	}
	if ps.Time != "" {
		msg.Time, err = time.Parse(twins.DateTimeFormat, ps.Time)
		if err != nil {
			return err
		}
	}
	msg.Topic = ps.Topic
	msg.Headers = ps.Headers
	msg.Path = ps.Path
	msg.Value = ps.Value
	msg.Status = ps.Status
	msg.Revision = ps.Revision
	return nil
}

func (msg *Envelope) MarshalJSON() ([]byte, error) {
	ps := struct {
		Time     string      `json:"time,omitempty"`
		Topic    *Topic      `json:"topic"`
		Headers  *Headers    `json:"headers,omitempty"`
		Path     *Path       `json:"path"`
		Value    interface{} `json:"value,omitempty"`
		Status   int         `json:"status,omitempty"`
		Revision int64       `json:"revision,omitempty"`
	}{
		Time:     msg.Time.Format(twins.DateTimeFormat),
		Topic:    msg.Topic,
		Headers:  msg.Headers,
		Path:     msg.Path,
		Value:    msg.Value,
		Status:   msg.Status,
		Revision: msg.Revision,
	}

	return json.Marshal(ps)
}

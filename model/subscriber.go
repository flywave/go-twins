package model

import "encoding/json"

type Subscribers map[string]Subscriber

type Subscriber struct {
	Protocol    string       `json:"protocol"`
	Status      StreamStatus `json:"status"`
	SubscribeAt int64        `json:"subscribe_at"`
	CloseAt     int64        `json:"close_at"`
}

func UnmarshalSubscriber(buf []byte, msg *Subscriber) error {
	return json.Unmarshal(buf, msg)
}

func MarshalSubscriber(msg *Subscriber) ([]byte, error) {
	return json.Marshal(msg)
}

func (sub *Subscriber) WithStatus(st StreamStatus) *Subscriber {
	sub.Status = st
	return sub
}

func (sub *Subscriber) WithProtocol(proto string) *Subscriber {
	sub.Protocol = proto
	return sub
}

func (sub *Subscriber) WithSubscribeAt(sat int64) *Subscriber {
	sub.SubscribeAt = sat
	return sub
}

func (sub *Subscriber) WithCloseAt(sat int64) *Subscriber {
	sub.CloseAt = sat
	return sub
}

func (t *Subscriber) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

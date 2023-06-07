package protocol

type Envelope struct {
	Topic     *Topic      `json:"topic"`
	Headers   *Headers    `json:"headers,omitempty"`
	Path      *Path       `json:"path"`
	Value     interface{} `json:"value,omitempty"`
	Fields    string      `json:"fields,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Status    int         `json:"status,omitempty"`
	Revision  int64       `json:"revision,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
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

func (msg *Envelope) WithFields(fields string) *Envelope {
	msg.Fields = fields
	return msg
}

func (msg *Envelope) WithExtra(extra interface{}) *Envelope {
	msg.Extra = extra
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

func (msg *Envelope) WithTimestamp(timestamp string) *Envelope {
	msg.Timestamp = timestamp
	return msg
}

package model

import "encoding/json"

type ConnectivityStatus string

const (
	CONNECTION_STATUS_OPEN    ConnectivityStatus = "open"
	CONNECTION_STATUS_CLOSED  ConnectivityStatus = "closed"
	CONNECTION_STATUS_FAILED  ConnectivityStatus = "failed"
	CONNECTION_STATUS_UNKNOWN ConnectivityStatus = "unknown"
)

type Connection struct {
	Name   string             `json:"name"`
	Status ConnectivityStatus `json:"status"`
	Tags   []string           `json:"tags"`
}

func UnmarshalConnection(buf []byte, msg *Connection) error {
	return json.Unmarshal(buf, msg)
}

func MarshalConnection(msg *Connection) ([]byte, error) {
	return json.Marshal(msg)
}

func (conn *Connection) WithName(name string) *Connection {
	conn.Name = name
	return conn
}

func (conn *Connection) WithStatus(status ConnectivityStatus) *Connection {
	conn.Status = status
	return conn
}

func (conn *Connection) WithTags(tags []string) *Connection {
	conn.Tags = append(conn.Tags, tags...)
	return conn
}

func (conn *Connection) WithTag(tag string) *Connection {
	conn.Tags = append(conn.Tags, tag)
	return conn
}

func (t *Connection) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

package model

import "encoding/json"

type DeviceType string

const (
	DeviceTypeGateway DeviceType = "gateway"
	DeviceTypeSensor  DeviceType = "sensor"
	DeviceTypePLC     DeviceType = "plc"
	DeviceTypeDCS     DeviceType = "dcs"
	DeviceTypeDCM     DeviceType = "dcm"
	DeviceTypeDTU     DeviceType = "dtu"
	DeviceTypeRTU     DeviceType = "rtu"
	DeviceTypeCamera  DeviceType = "camera"
	DeviceTypeMachine DeviceType = "machine"
	DeviceTypeEdge    DeviceType = "edge"
	DeviceTypeUnkown  DeviceType = "unkown"
)

type Profile struct {
	Name         string     `json:"name"`
	Product      string     `json:"product"`
	Manufacturer string     `json:"manufacturer"`
	Type         DeviceType `json:"type"`
	Version      string     `json:"version"`
	Firmware     string     `json:"firmware,omitempty"`
	Protocol     string     `json:"protocol,omitempty"`
	Transport    string     `json:"transport,omitempty"`
	Tags         []string   `json:"tags"`
}

func UnmarshalProfile(buf []byte, msg *Profile) error {
	return json.Unmarshal(buf, msg)
}

func MarshalProfile(msg *Profile) ([]byte, error) {
	return json.Marshal(msg)
}

func (c *Profile) WithName(name string) *Profile {
	c.Name = name
	return c
}

func (c *Profile) WithProduct(pd string) *Profile {
	c.Product = pd
	return c
}

func (c *Profile) WithManufacturer(man string) *Profile {
	c.Manufacturer = man
	return c
}

func (c *Profile) WithType(tp DeviceType) *Profile {
	c.Type = tp
	return c
}

func (c *Profile) WithVersion(vs string) *Profile {
	c.Version = vs
	return c
}

func (c *Profile) WithFirmware(fm string) *Profile {
	c.Firmware = fm
	return c
}

func (c *Profile) WithProtocol(proto string) *Profile {
	c.Protocol = proto
	return c
}

func (c *Profile) WithTransport(tran string) *Profile {
	c.Transport = tran
	return c
}

func (c *Profile) WithTags(tags []string) *Profile {
	c.Tags = append(c.Tags, tags...)
	return c
}

func (c *Profile) WithTag(tag string) *Profile {
	c.Tags = append(c.Tags, tag)
	return c
}

func (c *Profile) ToJson() string {
	b, _ := json.Marshal(c)
	return string(b)
}

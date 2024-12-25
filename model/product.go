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

type Product struct {
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

func UnmarshalProduct(buf []byte, msg *Product) error {
	return json.Unmarshal(buf, msg)
}

func MarshalProduct(msg *Product) ([]byte, error) {
	return json.Marshal(msg)
}

func (c *Product) WithName(name string) *Product {
	c.Name = name
	return c
}

func (c *Product) WithProduct(pd string) *Product {
	c.Product = pd
	return c
}

func (c *Product) WithManufacturer(man string) *Product {
	c.Manufacturer = man
	return c
}

func (c *Product) WithType(tp DeviceType) *Product {
	c.Type = tp
	return c
}

func (c *Product) WithVersion(vs string) *Product {
	c.Version = vs
	return c
}

func (c *Product) WithFirmware(fm string) *Product {
	c.Firmware = fm
	return c
}

func (c *Product) WithProtocol(proto string) *Product {
	c.Protocol = proto
	return c
}

func (c *Product) WithTransport(tran string) *Product {
	c.Transport = tran
	return c
}

func (c *Product) WithTags(tags []string) *Product {
	c.Tags = append(c.Tags, tags...)
	return c
}

func (c *Product) WithTag(tag string) *Product {
	c.Tags = append(c.Tags, tag)
	return c
}

func (c *Product) ToJson() string {
	b, _ := json.Marshal(c)
	return string(b)
}

package model

import "encoding/json"

type HealthStatus string

const (
	HEALTH_STATUS_UNACTIVATED HealthStatus = "unactivated"
	HEALTH_STATUS_UNHEALTHY   HealthStatus = "unhealthy"
	HEALTH_STATUS_HEALTHY     HealthStatus = "healthy"
	HEALTH_STATUS_OFFLINE     HealthStatus = "offline"
)

type Device struct {
	Name         string       `json:"name"`
	SerialNumber string       `json:"serial_number"`
	Product      *Product     `json:"Product,omitempty"`
	Strategys    StrategyList `json:"strategys,omitempty"`
	Status       HealthStatus `json:"status"`
	Attributes   Attributes   `json:"attributes,omitempty"`
}

func UnmarshalDevice(buf []byte, msg *Device) error {
	return json.Unmarshal(buf, msg)
}

func MarshalDevice(msg *Device) ([]byte, error) {
	return json.Marshal(msg)
}

func (dev *Device) WithName(name string) *Device {
	dev.Name = name
	return dev
}

func (dev *Device) WithStatus(status HealthStatus) *Device {
	dev.Status = status
	return dev
}

func (dev *Device) WithSerialNumber(sn string) *Device {
	dev.SerialNumber = sn
	return dev
}

func (dev *Device) WithProduct(pf *Product) *Device {
	dev.Product = pf
	return dev
}

func (dev *Device) WithAttributes(attrs Attributes) *Device {
	dev.Attributes = attrs
	return dev
}

func (dev *Device) WithAttribute(id string, value string) *Device {
	if dev.Attributes == nil {
		dev.Attributes = make(Attributes)
	}
	dev.Attributes[id] = value
	return dev
}

func (dev *Device) WithStrategys(strategyes StrategyList) *Device {
	dev.Strategys = strategyes
	return dev
}

func (dev *Device) WithStrategy(id string, value *Strategy) *Device {
	if dev.Strategys == nil {
		dev.Strategys = make(StrategyList)
	}
	dev.Strategys[id] = value
	return dev
}

func (t *Device) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

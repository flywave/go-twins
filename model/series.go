package model

import (
	"encoding/json"
	"time"

	"github.com/flywave/go-twins"
)

type SeriesPoint struct {
	Time       time.Time         `json:"time,omitempty"`
	Name       string            `json:"name"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Properties Properties        `json:"properties,omitempty"`
}

func (p *SeriesPoint) UnmarshalJSON(d []byte) error {
	ps := &struct {
		Time       string            `json:"time,omitempty"`
		Name       string            `json:"name"`
		Attributes map[string]string `json:"attributes,omitempty"`
		Properties Properties        `json:"properties,omitempty"`
	}{}

	err := json.Unmarshal(d, ps)
	if err != nil {
		return err
	}
	if ps.Time != "" {
		p.Time, err = time.Parse(twins.DateTimeFormat, ps.Time)
		if err != nil {
			return err
		}
	}
	p.Name = ps.Name
	p.Attributes = ps.Attributes
	p.Properties = ps.Properties
	return nil
}

func (p *SeriesPoint) MarshalJSON() ([]byte, error) {
	ps := struct {
		Time       string            `json:"time,omitempty"`
		Name       string            `json:"name"`
		Attributes map[string]string `json:"attributes,omitempty"`
		Properties Properties        `json:"properties,omitempty"`
	}{
		Time:       p.Time.Format(twins.DateTimeFormat),
		Name:       p.Name,
		Attributes: p.Attributes,
		Properties: p.Properties,
	}

	return json.Marshal(ps)
}

func UnmarshalSeriesPoint(buf []byte, msg *SeriesPoint) error {
	return json.Unmarshal(buf, msg)
}

func MarshalSeriesPoint(msg *SeriesPoint) ([]byte, error) {
	return json.Marshal(msg)
}

type Series []SeriesPoint

func (a Series) Len() int { return len(a) }

func (a Series) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }

func (a Series) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func UnmarshalSeries(buf []byte, msg *Series) error {
	return json.Unmarshal(buf, msg)
}

func MarshalSeries(msg *Series) ([]byte, error) {
	return json.Marshal(msg)
}

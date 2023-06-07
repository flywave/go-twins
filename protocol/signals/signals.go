package signals

type SignalType string

const (
	SignalTypeCommand SignalType = "command"
	SignalTypeEvent   SignalType = "event"
	SignalTypeAlarm   SignalType = "alarm"
	SignalTypeMessage SignalType = "message"
	SignalTypeErrors  SignalType = "errors"
)

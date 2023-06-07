package protocol

import (
	"regexp"
)

type PlaceHolders string

const (
	PLACE_HOLDERS_THING_ID              = "thing:id"
	PLACE_HOLDERS_THING_NAME            = "thing:name"
	PLACE_HOLDERS_FEATURE_ID            = "feature:id"
	PLACE_HOLDERS_FEATURE_NAME          = "feature:name"
	PLACE_HOLDERS_DEVICE_ID             = "device:id"
	PLACE_HOLDERS_DEVICE_SN             = "device:serial-number"
	PLACE_HOLDERS_SOURCE_ADDRESS        = "source:address"
	PLACE_HOLDERS_HEADER_REPLY_TO       = "header:reply-to"
	PLACE_HOLDERS_HEADER_CORRELATION    = "header:correlation-id"
	PLACE_HOLDERS_HEADER_CONTENT_TYPE   = "header:content-type"
	PLACE_HOLDERS_HEADER_MESSAGE_ID     = "header:message-id"
	PLACE_HOLDERS_HEADER_DEVICE_ID      = "header:device-id"
	PLACE_HOLDERS_HEADER_QOS            = "header:qos"
	PLACE_HOLDERS_TOPIC_CHANNEL         = "topic:channel"
	PLACE_HOLDERS_TOPIC_CRITERION       = "topic:criterion"
	PLACE_HOLDERS_TOPIC_ACTION          = "topic:action"
	PLACE_HOLDERS_TIME_NOW              = "time:now"
	PLACE_HOLDERS_TIME_NOW_EPOCH_MILLIS = "time:now_epoch_millis"
	PLACE_HOLDERS_HEADER_FORMAT         = "header:%s"

	PLACE_HOLDERS_TEMPLATE_REGEXP = `{{(.*?)}}`
)

type ValueFinder interface {
	GetValue(name string) string
}

func HasPlaceHolders(prop string) bool {
	re := regexp.MustCompile(PLACE_HOLDERS_TEMPLATE_REGEXP)
	return re.Match([]byte(prop))
}

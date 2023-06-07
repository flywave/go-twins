package protocol

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type TopicCriterion string

const (
	CriterionCommands TopicCriterion = "commands"
	CriterionEvents   TopicCriterion = "events"
	CriterionMessages TopicCriterion = "messages"
	CriterionErrors   TopicCriterion = "errors"
	CriterionAlarms   TopicCriterion = "alarms"
)

type EntityType string

const (
	EntityUnknown     EntityType = ""
	EntityThings      EntityType = "things"
	EntityStreams     EntityType = "streams"
	EntityConnections EntityType = "connections"
	EntityDevices     EntityType = "devices"
)

type TopicAction string

const (
	ActionCreateOrModify TopicAction = "createmodify"
	ActionCreated        TopicAction = "created"
	ActionModified       TopicAction = "modified"
	ActionDelete         TopicAction = "delete"
	ActionDeleted        TopicAction = "deleted"
	ActionMerge          TopicAction = "merge"
	ActionMerged         TopicAction = "merged"
	ActionTrigger        TopicAction = "trigger"
	ActionTriggered      TopicAction = "triggered"
	ActionClear          TopicAction = "clear"
	ActionCleared        TopicAction = "cleared"
	ActionSubscribe      TopicAction = "subscribe"
	ActionSubscribed     TopicAction = "subscribed"
	ActionUnSubscribe    TopicAction = "unsubscribe"
	ActionUnSubscribed   TopicAction = "unsubscribed"
	ActionStatusChange   TopicAction = "statuschange"
	ActionStatusChanged  TopicAction = "statuschanged"
	ActionFailed         TopicAction = "failed"
)

const TopicPlaceholder = "_"

const (
	topicFormat         = "@topic/%s/%s/%s/%s/%s"
	topicFormatNoAction = "@topic/%s/%s/%s/%s"
)

/**
* @topic/tenant1/channel1/things/alarms
* @topic/tenant1/channel1/things/commands/create
* @topic/tenant1/channel1/things/events/created
**/
var regexTopic = regexp.MustCompile("^@topic/([^/]+)/([^/]+)/(" + string(EntityThings) + "|" + string(EntityStreams) + "|" + string(EntityConnections) + "|" + string(EntityDevices) + ")/([^/]+)(/([^/]{1}.*))?$")

type Topic struct {
	TenantName  string
	ChannelName string
	Entity      EntityType
	Criterion   TopicCriterion
	Action      TopicAction
}

func NewTopic(tp string) (*Topic, error) {
	topic := &Topic{}
	if err := parseTopic(tp, topic); err != nil {
		return nil, err
	}
	return topic, nil
}

func parseTopic(tp string, topic *Topic) error {
	matches := regexTopic.FindAllStringSubmatch(tp, -1)
	if matches == nil {
		return errors.New("invalid topic: " + tp)
	}

	elements := matches[0]

	topic.TenantName = elements[1]
	topic.ChannelName = elements[2]

	topic.Criterion = TopicCriterion(elements[3])
	topic.Criterion = TopicCriterion(elements[4])

	if elements[6] != "" {
		topic.Action = TopicAction(elements[6])
	}
	return nil
}

func (topic *Topic) String() string {
	if len(topic.Action) == 0 {
		return fmt.Sprintf(topicFormatNoAction, topic.TenantName, topic.ChannelName, topic.Entity, topic.Criterion)
	}
	return fmt.Sprintf(topicFormat, topic.TenantName, topic.ChannelName, topic.Entity, topic.Criterion, topic.Action)
}

func (topic Topic) MarshalJSON() ([]byte, error) {
	topicStr := topic.String()
	matches := regexTopic.FindAllStringSubmatch(topicStr, -1)
	if matches == nil {
		return nil, errors.New("invalid topic: " + topicStr)
	}
	return json.Marshal(topicStr)
}

func (topic *Topic) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	matches := regexTopic.FindAllStringSubmatch(v, -1)
	if matches == nil {
		return errors.New("invalid topic: " + v)
	}

	elements := matches[0]

	topic.TenantName = elements[1]
	topic.ChannelName = elements[2]

	topic.Entity = EntityType(elements[3])
	topic.Criterion = TopicCriterion(elements[4])

	if elements[6] != "" {
		topic.Action = TopicAction(elements[6])
	}
	return nil
}

func (topic *Topic) Value() (driver.Value, error) {
	j, err := json.Marshal(topic)
	if err != nil {
		return nil, err
	}
	return string(j), err
}

func (topic *Topic) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	buf, ok := value.([]byte)
	if ok {
		return json.Unmarshal(buf, topic)
	}

	str, ok := value.(string)
	if ok {
		return json.Unmarshal([]byte(str), topic)
	}

	return errors.New("received value is neither a byte slice nor string")
}

func (topic *Topic) IsCommand() bool {
	return topic.Criterion == CriterionCommands
}

func (topic *Topic) IsEvent() bool {
	return topic.Criterion == CriterionEvents
}

func (topic *Topic) IsMessage() bool {
	return topic.Criterion == CriterionMessages
}

func (topic *Topic) IsError() bool {
	return topic.Criterion == CriterionErrors
}

func (topic *Topic) IsAlarm() bool {
	return topic.Criterion == CriterionAlarms
}

func (topic *Topic) GetNameSpace() string {
	return topic.TenantName
}

func (topic *Topic) WithTenantName(ns string) *Topic {
	topic.TenantName = ns
	return topic
}

func (topic *Topic) WithChannelName(channelName string) *Topic {
	topic.ChannelName = channelName
	return topic
}

func (topic *Topic) WithCriterion(criterion TopicCriterion) *Topic {
	topic.Criterion = criterion
	return topic
}

func (topic *Topic) WithAction(action TopicAction) *Topic {
	topic.Action = action
	return topic
}

func (topic *Topic) WithEntity(entity EntityType) *Topic {
	topic.Entity = entity
	return topic
}

func (topic *Topic) HasWildCard() bool {
	return strings.Contains(topic.String(), "*")
}

func (topic *Topic) Match(pattern string) bool {
	result, _ := regexp.MatchString(WildCardToRegexp(pattern), topic.String())
	return result
}

func (topic *Topic) HasPlaceHolders() bool {
	return HasPlaceHolders(topic.String())
}

type TopicList []Topic

func (e TopicList) Value() (driver.Value, error) {
	j, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(j), err
}

func (e *TopicList) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	buf, ok := value.([]byte)
	if ok {
		return json.Unmarshal(buf, e)
	}

	str, ok := value.(string)
	if ok {
		return json.Unmarshal([]byte(str), e)
	}

	return errors.New("received value is neither a byte slice nor string")
}

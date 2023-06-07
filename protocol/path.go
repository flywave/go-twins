package protocol

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	pathWillCard                            = "*"
	pathRoot                                = "@"
	pathThings                              = "@things"
	pathThingFormat                         = pathThings + "/%s"
	pathThingFeatures                       = pathThingFormat + "/features"
	pathThingAttributes                     = pathThingFormat + "/attributes"
	pathThingAttributeFormat                = pathThingAttributes + "/%s"
	pathThingFeatureFormat                  = pathThingFeatures + "/%s"
	pathThingFeaturePropertiesFormat        = pathThingFeatureFormat + "/properties"
	pathThingFeaturePropertyFormat          = pathThingFeaturePropertiesFormat + "/%s"
	pathThingFeatureDesiredPropertiesFormat = pathThingFeatureFormat + "/desired"
	pathThingFeatureDesiredPropertyFormat   = pathThingFeatureDesiredPropertiesFormat + "/%s"
	pathThingFeaturePropertyTimeseries      = pathThingFeaturePropertyFormat + "/timeseries"
	pathThingFeatureAttributes              = pathThingFeatureFormat + "/attributes"
	pathThingFeatureAttributeFormat         = pathThingFeatureAttributes + "/%s"
	pathDevices                             = "@devices"
	pathDeviceFormat                        = pathDevices + "/%s"
	pathDeviceProperties                    = pathDeviceFormat + "/attributes"
	pathDevicePropertyFormat                = pathDeviceProperties + "/%s"
	pathDeviceHealthStatus                  = pathDeviceFormat + "/status"
	pathDeviceStrategys                     = pathDeviceFormat + "/strategys"
	pathDeviceStrategysFormat               = pathDeviceStrategys + "/%s"
	pathDeviceIndicators                    = pathDeviceStrategysFormat + "/indicators"
	pathDeviceStrategysAttributes           = pathDeviceStrategysFormat + "/attributes"
	pathDeviceStrategysAttributeFormat      = pathDeviceStrategysAttributes + "/%s"
	pathDeviceIndicatorFormat               = pathDeviceIndicators + "/%s"
	pathDeviceIndicatorTimeseries           = pathDeviceIndicatorFormat + "/timeseries"
	pathDeviceProfiles                      = pathDeviceFormat + "/profiles"
	pathDeviceProfileFormat                 = pathDeviceProfiles + "/%s" // (name | product | manufacturer | version | firmware | protocol | transport | tags)
	pathConnections                         = "@connections"
	pathConnectionId                        = "@connectionid"
	pathConnectionFormat                    = pathConnections + "/%s"
	pathConnectionStatus                    = pathConnectionFormat + "/status"
	pathStreams                             = "@streams"
	pathStreamFormat                        = pathStreams + "/%s"
	pathStreamStatus                        = pathStreamFormat + "/status"
	pathStreamVideos                        = pathStreamFormat + "/videos"
	pathStreamAudios                        = pathStreamFormat + "/audios"
	pathConnectionSubscribers               = pathStreamFormat + "/subscribers"
	pathMessagesFormat                      = "%s/messages/%s/%s"
	pathFeatures                            = "@features"
	pathFeatureFormat                       = pathFeatures + "/%s"
	pathFeaturePropertiesFormat             = pathFeatureFormat + "/properties"
	pathFeaturePropertyFormat               = pathFeaturePropertiesFormat + "/%s"
	pathFeatureDesiredPropertiesFormat      = pathFeatureFormat + "/desired"
	pathFeatureDesiredPropertyFormat        = pathFeatureDesiredPropertiesFormat + "/%s"
	pathFeatureAttributes                   = pathFeatureFormat + "/attributes"
	pathFeatureAttributeFormat              = pathFeatureAttributes + "/%s"
	pathProperties                          = "@properties"
	pathPropertieFormat                     = pathProperties + "/%s"
	pathDesireds                            = "@desired"
	pathDesiredFormat                       = pathDesireds + "/%s"
	pathAttributes                          = "@attributes"
	pathAttributeFormat                     = pathAttributes + "/%s"
	pathStrategys                           = "@strategys"
	pathStrategyFormat                      = pathStrategys + "/%s"
	pathStrategysIndicators                 = pathStrategyFormat + "/indicators"
	pathStrategysIndicatorFormat            = pathStrategysIndicators + "/%s"
	pathStrategysAttributes                 = pathStrategyFormat + "/attributes"
	pathStrategysAttributeFormat            = pathStrategysAttributes + "/%s"
	pathIndicators                          = "@indicators"
	pathIndicatorFormat                     = pathIndicators + "/%s"
	pathProfiles                            = "@profiles"
	pathProfileFormat                       = pathProfiles + "/%s"
	pathStatus                              = "@status"
	pathVideos                              = "@videos"
	pathAudios                              = "@audios"
	pathSubscribers                         = "@subscribers"
)

var regexThingsPath = regexp.MustCompile("^@(things)/([^/]+)(/(features)/([^/]+)/(messages)/(incoming|outgoing)/([^/]+)|/(features)(/([^/]+)(/(properties)(/([^/]{1}.*)/(timeseries))?)?)?|/(features)(/([^/]+)(/(properties|desired|attributes)(/([^/]{1}.*))?)?)?|/(attributes)(/([^/]{1}.*))?|/(messages)/(incoming|outgoing)/([^/]+))?$")
var regexDevicesPath = regexp.MustCompile("^@(devices)/([^/]+)(/(status)|/(strategys)(/([^/]+)(/(indicators|attributes)(/([^/]+)(/(timeseries))?)?)?)?|/(attributes)(/([^/]{1}.*))?|/(profiles)(/(name|product|manufacturer|version|firmware|protocol|transport|tags))?)?$")
var regexConnectionsPath = regexp.MustCompile("^@(connections)/([^/]+)(/(status))?$")
var regexStreamsPath = regexp.MustCompile("^@(streams)/([^/]+)(/(status|videos|audios|subscribers))?$")
var regexFeaturesPath = regexp.MustCompile("^@(features)(/([^/]+)(/(properties|desired|attributes)(/([^/]{1}.*))?)?)?$")
var regexPropertiesPath = regexp.MustCompile("^@(properties)(/([^/]{1}.*))?$")
var regexDesiredPath = regexp.MustCompile("^@(desired)(/([^/]{1}.*))?$")
var regexAttributesPath = regexp.MustCompile("^@(attributes)(/([^/]{1}.*))?$")
var regexStrategysPath = regexp.MustCompile("^@(strategys)(/([^/]+)(/(indicators|attributes)(/([^/]+))?)?)?$")
var regexIndicatorsPath = regexp.MustCompile("^@(indicators)(/([^/]+)?)?$")
var regexProfilesPath = regexp.MustCompile("^@(profiles)(/([^/]+)?)?$")

type PathType string

const (
	PathTypeRoot                          PathType = "root_path"
	PathTypeThing                         PathType = "thing_path"
	PathTypeThingAttributes               PathType = "thing_attributes_path"
	PathTypeThingMessages                 PathType = "thing_message_path"
	PathTypeThingFeatures                 PathType = "thing_features_path"
	PathTypeThingFeatureMessages          PathType = "thing_feature_message_path"
	PathTypeThingFeatureProperties        PathType = "thing_feature_properties_path"
	PathTypeThingFeatureDesiredProperties PathType = "thing_feature_desired_path"
	PathTypeThingFeatureAttributes        PathType = "thing_feature_attributes_path"
	PathTypeDevice                        PathType = "device_path"
	PathTypeDeviceStatus                  PathType = "device_status_path"
	PathTypeDeviceAttributes              PathType = "device_attributes_path"
	PathTypeDeviceStrategys               PathType = "device_strategys_path"
	PathTypeDeviceStrategyIndicators      PathType = "device_strategys_indicators_path"
	PathTypeDeviceStrategyAttributes      PathType = "device_strategys_attributes_path"
	PathTypeDeviceProfiles                PathType = "device_profiles_path"
	PathTypeConnection                    PathType = "connection_path"
	PathTypeConnectionStatus              PathType = "connection_status_path"
	PathTypeStream                        PathType = "stream_path"
	PathTypeStreamStatus                  PathType = "stream_status_path"
	PathTypeStreamVideos                  PathType = "stream_videos_path"
	PathTypeStreamAudios                  PathType = "stream_audios_path"
	PathTypeStream_SUBSCRIBERS            PathType = "stream_subscribers_path"
	PathTypeFeatures                      PathType = "feature_path"
	PathTypeFeaturesProperties            PathType = "feature_properties_path"
	PathTypeFeaturesDesiredProperties     PathType = "feature_desired_path"
	PathTypeFeaturesAttributes            PathType = "feature_attributes_path"
	PathTypeProperties                    PathType = "properties_path"
	PathTypeDesiredProperties             PathType = "desired_path"
	PathTypeAttributes                    PathType = "attributes_path"
	PathTypeStrategys                     PathType = "strategys_path"
	PathTypeStrategyIndicators            PathType = "strategys_indicators_path"
	PathTypeStrategyAttributes            PathType = "strategys_attributes_path"
	PathTypeIndicators                    PathType = "indicators_path"
	PathTypeProfiles                      PathType = "profiles_path"
	PathTypeStatus                        PathType = "status_path"
	PathTypeVideos                        PathType = "videos_path"
	PathTypeAudios                        PathType = "audios_path"
	PathTypeSubscribers                   PathType = "subscribers_path"
)

type Path struct {
	Entity EntityPath
}

func NewRootPath() *Path {
	return &Path{Entity: &RootPath{}}
}

func NewPath(str string) (*Path, error) {
	e, err := parseEntityPath(str)
	if err != nil {
		return nil, err
	}
	return &Path{Entity: e}, nil
}

func (p *Path) Empty() bool {
	return p.Entity == nil
}

func (p *Path) Name() string {
	return p.Entity.Name()
}

func (p *Path) String() string {
	if p.Entity == nil {
		return ""
	}
	return p.Entity.String()
}

func (p *Path) Clone() *Path {
	po, _ := NewPath(p.String())
	return po
}

func (p Path) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Path) IsJoinOf(target *Path) bool {
	return isJoinOf(p.Type(), target.Type())
}

func (p *Path) Join(target *Path) (*Path, error) {
	src := p.String()
	targetPath := target.String()

	if strings.HasPrefix(targetPath, "@") {
		targetPath = strings.ReplaceAll(targetPath, "@", "")
	}

	newPath := strings.Join([]string{src, targetPath}, "/")

	return NewPath(newPath)
}

func (p *Path) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	e, err := parseEntityPath(v)
	if err != nil {
		return err
	}
	p.Entity = e
	return nil
}

func (p *Path) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return string(j), err
}

func (p *Path) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	buf, ok := value.([]byte)
	if ok {
		return json.Unmarshal(buf, p)
	}

	str, ok := value.(string)
	if ok {
		return json.Unmarshal([]byte(str), p)
	}

	return errors.New("received value is neither a byte slice nor string")
}

func (p *Path) Type() PathType {
	return p.Entity.Type()
}

func (p *Path) EntityType() EntityType {
	return p.Entity.EntityType()
}

func (p *Path) WithThings() *Path {
	p.Entity = &ThingPath{}
	return p
}

func (p *Path) WithThing(thing string) *Path {
	p.Entity = &ThingPath{Thing: thing}
	return p
}

func (p *Path) GetThing() *ThingPath {
	if p.Entity != nil && p.Type() == PathTypeThing {
		return p.Entity.(*ThingPath)
	}
	return nil
}

func (p *Path) WithThingAttributes(thing string) *Path {
	p.Entity = &ThingAttributesPath{ThingPath: ThingPath{Thing: thing}}
	return p
}

func (p *Path) WithThingAttribute(thing, attribute string) *Path {
	p.Entity = &ThingAttributesPath{ThingPath: ThingPath{Thing: thing}, Attribute: attribute}
	return p
}

func (p *Path) GetThingAttributes() *ThingAttributesPath {
	if p.Entity != nil && p.Type() == PathTypeThingAttributes {
		return p.Entity.(*ThingAttributesPath)
	}
	return nil
}

func (p *Path) WithThingMessages(thing string, direction DirectionType, subject string) *Path {
	p.Entity = &ThingMessagesPath{ThingPath: ThingPath{Thing: thing}, Direction: direction, Subject: subject}
	return p
}

func (p *Path) GetThingMessages() *ThingMessagesPath {
	if p.Entity != nil && p.Type() == PathTypeThingMessages {
		return p.Entity.(*ThingMessagesPath)
	}
	return nil
}

func (p *Path) WithThingFeatures(thing string) *Path {
	p.Entity = &ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}}
	return p
}

func (p *Path) WithThingFeature(thing, feature string) *Path {
	p.Entity = &ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}
	return p
}

func (p *Path) GetThingFeatures() *ThingFeaturesPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatures {
		return p.Entity.(*ThingFeaturesPath)
	}
	return nil
}

func (p *Path) WithThingFeatureMessages(thing, feature string, direction DirectionType, subject string) *Path {
	p.Entity = &ThingFeatureMessagesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}, Direction: direction, Subject: subject}
	return p
}

func (p *Path) GetThingFeatureMessages() *ThingFeatureMessagesPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatureMessages {
		return p.Entity.(*ThingFeatureMessagesPath)
	}
	return nil
}

func (p *Path) WithThingFeatureProperties(thing, feature string) *Path {
	p.Entity = &ThingFeaturePropertiesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}, TimeSeries: false}
	return p
}

func (p *Path) WithThingFeaturePropertie(thing, feature, propertie string) *Path {
	p.Entity = &ThingFeaturePropertiesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}, Propertie: propertie, TimeSeries: false}
	return p
}

func (p *Path) WithThingFeaturePropertiesTimeSeries(thing, feature, propertie string) *Path {
	p.Entity = &ThingFeaturePropertiesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}, Propertie: propertie, TimeSeries: true}
	return p
}

func (p *Path) GetThingFeatureProperties() *ThingFeaturePropertiesPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatureProperties {
		return p.Entity.(*ThingFeaturePropertiesPath)
	}
	return nil
}

func (p *Path) WithThingFeatureDesireds(thing, feature string) *Path {
	p.Entity = &ThingFeatureDesiredPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}}
	return p
}

func (p *Path) WithThingFeatureDesired(thing, feature, propertie string) *Path {
	p.Entity = &ThingFeatureDesiredPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}, Propertie: propertie}
	return p
}

func (p *Path) GetThingFeatureDesired() *ThingFeatureDesiredPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatureDesiredProperties {
		return p.Entity.(*ThingFeatureDesiredPath)
	}
	return nil
}

func (p *Path) WithThingFeatureAttributes(thing, feature string) *Path {
	p.Entity = &ThingFeatureAttributesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}}
	return p
}

func (p *Path) WithThingFeatureAttribute(thing, feature, attribute string) *Path {
	p.Entity = &ThingFeatureAttributesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: thing}, Feature: feature}, Attribute: attribute}
	return p
}

func (p *Path) GetThingFeatureAttributes() *ThingFeatureAttributesPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatureAttributes {
		return p.Entity.(*ThingFeatureAttributesPath)
	}
	return nil
}

func (p *Path) WithDevices() *Path {
	p.Entity = &DevicePath{}
	return p
}

func (p *Path) WithDevice(device string) *Path {
	p.Entity = &DevicePath{Device: device}
	return p
}

func (p *Path) GetDevice() *DevicePath {
	if p.Entity != nil && p.Type() == PathTypeDevice {
		return p.Entity.(*DevicePath)
	}
	return nil
}

func (p *Path) WithDeviceStatus(device string) *Path {
	p.Entity = &DeviceStatusPath{DevicePath: DevicePath{Device: device}}
	return p
}

func (p *Path) GetDeviceStatus() *DeviceStatusPath {
	if p.Entity != nil && p.Type() == PathTypeDeviceStatus {
		return p.Entity.(*DeviceStatusPath)
	}
	return nil
}

func (p *Path) WithDeviceAttributes(device string) *Path {
	p.Entity = &DeviceAttributesPath{DevicePath: DevicePath{Device: device}}
	return p
}

func (p *Path) WithDeviceAttribute(device, attribute string) *Path {
	p.Entity = &DeviceAttributesPath{DevicePath: DevicePath{Device: device}, Attribute: attribute}
	return p
}

func (p *Path) GetDeviceAttributes() *DeviceAttributesPath {
	if p.Entity != nil && p.Type() == PathTypeDeviceAttributes {
		return p.Entity.(*DeviceAttributesPath)
	}
	return nil
}

func (p *Path) WithDeviceStrategys(device string) *Path {
	p.Entity = &DeviceStrategysPath{DevicePath: DevicePath{Device: device}}
	return p
}

func (p *Path) WithDeviceStrategy(device, strategy string) *Path {
	p.Entity = &DeviceStrategysPath{DevicePath: DevicePath{Device: device}, Strategy: strategy}
	return p
}

func (p *Path) GetDeviceStrategys() *DeviceStrategysPath {
	if p.Entity != nil && p.Type() == PathTypeDeviceStrategys {
		return p.Entity.(*DeviceStrategysPath)
	}
	return nil
}

func (p *Path) WithDeviceStrategyIndicators(device, strategy string) *Path {
	p.Entity = &DeviceStrategyIndicatorsPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: device}, Strategy: strategy}, TimeSeries: false}
	return p
}

func (p *Path) WithDeviceStrategyIndicator(device, strategy, indicator string) *Path {
	p.Entity = &DeviceStrategyIndicatorsPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: device}, Strategy: strategy}, Indicator: indicator, TimeSeries: false}
	return p
}

func (p *Path) WithDeviceStrategyIndicatorTimeSeries(device, strategy, indicator string) *Path {
	p.Entity = &DeviceStrategyIndicatorsPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: device}, Strategy: strategy}, Indicator: indicator, TimeSeries: true}
	return p
}

func (p *Path) GetDeviceStrategyIndicators() *DeviceStrategyIndicatorsPath {
	if p.Entity != nil && p.Type() == PathTypeDeviceStrategyIndicators {
		return p.Entity.(*DeviceStrategyIndicatorsPath)
	}
	return nil
}

func (p *Path) WithDeviceStrategyAttributes(device, strategy string) *Path {
	p.Entity = &DeviceStrategyAttributesPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: device}, Strategy: strategy}}
	return p
}

func (p *Path) WithDeviceStrategyAttribute(device, strategy, attribute string) *Path {
	p.Entity = &DeviceStrategyAttributesPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: device}, Strategy: strategy}, Attribute: attribute}
	return p
}

func (p *Path) GetDeviceStrategyAttributes() *DeviceStrategyAttributesPath {
	if p.Entity != nil && p.Type() == PathTypeDeviceStrategyAttributes {
		return p.Entity.(*DeviceStrategyAttributesPath)
	}
	return nil
}

func (p *Path) WithDeviceProfiles(device string) *Path {
	p.Entity = &DeviceProfilesPath{DevicePath: DevicePath{Device: device}}
	return p
}

func (p *Path) WithDeviceProfile(device, profile string) *Path {
	p.Entity = &DeviceProfilesPath{DevicePath: DevicePath{Device: device}, Profile: profile}
	return p
}

func (p *Path) GetDeviceProfiles() *DeviceProfilesPath {
	if p.Entity != nil && p.Type() == PathTypeDeviceProfiles {
		return p.Entity.(*DeviceProfilesPath)
	}
	return nil
}

func (p *Path) WithConnections() *Path {
	p.Entity = &ConnectionPath{}
	return p
}

func (p *Path) WithConnection(connection string) *Path {
	p.Entity = &ConnectionPath{Connection: connection}
	return p
}

func (p *Path) GetConnection() *ConnectionPath {
	if p.Entity != nil && p.Type() == PathTypeConnection {
		return p.Entity.(*ConnectionPath)
	}
	return nil
}

func (p *Path) WithConnectionStatus(connection string) *Path {
	p.Entity = &ConnectionStatusPath{ConnectionPath: ConnectionPath{Connection: connection}}
	return p
}

func (p *Path) GetConnectionStatus() *ConnectionStatusPath {
	if p.Entity != nil && p.Type() == PathTypeConnectionStatus {
		return p.Entity.(*ConnectionStatusPath)
	}
	return nil
}

func (p *Path) WithStreams() *Path {
	p.Entity = &StreamPath{}
	return p
}

func (p *Path) WithStream(stream string) *Path {
	p.Entity = &StreamPath{Stream: stream}
	return p
}

func (p *Path) GetStream() *StreamPath {
	if p.Entity != nil && p.Type() == PathTypeStream {
		return p.Entity.(*StreamPath)
	}
	return nil
}

func (p *Path) WithStreamStatus(stream string) *Path {
	p.Entity = &StreamStatusPath{StreamPath: StreamPath{Stream: stream}}
	return p
}

func (p *Path) GetStreamStatus() *StreamStatusPath {
	if p.Entity != nil && p.Type() == PathTypeStreamStatus {
		return p.Entity.(*StreamStatusPath)
	}
	return nil
}

func (p *Path) WithStreamVideos(stream string) *Path {
	p.Entity = &StreamVideosPath{StreamPath: StreamPath{Stream: stream}}
	return p
}

func (p *Path) GetStreamVideos() *StreamVideosPath {
	if p.Entity != nil && p.Type() == PathTypeStreamVideos {
		return p.Entity.(*StreamVideosPath)
	}
	return nil
}

func (p *Path) WithStreamAudios(stream string) *Path {
	p.Entity = &StreamAudiosPath{StreamPath: StreamPath{Stream: stream}}
	return p
}

func (p *Path) GetStreamAudios() *StreamAudiosPath {
	if p.Entity != nil && p.Type() == PathTypeStreamAudios {
		return p.Entity.(*StreamAudiosPath)
	}
	return nil
}

func (p *Path) WithStreamSubscribers(stream string) *Path {
	p.Entity = &StreamSubscribersPath{StreamPath: StreamPath{Stream: stream}}
	return p
}

func (p *Path) GetStreamSubscribers() *StreamSubscribersPath {
	if p.Entity != nil && p.Type() == PathTypeStream_SUBSCRIBERS {
		return p.Entity.(*StreamSubscribersPath)
	}
	return nil
}

func (p *Path) WithFeatures() *Path {
	p.Entity = &FeaturesPath{}
	return p
}

func (p *Path) WithFeature(feature string) *Path {
	p.Entity = &FeaturesPath{Feature: feature}
	return p
}

func (p *Path) GetFeatures() *FeaturesPath {
	if p.Entity != nil && p.Type() == PathTypeFeatures {
		return p.Entity.(*FeaturesPath)
	}
	return nil
}

func (p *Path) WithFeatureProperties(feature string) *Path {
	p.Entity = &FeaturePropertiesPath{FeaturesPath: FeaturesPath{Feature: feature}}
	return p
}

func (p *Path) WithFeaturePropertie(feature, propertie string) *Path {
	p.Entity = &FeaturePropertiesPath{FeaturesPath: FeaturesPath{Feature: feature}, Propertie: propertie}
	return p
}

func (p *Path) GetFeatureProperties() *FeaturePropertiesPath {
	if p.Entity != nil && p.Type() == PathTypeFeaturesProperties {
		return p.Entity.(*FeaturePropertiesPath)
	}
	return nil
}

func (p *Path) WithFeatureDesireds(feature string) *Path {
	p.Entity = &FeatureDesiredPath{FeaturesPath: FeaturesPath{Feature: feature}}
	return p
}

func (p *Path) WithFeatureDesired(feature, propertie string) *Path {
	p.Entity = &FeatureDesiredPath{FeaturesPath: FeaturesPath{Feature: feature}, Propertie: propertie}
	return p
}

func (p *Path) GetFeatureDesired() *FeatureDesiredPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatureDesiredProperties {
		return p.Entity.(*FeatureDesiredPath)
	}
	return nil
}

func (p *Path) WithFeatureAttributes(feature string) *Path {
	p.Entity = &FeatureAttributesPath{FeaturesPath: FeaturesPath{Feature: feature}}
	return p
}

func (p *Path) WithFeatureAttribute(feature, attribute string) *Path {
	p.Entity = &FeatureAttributesPath{FeaturesPath: FeaturesPath{Feature: feature}, Attribute: attribute}
	return p
}

func (p *Path) GetFeatureAttributes() *FeatureAttributesPath {
	if p.Entity != nil && p.Type() == PathTypeThingFeatureAttributes {
		return p.Entity.(*FeatureAttributesPath)
	}
	return nil
}

func (p *Path) WithProperties() *Path {
	p.Entity = &PropertiesPath{}
	return p
}

func (p *Path) WithPropertie(propertie string) *Path {
	p.Entity = &PropertiesPath{Propertie: propertie}
	return p
}

func (p *Path) GetProperties() *PropertiesPath {
	if p.Entity != nil && p.Type() == PathTypeProperties {
		return p.Entity.(*PropertiesPath)
	}
	return nil
}

func (p *Path) WithDesireds() *Path {
	p.Entity = &DesiredPath{}
	return p
}

func (p *Path) WithDesired(propertie string) *Path {
	p.Entity = &DesiredPath{Propertie: propertie}
	return p
}

func (p *Path) GetDesired() *DesiredPath {
	if p.Entity != nil && p.Type() == PathTypeDesiredProperties {
		return p.Entity.(*DesiredPath)
	}
	return nil
}

func (p *Path) WithAttributes() *Path {
	p.Entity = &AttributesPath{}
	return p
}

func (p *Path) WithAttribute(attribute string) *Path {
	p.Entity = &AttributesPath{Attribute: attribute}
	return p
}

func (p *Path) GetAttribute() *AttributesPath {
	if p.Entity != nil && p.Type() == PathTypeAttributes {
		return p.Entity.(*AttributesPath)
	}
	return nil
}

func (p *Path) WithStrategys() *Path {
	p.Entity = &StrategysPath{}
	return p
}

func (p *Path) WithStrategy(strategy string) *Path {
	p.Entity = &StrategysPath{Strategy: strategy}
	return p
}

func (p *Path) GetStrategys() *StrategysPath {
	if p.Entity != nil && p.Type() == PathTypeStrategys {
		return p.Entity.(*StrategysPath)
	}
	return nil
}

func (p *Path) WithStrategyIndicators(strategy string) *Path {
	p.Entity = &StrategyIndicatorsPath{StrategysPath: StrategysPath{Strategy: strategy}}
	return p
}

func (p *Path) WithStrategyIndicator(strategy, indicator string) *Path {
	p.Entity = &StrategyIndicatorsPath{StrategysPath: StrategysPath{Strategy: strategy}, Indicator: indicator}
	return p
}

func (p *Path) GetStrategyIndicators() *StrategyIndicatorsPath {
	if p.Entity != nil && p.Type() == PathTypeStrategyIndicators {
		return p.Entity.(*StrategyIndicatorsPath)
	}
	return nil
}

func (p *Path) WithStrategyAttributes(strategy string) *Path {
	p.Entity = &StrategyAttributesPath{StrategysPath: StrategysPath{Strategy: strategy}}
	return p
}

func (p *Path) WithStrategyAttribute(strategy, attribute string) *Path {
	p.Entity = &StrategyAttributesPath{StrategysPath: StrategysPath{Strategy: strategy}, Attribute: attribute}
	return p
}

func (p *Path) GetStrategyAttributes() *StrategyAttributesPath {
	if p.Entity != nil && p.Type() == PathTypeStrategyAttributes {
		return p.Entity.(*StrategyAttributesPath)
	}
	return nil
}

func (p *Path) WithIndicators(strategy, indicator string) *Path {
	p.Entity = &IndicatorsPath{Indicator: indicator}
	return p
}

func (p *Path) GetIndicators() *IndicatorsPath {
	if p.Entity != nil && p.Type() == PathTypeIndicators {
		return p.Entity.(*IndicatorsPath)
	}
	return nil
}

func (p *Path) WithProfiles() *Path {
	p.Entity = &ProfilesPath{}
	return p
}

func (p *Path) WithProfile(profile string) *Path {
	p.Entity = &ProfilesPath{Profile: profile}
	return p
}

func (p *Path) GetProfile() *ProfilesPath {
	if p.Entity != nil && p.Type() == PathTypeIndicators {
		return p.Entity.(*ProfilesPath)
	}
	return nil
}

func (p *Path) WithStatus(profile string) *Path {
	p.Entity = &StatusPath{}
	return p
}

func (p *Path) GetStatus() *StatusPath {
	if p.Entity != nil && p.Type() == PathTypeStatus {
		return p.Entity.(*StatusPath)
	}
	return nil
}

func (p *Path) WithVideos() *Path {
	p.Entity = &VideosPath{}
	return p
}

func (p *Path) GetVideos() *VideosPath {
	if p.Entity != nil && p.Type() == PathTypeVideos {
		return p.Entity.(*VideosPath)
	}
	return nil
}

func (p *Path) WithAudios() *Path {
	p.Entity = &AudiosPath{}
	return p
}

func (p *Path) GetAudios() *AudiosPath {
	if p.Entity != nil && p.Type() == PathTypeAudios {
		return p.Entity.(*AudiosPath)
	}
	return nil
}

func (p *Path) WithSubscribers() *Path {
	p.Entity = &SubscribersPath{}
	return p
}

func (p *Path) GetSubscribers() *SubscribersPath {
	if p.Entity != nil && p.Type() == PathTypeSubscribers {
		return p.Entity.(*SubscribersPath)
	}
	return nil
}

func GetPathEntityType(str string) EntityType {
	if strings.HasPrefix(str, "@things") {
		return EntityThings
	} else if strings.HasPrefix(str, "@devices") {
		return EntityDevices
	} else if strings.HasPrefix(str, "@connections") {
		return EntityConnections
	} else if strings.HasPrefix(str, "@streams") {
		return EntityStreams
	}
	return EntityUnknown
}

func ValidPath(str string) bool {
	if strings.HasPrefix(str, "@things") {
		return regexThingsPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@devices") {
		return regexDevicesPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@connections") {
		return regexConnectionsPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@streams") {
		return regexStreamsPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@features") {
		return regexFeaturesPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@properties") {
		return regexPropertiesPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@desired") {
		return regexDesiredPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@attributes") {
		return regexAttributesPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@strategys") {
		return regexStrategysPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@indicators") {
		return regexIndicatorsPath.Match([]byte(str))
	} else if strings.HasPrefix(str, "@profiles") {
		return regexProfilesPath.Match([]byte(str))
	} else if str == "@" || str == "@status" || str == "@videos" || str == "@audios" || str == "@subscribers" {
		return true
	}
	return false
}

func parseEntityPath(str string) (EntityPath, error) {
	if strings.HasPrefix(str, "@things") && regexThingsPath.Match([]byte(str)) {
		matches := regexThingsPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "things" {
				if matches[0][24] == "attributes" {
					return &ThingAttributesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Attribute: matches[0][26]}, nil
				}
				if matches[0][4] == "features" {
					if matches[0][6] == "messages" {
						return &ThingFeatureMessagesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][5]}, Direction: DirectionType(matches[0][7]), Subject: matches[0][8]}, nil
					}
					return &ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][5]}, nil
				}
				if matches[0][9] == "features" {
					if matches[0][13] == "properties" {
						timeseries := false
						if matches[0][16] == "timeseries" {
							timeseries = true
						}
						return &ThingFeaturePropertiesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][11]}, Propertie: matches[0][15], TimeSeries: timeseries}, nil
					}
					return &ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][11]}, nil
				}
				if matches[0][17] == "features" {
					if matches[0][21] == "desired" {
						return &ThingFeatureDesiredPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][19]}, Propertie: matches[0][23]}, nil
					} else if matches[0][21] == "properties" {
						return &ThingFeaturePropertiesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][19]}, Propertie: matches[0][23], TimeSeries: false}, nil
					} else if matches[0][21] == "attributes" {
						return &ThingFeatureAttributesPath{ThingFeaturesPath: ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][19]}, Attribute: matches[0][23]}, nil
					}
					return &ThingFeaturesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Feature: matches[0][19]}, nil
				}
				if matches[0][27] == "messages" {
					return &ThingMessagesPath{ThingPath: ThingPath{Thing: matches[0][2]}, Direction: DirectionType(matches[0][28]), Subject: matches[0][29]}, nil
				}
				return &ThingPath{Thing: matches[0][2]}, nil
			}
		}
		return nil, errors.New("error parse things")
	} else if strings.HasPrefix(str, "@devices") && regexDevicesPath.Match([]byte(str)) {
		matches := regexDevicesPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "devices" {
				if matches[0][4] == "status" {
					return &DeviceStatusPath{DevicePath: DevicePath{Device: matches[0][2]}}, nil
				}
				if matches[0][14] == "attributes" {
					return &DeviceAttributesPath{DevicePath: DevicePath{Device: matches[0][2]}, Attribute: matches[0][16]}, nil
				}
				if matches[0][5] == "strategys" {
					if matches[0][9] == "indicators" {
						timeseries := false
						if matches[0][13] == "timeseries" {
							timeseries = true
						}
						return &DeviceStrategyIndicatorsPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: matches[0][2]}, Strategy: matches[0][7]}, Indicator: matches[0][11], TimeSeries: timeseries}, nil
					} else if matches[0][9] == "attributes" {
						return &DeviceStrategyAttributesPath{DeviceStrategysPath: DeviceStrategysPath{DevicePath: DevicePath{Device: matches[0][2]}, Strategy: matches[0][7]}, Attribute: matches[0][11]}, nil
					}
					return &DeviceStrategysPath{DevicePath: DevicePath{Device: matches[0][2]}, Strategy: matches[0][7]}, nil
				}
				if matches[0][17] == "profiles" {
					return &DeviceProfilesPath{DevicePath: DevicePath{Device: matches[0][2]}, Profile: matches[0][19]}, nil
				}
				return &DevicePath{Device: matches[0][2]}, nil
			}
		}
		return nil, errors.New("error parse devices")
	} else if strings.HasPrefix(str, "@connections") && regexConnectionsPath.Match([]byte(str)) {
		matches := regexConnectionsPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "connections" {
				if matches[0][4] == "status" {
					return &ConnectionStatusPath{ConnectionPath: ConnectionPath{Connection: matches[0][2]}}, nil
				}
				return &ConnectionPath{Connection: matches[0][2]}, nil
			}
		}
		return nil, errors.New("error parse connections")
	} else if strings.HasPrefix(str, "@attributes") && regexAttributesPath.Match([]byte(str)) {
		matches := regexAttributesPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "attributes" {
				return &AttributesPath{Attribute: matches[0][3]}, nil
			}
		}
		return nil, errors.New("error parse attributes")
	} else if strings.HasPrefix(str, "@features") && regexFeaturesPath.Match([]byte(str)) {
		matches := regexFeaturesPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "features" {
				if matches[0][5] == "properties" {
					return &FeaturePropertiesPath{FeaturesPath: FeaturesPath{Feature: matches[0][3]}, Propertie: matches[0][7]}, nil
				} else if matches[0][5] == "desired" {
					return &FeatureDesiredPath{FeaturesPath: FeaturesPath{Feature: matches[0][3]}, Propertie: matches[0][7]}, nil
				} else if matches[0][5] == "attributes" {
					return &FeatureAttributesPath{FeaturesPath: FeaturesPath{Feature: matches[0][3]}, Attribute: matches[0][7]}, nil
				} else {
					return &FeaturesPath{Feature: matches[0][3]}, nil
				}
			}
		}
		return nil, errors.New("error parse features")
	} else if strings.HasPrefix(str, "@properties") && regexPropertiesPath.Match([]byte(str)) {
		matches := regexPropertiesPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "properties" {
				return &PropertiesPath{Propertie: matches[0][3]}, nil
			}
		}
		return nil, errors.New("error parse attributes")
	} else if strings.HasPrefix(str, "@desired") && regexDesiredPath.Match([]byte(str)) {
		matches := regexDesiredPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "desired" {
				return &DesiredPath{Propertie: matches[0][3]}, nil
			}
		}
		return nil, errors.New("error parse attributes")
	} else if strings.HasPrefix(str, "@streams") && regexStreamsPath.Match([]byte(str)) {
		matches := regexStreamsPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "streams" {
				if matches[0][4] == "status" {
					return &StreamStatusPath{StreamPath: StreamPath{Stream: matches[0][2]}}, nil
				} else if matches[0][4] == "videos" {
					return &StreamVideosPath{StreamPath: StreamPath{Stream: matches[0][2]}}, nil
				} else if matches[0][4] == "audios" {
					return &StreamAudiosPath{StreamPath: StreamPath{Stream: matches[0][2]}}, nil
				} else if matches[0][4] == "subscribers" {
					return &StreamSubscribersPath{StreamPath: StreamPath{Stream: matches[0][2]}}, nil
				}
				return &StreamPath{Stream: matches[0][2]}, nil
			}
		}
		return nil, errors.New("error parse streams")
	} else if strings.HasPrefix(str, "@strategys") && regexStrategysPath.Match([]byte(str)) {
		matches := regexStrategysPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "strategys" {
				if matches[0][5] == "indicators" {
					return &StrategyIndicatorsPath{StrategysPath: StrategysPath{Strategy: matches[0][3]}, Indicator: matches[0][7]}, nil
				} else if matches[0][5] == "attributes" {
					return &StrategyAttributesPath{StrategysPath: StrategysPath{Strategy: matches[0][3]}, Attribute: matches[0][7]}, nil
				} else {
					return &StrategysPath{Strategy: matches[0][3]}, nil
				}
			}
		}
		return nil, errors.New("error parse strategys")
	} else if strings.HasPrefix(str, "@indicators") && regexIndicatorsPath.Match([]byte(str)) {
		matches := regexIndicatorsPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "indicators" {
				return &IndicatorsPath{Indicator: matches[0][3]}, nil
			}
		}
		return nil, errors.New("error parse profiles")
	} else if strings.HasPrefix(str, "@profiles") && regexProfilesPath.Match([]byte(str)) {
		matches := regexProfilesPath.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if matches[0][1] == "profiles" {
				return &ProfilesPath{Profile: matches[0][3]}, nil
			}
		}
		return nil, errors.New("error parse profiles")
	} else if str == "@status" {
		return &StatusPath{}, nil
	} else if str == "@videos" {
		return &VideosPath{}, nil
	} else if str == "@audios" {
		return &AudiosPath{}, nil
	} else if str == "@subscribers" {
		return &SubscribersPath{}, nil
	} else if str == "@" {
		return &RootPath{}, nil
	}
	return nil, nil
}

func (p *Path) HasWildCard() bool {
	return strings.Contains(p.String(), "*")
}

func WildCardToRegexp(pattern string) string {
	var result strings.Builder
	for i, literal := range strings.Split(pattern, pathWillCard) {

		if i > 0 {
			result.WriteString(".*")
		}

		result.WriteString(regexp.QuoteMeta(literal))
	}
	return result.String()
}

func (p *Path) Match(pattern string) bool {
	result, _ := regexp.MatchString(WildCardToRegexp(pattern), p.String())
	return result
}

func (p *Path) HasPlaceHolders() bool {
	return HasPlaceHolders(p.String())
}

type EntityPath interface {
	Name() string
	String() string
	Type() PathType
	EntityType() EntityType
	IsParent(target *Path) bool
}

type RootPath struct {
}

func (t *RootPath) Type() PathType {
	return PathTypeRoot
}

func (t *RootPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *RootPath) String() string {
	return pathRoot
}

func (t *RootPath) Name() string {
	return pathRoot
}

func (t *RootPath) IsParent(target *Path) bool {
	return false
}

type ThingPath struct {
	Thing string
}

func (t *ThingPath) Type() PathType {
	return PathTypeThing
}

func (t *ThingPath) EntityType() EntityType {
	return EntityThings
}

func (t *ThingPath) ThingName() string {
	if t.Thing != "" {
		return t.Thing
	}
	return "things"
}

func (t *ThingPath) Name() string {
	if t.Thing != "" {
		return t.Thing
	}
	return "things"
}

func (t *ThingPath) String() string {
	if t.Thing != "" {
		return fmt.Sprintf(pathThingFormat, t.Thing)
	}
	return pathThings
}

func (t *ThingPath) IsParent(target *Path) bool {
	return false
}

type ThingAttributesPath struct {
	ThingPath
	Attribute string
}

func (t *ThingAttributesPath) Type() PathType {
	return PathTypeThingAttributes
}

func (t *ThingAttributesPath) String() string {
	if t.Thing != "" {
		if t.Attribute != "" {
			return fmt.Sprintf(pathThingAttributeFormat, t.Thing, t.Attribute)
		}
		return fmt.Sprintf(pathThingAttributes, t.Thing)
	}
	return pathThings
}

func (t *ThingAttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *ThingAttributesPath) IsParent(target *Path) bool {
	return false
}

func (t *ThingAttributesPath) GetAttribute() *AttributesPath {
	return &AttributesPath{Attribute: t.Attribute}
}

type ThingMessagesPath struct {
	ThingPath
	Direction DirectionType
	Subject   string
}

func (t *ThingMessagesPath) Type() PathType {
	return PathTypeThingMessages
}

func (t *ThingMessagesPath) String() string {
	if t.Thing != "" {
		if t.Direction != "" && t.Subject != "" {
			return fmt.Sprintf(pathMessagesFormat, t.ThingPath.String(), t.Direction, t.Subject)
		}
	}
	return pathThings
}

func (t *ThingMessagesPath) Name() string {
	return "messages"
}

func (t *ThingMessagesPath) IsParent(target *Path) bool {
	return false
}

type ThingFeaturesPath struct {
	ThingPath
	Feature string
}

func (t *ThingFeaturesPath) Type() PathType {
	return PathTypeThingFeatures
}

func (t *ThingFeaturesPath) String() string {
	if t.Thing != "" {
		if t.Feature != "" {
			return fmt.Sprintf(pathThingFeatureFormat, t.Thing, t.Feature)
		}
		return fmt.Sprintf(pathThingFeatures, t.Thing)
	}
	return pathThings
}

func (t *ThingFeaturesPath) Name() string {
	if t.Feature != "" {
		return t.Feature
	}
	return "features"
}

func (t *ThingFeaturesPath) IsParent(target *Path) bool {
	return false
}

func (t *ThingFeaturesPath) GetFeatures() *FeaturesPath {
	return &FeaturesPath{Feature: t.Feature}
}

type ThingFeatureMessagesPath struct {
	ThingFeaturesPath
	Direction DirectionType
	Subject   string
}

func (t *ThingFeatureMessagesPath) Type() PathType {
	return PathTypeThingFeatureMessages
}

func (t *ThingFeatureMessagesPath) String() string {
	if t.Thing != "" && t.Feature != "" && t.Direction != "" && t.Subject != "" {
		return fmt.Sprintf(pathMessagesFormat, t.ThingFeaturesPath.String(), t.Direction, t.Subject)
	}
	return pathThings
}

func (t *ThingFeatureMessagesPath) Name() string {
	return "messages"
}

func (t *ThingFeatureMessagesPath) IsParent(target *Path) bool {
	return false
}

type ThingFeaturePropertiesPath struct {
	ThingFeaturesPath
	Propertie  string
	TimeSeries bool
}

func (t *ThingFeaturePropertiesPath) Type() PathType {
	return PathTypeThingFeatureProperties
}

func (t *ThingFeaturePropertiesPath) String() string {
	if t.Thing != "" {
		if t.Feature != "" {
			if t.Propertie != "" {
				if t.TimeSeries {
					return fmt.Sprintf(pathThingFeaturePropertyTimeseries, t.Thing, t.Feature, t.Propertie)
				}
				return fmt.Sprintf(pathThingFeaturePropertyFormat, t.Thing, t.Feature, t.Propertie)
			}
			return fmt.Sprintf(pathThingFeaturePropertiesFormat, t.Thing, t.Feature)
		}
		return fmt.Sprintf(pathThingFeatures, t.Thing)
	}
	return pathThings
}

func (t *ThingFeaturePropertiesPath) Name() string {
	if t.Propertie != "" {
		return t.Propertie
	}
	return "properties"
}

func (t *ThingFeaturePropertiesPath) IsParent(target *Path) bool {
	return false
}

func (t *ThingFeaturePropertiesPath) GetProperties() *PropertiesPath {
	return &PropertiesPath{Propertie: t.Propertie}
}

func (t *ThingFeaturePropertiesPath) GetFeatureProperties() *FeaturePropertiesPath {
	return &FeaturePropertiesPath{FeaturesPath: FeaturesPath{Feature: t.Feature}, Propertie: t.Propertie}
}

type ThingFeatureDesiredPath struct {
	ThingFeaturesPath
	Propertie string
}

func (t *ThingFeatureDesiredPath) Type() PathType {
	return PathTypeThingFeatureDesiredProperties
}

func (t *ThingFeatureDesiredPath) String() string {
	if t.Thing != "" {
		if t.Feature != "" {
			if t.Propertie != "" {
				return fmt.Sprintf(pathThingFeatureDesiredPropertyFormat, t.Thing, t.Feature, t.Propertie)
			}
			return fmt.Sprintf(pathThingFeatureDesiredPropertiesFormat, t.Thing, t.Feature)
		}
		return fmt.Sprintf(pathThingFeatures, t.Thing)
	}
	return pathThings
}

func (t *ThingFeatureDesiredPath) Name() string {
	if t.Propertie != "" {
		return t.Propertie
	}
	return "desired"
}

func (t *ThingFeatureDesiredPath) IsParent(target *Path) bool {
	return false
}

func (t *ThingFeatureDesiredPath) GetDesired() *DesiredPath {
	return &DesiredPath{Propertie: t.Propertie}
}

func (t *ThingFeatureDesiredPath) GetFeatureDesired() *FeatureDesiredPath {
	return &FeatureDesiredPath{FeaturesPath: FeaturesPath{Feature: t.Feature}, Propertie: t.Propertie}
}

type ThingFeatureAttributesPath struct {
	ThingFeaturesPath
	Attribute string
}

func (t *ThingFeatureAttributesPath) Type() PathType {
	return PathTypeThingFeatureAttributes
}

func (t *ThingFeatureAttributesPath) String() string {
	if t.Thing != "" {
		if t.Feature != "" {
			if t.Attribute != "" {
				return fmt.Sprintf(pathThingFeatureAttributeFormat, t.Thing, t.Feature, t.Attribute)
			}
			return fmt.Sprintf(pathThingFeatureAttributes, t.Thing, t.Feature)
		}
		return fmt.Sprintf(pathThingFeatures, t.Thing)
	}
	return pathThings
}

func (t *ThingFeatureAttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *ThingFeatureAttributesPath) IsParent(target *Path) bool {
	return false
}

func (t *ThingFeatureAttributesPath) GetAttribute() *AttributesPath {
	return &AttributesPath{Attribute: t.Attribute}
}

func (t *ThingFeatureAttributesPath) GetFeatureAttributes() *FeatureAttributesPath {
	return &FeatureAttributesPath{FeaturesPath: FeaturesPath{Feature: t.Feature}, Attribute: t.Attribute}
}

type DevicePath struct {
	Device string
}

func (t *DevicePath) Type() PathType {
	return PathTypeDevice
}

func (t *DevicePath) EntityType() EntityType {
	return EntityDevices
}

func (t *DevicePath) String() string {
	if t.Device != "" {
		return fmt.Sprintf(pathDeviceFormat, t.Device)
	}
	return pathDevices
}

func (t *DevicePath) Name() string {
	if t.Device != "" {
		return t.Device
	}
	return "devices"
}

func (t *DevicePath) IsParent(target *Path) bool {
	return false
}

type DeviceStatusPath struct {
	DevicePath
}

func (t *DeviceStatusPath) Type() PathType {
	return PathTypeDeviceStatus
}

func (t *DeviceStatusPath) String() string {
	if t.Device != "" {
		return fmt.Sprintf(pathDeviceHealthStatus, t.Device)
	}
	return pathDevices
}

func (t *DeviceStatusPath) Name() string {
	return "status"
}

func (t *DeviceStatusPath) IsParent(target *Path) bool {
	return false
}

func (t *DeviceStatusPath) GetStatus() *StatusPath {
	return &StatusPath{}
}

type DeviceAttributesPath struct {
	DevicePath
	Attribute string
}

func (t *DeviceAttributesPath) Type() PathType {
	return PathTypeDeviceAttributes
}

func (t *DeviceAttributesPath) String() string {
	if t.Device != "" {
		if t.Attribute != "" {
			return fmt.Sprintf(pathDevicePropertyFormat, t.Device, t.Attribute)
		}
		return fmt.Sprintf(pathDeviceProperties, t.Device)
	}
	return pathDevices
}

func (t *DeviceAttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *DeviceAttributesPath) IsParent(target *Path) bool {
	return false
}

func (t *DeviceAttributesPath) GetAttribute() *AttributesPath {
	return &AttributesPath{Attribute: t.Attribute}
}

type DeviceStrategysPath struct {
	DevicePath
	Strategy string
}

func (t *DeviceStrategysPath) Type() PathType {
	return PathTypeDeviceStrategys
}

func (t *DeviceStrategysPath) String() string {
	if t.Device != "" {
		if t.Strategy != "" {
			return fmt.Sprintf(pathDeviceStrategysFormat, t.Device, t.Strategy)
		}
		return fmt.Sprintf(pathDeviceStrategys, t.Device)
	}
	return pathDevices
}

func (t *DeviceStrategysPath) Name() string {
	if t.Strategy != "" {
		return t.Strategy
	}
	return "strategys"
}

func (t *DeviceStrategysPath) IsParent(target *Path) bool {
	return false
}

func (t *DeviceStrategysPath) GetStrategys() *StrategysPath {
	return &StrategysPath{Strategy: t.Strategy}
}

type DeviceStrategyIndicatorsPath struct {
	DeviceStrategysPath
	Indicator  string
	TimeSeries bool
}

func (t *DeviceStrategyIndicatorsPath) Type() PathType {
	return PathTypeDeviceStrategyIndicators
}

func (t *DeviceStrategyIndicatorsPath) String() string {
	if t.Device != "" {
		if t.Strategy != "" {
			if t.Indicator != "" {
				if t.TimeSeries {
					return fmt.Sprintf(pathDeviceIndicatorTimeseries, t.Device, t.Strategy, t.Indicator)
				}
				return fmt.Sprintf(pathDeviceIndicatorFormat, t.Device, t.Strategy, t.Indicator)
			}
			return fmt.Sprintf(pathDeviceStrategysFormat, t.Device, t.Strategy)
		}
		return fmt.Sprintf(pathDeviceStrategys, t.Device)
	}
	return pathDevices
}

func (t *DeviceStrategyIndicatorsPath) Name() string {
	if t.Indicator != "" {
		return t.Indicator
	}
	return "indicators"
}

func (t *DeviceStrategyIndicatorsPath) IsParent(target *Path) bool {
	return false
}

func (t *DeviceStrategyIndicatorsPath) GetIndicators() *IndicatorsPath {
	return &IndicatorsPath{Indicator: t.Indicator}
}

func (t *DeviceStrategyIndicatorsPath) GetStrategyIndicators() *StrategyIndicatorsPath {
	return &StrategyIndicatorsPath{StrategysPath: StrategysPath{Strategy: t.Strategy}, Indicator: t.Indicator}
}

type DeviceStrategyAttributesPath struct {
	DeviceStrategysPath
	Attribute string
}

func (t *DeviceStrategyAttributesPath) Type() PathType {
	return PathTypeDeviceStrategyAttributes
}

func (t *DeviceStrategyAttributesPath) String() string {
	if t.Device != "" {
		if t.Strategy != "" {
			if t.Attribute != "" {
				return fmt.Sprintf(pathDeviceStrategysAttributeFormat, t.Device, t.Strategy, t.Attribute)
			}
			return fmt.Sprintf(pathDeviceStrategysAttributes, t.Device, t.Strategy)
		}
		return fmt.Sprintf(pathDeviceStrategys, t.Device)
	}
	return pathDevices
}

func (t *DeviceStrategyAttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *DeviceStrategyAttributesPath) IsParent(target *Path) bool {
	return false
}

func (t *DeviceStrategyAttributesPath) GetStrategyAttributes() *StrategyAttributesPath {
	return &StrategyAttributesPath{StrategysPath: StrategysPath{Strategy: t.Strategy}, Attribute: t.Attribute}
}

type DeviceProfilesPath struct {
	DevicePath
	Profile string
}

func (t *DeviceProfilesPath) Type() PathType {
	return PathTypeDeviceProfiles
}

func (t *DeviceProfilesPath) String() string {
	if t.Device != "" {
		if t.Profile != "" {
			return fmt.Sprintf(pathDeviceProfileFormat, t.Device, t.Profile)
		}
		return fmt.Sprintf(pathDeviceProfiles, t.Device)
	}
	return pathDevices
}

func (t *DeviceProfilesPath) Name() string {
	if t.Profile != "" {
		return t.Profile
	}
	return "profiles"
}

func (t *DeviceProfilesPath) IsParent(target *Path) bool {
	return false
}

func (t *DeviceProfilesPath) GetProfile() *ProfilesPath {
	return &ProfilesPath{Profile: t.Profile}
}

type ConnectionPath struct {
	Connection string
}

func (t *ConnectionPath) Type() PathType {
	return PathTypeConnection
}

func (t *ConnectionPath) EntityType() EntityType {
	return EntityConnections
}

func (t *ConnectionPath) String() string {
	if t.Connection != "" {
		return fmt.Sprintf(pathConnectionFormat, t.Connection)
	}
	return pathConnections
}

func (t *ConnectionPath) Name() string {
	if t.Connection != "" {
		return t.Connection
	}
	return "connections"
}

func (t *ConnectionPath) IsParent(target *Path) bool {
	return false
}

type ConnectionStatusPath struct {
	ConnectionPath
}

func (t *ConnectionStatusPath) Type() PathType {
	return PathTypeConnectionStatus
}

func (t *ConnectionStatusPath) String() string {
	if t.Connection != "" {
		return fmt.Sprintf(pathConnectionStatus, t.Connection)
	}
	return pathConnections
}

func (t *ConnectionStatusPath) Name() string {
	return "status"
}

func (t *ConnectionStatusPath) IsParent(target *Path) bool {
	return false
}

type StreamPath struct {
	Stream string
}

func (t *StreamPath) Type() PathType {
	return PathTypeStream
}

func (t *StreamPath) EntityType() EntityType {
	return EntityStreams
}

func (t *StreamPath) String() string {
	if t.Stream != "" {
		return fmt.Sprintf(pathStreamFormat, t.Stream)
	}
	return pathStreams
}

func (t *StreamPath) Name() string {
	if t.Stream != "" {
		return t.Stream
	}
	return "streams"
}

func (t *StreamPath) IsParent(target *Path) bool {
	return false
}

type StreamStatusPath struct {
	StreamPath
}

func (t *StreamStatusPath) Type() PathType {
	return PathTypeStreamStatus
}

func (t *StreamStatusPath) String() string {
	if t.Stream != "" {
		return fmt.Sprintf(pathStreamStatus, t.Stream)
	}
	return pathStreams
}

func (t *StreamStatusPath) Name() string {
	return "status"
}

func (t *StreamStatusPath) IsParent(target *Path) bool {
	return false
}

type StreamVideosPath struct {
	StreamPath
}

func (t *StreamVideosPath) Type() PathType {
	return PathTypeStreamVideos
}

func (t *StreamVideosPath) String() string {
	if t.Stream != "" {
		return fmt.Sprintf(pathStreamVideos, t.Stream)
	}
	return pathStreams
}

func (t *StreamVideosPath) Name() string {
	return "videos"
}

func (t *StreamVideosPath) IsParent(target *Path) bool {
	return false
}

type StreamAudiosPath struct {
	StreamPath
}

func (t *StreamAudiosPath) Type() PathType {
	return PathTypeStreamAudios
}

func (t *StreamAudiosPath) String() string {
	if t.Stream != "" {
		return fmt.Sprintf(pathStreamAudios, t.Stream)
	}
	return pathStreams
}

func (t *StreamAudiosPath) Name() string {
	return "audios"
}

func (t *StreamAudiosPath) IsParent(target *Path) bool {
	return false
}

type StreamSubscribersPath struct {
	StreamPath
}

func (t *StreamSubscribersPath) Type() PathType {
	return PathTypeStream_SUBSCRIBERS
}

func (t *StreamSubscribersPath) String() string {
	if t.Stream != "" {
		return fmt.Sprintf(pathConnectionSubscribers, t.Stream)
	}
	return pathStreams
}

func (t *StreamSubscribersPath) Name() string {
	return "subscribers"
}

func (t *StreamSubscribersPath) IsParent(target *Path) bool {
	return false
}

type FeaturesPath struct {
	Feature string
}

func (t *FeaturesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *FeaturesPath) Type() PathType {
	return PathTypeFeatures
}

func (t *FeaturesPath) String() string {
	if t.Feature != "" {
		return fmt.Sprintf(pathFeatureFormat, t.Feature)
	}
	return fmt.Sprintf(pathFeatures)
}

func (t *FeaturesPath) Name() string {
	if t.Feature != "" {
		return t.Feature
	}
	return "features"
}

func (t *FeaturesPath) IsParent(target *Path) bool {
	return false
}

type FeaturePropertiesPath struct {
	FeaturesPath
	Propertie string
}

func (t *FeaturePropertiesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *FeaturePropertiesPath) Type() PathType {
	return PathTypeFeaturesProperties
}

func (t *FeaturePropertiesPath) String() string {
	if t.Feature != "" {
		if t.Propertie != "" {
			return fmt.Sprintf(pathFeaturePropertyFormat, t.Feature, t.Propertie)
		}
		return fmt.Sprintf(pathFeaturePropertiesFormat, t.Feature)
	}
	return fmt.Sprintf(pathFeatures)
}

func (t *FeaturePropertiesPath) Name() string {
	if t.Propertie != "" {
		return t.Propertie
	}
	return "properties"
}

func (t *FeaturePropertiesPath) IsParent(target *Path) bool {
	return false
}

func (t *FeaturePropertiesPath) GetProperties() *PropertiesPath {
	return &PropertiesPath{Propertie: t.Propertie}
}

type FeatureDesiredPath struct {
	FeaturesPath
	Propertie string
}

func (t *FeatureDesiredPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *FeatureDesiredPath) Type() PathType {
	return PathTypeFeaturesDesiredProperties
}

func (t *FeatureDesiredPath) String() string {
	if t.Feature != "" {
		if t.Propertie != "" {
			return fmt.Sprintf(pathFeatureDesiredPropertyFormat, t.Feature, t.Propertie)
		}
		return fmt.Sprintf(pathFeatureDesiredPropertiesFormat, t.Feature)
	}
	return fmt.Sprintf(pathFeatures)
}

func (t *FeatureDesiredPath) Name() string {
	if t.Propertie != "" {
		return t.Propertie
	}
	return "desired"
}

func (t *FeatureDesiredPath) IsParent(target *Path) bool {
	return false
}

func (t *FeatureDesiredPath) GetDesired() *DesiredPath {
	return &DesiredPath{Propertie: t.Propertie}
}

type FeatureAttributesPath struct {
	FeaturesPath
	Attribute string
}

func (t *FeatureAttributesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *FeatureAttributesPath) Type() PathType {
	return PathTypeFeaturesAttributes
}

func (t *FeatureAttributesPath) String() string {
	if t.Feature != "" {
		if t.Attribute != "" {
			return fmt.Sprintf(pathFeatureAttributeFormat, t.Feature, t.Attribute)
		}
		return fmt.Sprintf(pathFeatureAttributes, t.Feature)
	}
	return fmt.Sprintf(pathFeatures)
}

func (t *FeatureAttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *FeatureAttributesPath) IsParent(target *Path) bool {
	return false
}

func (t *FeatureAttributesPath) GetAttribute() *AttributesPath {
	return &AttributesPath{Attribute: t.Attribute}
}

type PropertiesPath struct {
	Propertie string
}

func (t *PropertiesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *PropertiesPath) Type() PathType {
	return PathTypeProperties
}

func (t *PropertiesPath) String() string {
	if t.Propertie != "" {
		return fmt.Sprintf(pathPropertieFormat, t.Propertie)
	}
	return fmt.Sprintf(pathProperties)
}

func (t *PropertiesPath) Name() string {
	if t.Propertie != "" {
		return t.Propertie
	}
	return "properties"
}

func (t *PropertiesPath) IsParent(target *Path) bool {
	return false
}

type DesiredPath struct {
	Propertie string
}

func (t *DesiredPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *DesiredPath) Type() PathType {
	return PathTypeDesiredProperties
}

func (t *DesiredPath) String() string {
	if t.Propertie != "" {
		return fmt.Sprintf(pathDesiredFormat, t.Propertie)
	}
	return fmt.Sprintf(pathDesireds)
}

func (t *DesiredPath) Name() string {
	if t.Propertie != "" {
		return t.Propertie
	}
	return "desired"
}

func (t *DesiredPath) IsParent(target *Path) bool {
	return false
}

type AttributesPath struct {
	Attribute string
}

func (t *AttributesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *AttributesPath) Type() PathType {
	return PathTypeAttributes
}

func (t *AttributesPath) String() string {
	if t.Attribute != "" {
		return fmt.Sprintf(pathAttributeFormat, t.Attribute)
	}
	return pathAttributes
}

func (t *AttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *AttributesPath) IsParent(target *Path) bool {
	return false
}

type StrategysPath struct {
	Strategy string
}

func (t *StrategysPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *StrategysPath) Type() PathType {
	return PathTypeStrategys
}

func (t *StrategysPath) String() string {
	if t.Strategy != "" {
		return fmt.Sprintf(pathStrategyFormat, t.Strategy)
	}
	return fmt.Sprintf(pathStrategys)
}

func (t *StrategysPath) Name() string {
	if t.Strategy != "" {
		return t.Strategy
	}
	return "strategys"
}

func (t *StrategysPath) IsParent(target *Path) bool {
	return false
}

type StrategyIndicatorsPath struct {
	StrategysPath
	Indicator string
}

func (t *StrategyIndicatorsPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *StrategyIndicatorsPath) Type() PathType {
	return PathTypeStrategyIndicators
}

func (t *StrategyIndicatorsPath) String() string {
	if t.Strategy != "" {
		if t.Indicator != "" {
			return fmt.Sprintf(pathStrategysIndicatorFormat, t.Strategy, t.Indicator)
		}
		return fmt.Sprintf(pathStrategysIndicators, t.Strategy)
	}
	return pathStrategys
}

func (t *StrategyIndicatorsPath) Name() string {
	if t.Indicator != "" {
		return t.Indicator
	}
	return "indicators"
}

func (t *StrategyIndicatorsPath) IsParent(target *Path) bool {
	return false
}

func (t *StrategyIndicatorsPath) GetIndicators() *IndicatorsPath {
	return &IndicatorsPath{Indicator: t.Indicator}
}

type StrategyAttributesPath struct {
	StrategysPath
	Attribute string
}

func (t *StrategyAttributesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *StrategyAttributesPath) Type() PathType {
	return PathTypeStrategyAttributes
}

func (t *StrategyAttributesPath) String() string {
	if t.Strategy != "" {
		if t.Attribute != "" {
			return fmt.Sprintf(pathStrategysAttributeFormat, t.Strategy, t.Attribute)
		}
		return fmt.Sprintf(pathStrategysAttributes, t.Strategy)
	}
	return pathStrategys
}

func (t *StrategyAttributesPath) Name() string {
	if t.Attribute != "" {
		return t.Attribute
	}
	return "attributes"
}

func (t *StrategyAttributesPath) IsParent(target *Path) bool {
	return false
}

func (t *StrategyAttributesPath) GetAttribute() *AttributesPath {
	return &AttributesPath{Attribute: t.Attribute}
}

type IndicatorsPath struct {
	Indicator string
}

func (t *IndicatorsPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *IndicatorsPath) Type() PathType {
	return PathTypeIndicators
}

func (t *IndicatorsPath) String() string {
	if t.Indicator != "" {
		return fmt.Sprintf(pathIndicatorFormat, t.Indicator)
	}
	return pathIndicators
}

func (t *IndicatorsPath) Name() string {
	if t.Indicator != "" {
		return t.Indicator
	}
	return "indicators"
}

func (t *IndicatorsPath) IsParent(target *Path) bool {
	return false
}

type ProfilesPath struct {
	Profile string
}

func (t *ProfilesPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *ProfilesPath) Type() PathType {
	return PathTypeProfiles
}

func (t *ProfilesPath) String() string {
	if t.Profile != "" {
		return fmt.Sprintf(pathProfileFormat, t.Profile)
	}
	return pathProfiles
}

func (t *ProfilesPath) Name() string {
	if t.Profile != "" {
		return t.Profile
	}
	return "profiles"
}

func (t *ProfilesPath) IsParent(target *Path) bool {
	return false
}

type StatusPath struct {
}

func (t *StatusPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *StatusPath) Type() PathType {
	return PathTypeStatus
}

func (t *StatusPath) String() string {
	return pathStatus
}

func (t *StatusPath) Name() string {
	return "status"
}

func (t *StatusPath) IsParent(target *Path) bool {
	return false
}

type VideosPath struct {
}

func (t *VideosPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *VideosPath) Type() PathType {
	return PathTypeVideos
}

func (t *VideosPath) String() string {
	return pathVideos
}

func (t *VideosPath) Name() string {
	return "videos"
}

func (t *VideosPath) IsParent(target *Path) bool {
	return false
}

type AudiosPath struct {
}

func (t *AudiosPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *AudiosPath) Type() PathType {
	return PathTypeAudios
}

func (t *AudiosPath) String() string {
	return pathAudios
}

func (t *AudiosPath) Name() string {
	return "audios"
}

func (t *AudiosPath) IsParent(target *Path) bool {
	return false
}

type SubscribersPath struct {
}

func (t *SubscribersPath) EntityType() EntityType {
	return EntityUnknown
}

func (t *SubscribersPath) Type() PathType {
	return PathTypeSubscribers
}

func (t *SubscribersPath) String() string {
	return pathSubscribers
}

func (t *SubscribersPath) Name() string {
	return "subscribers"
}

func (t *SubscribersPath) IsParent(target *Path) bool {
	return false
}

type PathList []Path

func (e PathList) Value() (driver.Value, error) {
	j, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return string(j), err
}

func (e *PathList) Scan(value interface{}) error {
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

func isJoinOf(p, s PathType) bool {
	switch p {
	case PathTypeRoot:
		switch s {
		case PathTypeThing, PathTypeDevice, PathTypeStream, PathTypeConnection:
			return true
		}
	case PathTypeThing:
		switch s {
		case PathTypeFeatures, PathTypeFeaturesProperties, PathTypeFeaturesDesiredProperties, PathTypeFeaturesAttributes, PathTypeAttributes:
			return true
		}
	case PathTypeThingFeatures:
		switch s {
		case PathTypeProperties, PathTypeDesiredProperties, PathTypeAttributes:
			return true
		}
	case PathTypeDevice:
		switch s {
		case PathTypeStrategys, PathTypeStrategyIndicators, PathTypeStrategyAttributes, PathTypeAttributes, PathTypeProfiles, PathTypeStatus:
			return true
		}
	case PathTypeDeviceStrategys:
		switch s {
		case PathTypeIndicators, PathTypeAttributes:
			return true
		}
	case PathTypeConnection:
		switch s {
		case PathTypeStatus:
			return true
		}
	case PathTypeStream:
		switch s {
		case PathTypeStatus, PathTypeVideos, PathTypeAudios, PathTypeSubscribers:
			return true
		}
	case PathTypeFeatures:
		switch s {
		case PathTypeAttributes, PathTypeProperties, PathTypeDesiredProperties:
			return true
		}
	case PathTypeStrategys:
		switch s {
		case PathTypeIndicators, PathTypeAttributes:
			return true
		}
	}
	return false
}

func (p *Path) IsParent(target *Path) bool {
	if p.Entity == nil || target.Entity == nil {
		return false
	}

	return p.Entity.IsParent(target)
}

func (p *Path) RelativeOf(target *Path) (*Path, error) {
	return nil, nil
}

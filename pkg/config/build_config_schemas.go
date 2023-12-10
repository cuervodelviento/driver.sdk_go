package config

type ErrConfigExpectedNotCompatible struct {
}

func (e *ErrConfigExpectedNotCompatible) Error() string {
	return "Config expected not compatible"
}

// type Converter interface {
// 	AsGetChannels() (netsocs_config.GetChannelResponse, error)
// }

// func Convert()

package config

import (
	"encoding/json"
)

// AddConfigHandler adds a handler for a specific config key

type HandlerValue struct {
	DeviceData *ConfigMessageDeviceData
	Value      string
}

var handlersMap map[NetsocsConfigKey]handlerFunction

type handlerFunction func(valueMessage string, deviceData *ConfigMessageDeviceData) (string, error)

// type tConfigHandler[RESPONSE_TEMPLATE any, VALUE_EXPECTED_TEMPLATE any] func(valueMessage HandlerValue) (RESPONSE_TEMPLATE, error)

type FuncConfigHandler func(valueMessage HandlerValue) (interface{}, error)

func AddConfigHandler(configKey NetsocsConfigKey, configHandler FuncConfigHandler) error {
	if handlersMap == nil {
		handlersMap = make(map[NetsocsConfigKey]handlerFunction)
	}
	handlersMap[configKey] = func(valueMessage string, deviceData *ConfigMessageDeviceData) (string, error) {

		response, err := configHandler(HandlerValue{
			DeviceData: deviceData,
			Value:      valueMessage,
		})
		if err != nil {
			return "", err
		}
		jsondata, err := json.Marshal(response)
		return string(jsondata), err
	}
	return nil

}

// Thanks https://stackoverflow.com/a/71955439
func UnmarshalAny[T any](bytes []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, err
	}
	return out, nil
}

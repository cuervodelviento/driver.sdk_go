package event

import (
	"bytes"
	"fmt"
	"net/http"
)

type EventDispatcher struct {
	host string
	port int
	key  string
}

const DEFAULT_EVENT_PORT = 3070

func NewEventDispatcher(host string, key string) *EventDispatcher {
	return &EventDispatcher{
		host: host,
		port: DEFAULT_EVENT_PORT,
		key:  key,
	}
}

func (e *EventDispatcher) SetPort(port int) {
	e.port = port
}

func (e *EventDispatcher) Dispatch(eventKey string, deviceID int) error {
	return e.dispatch(eventKey, deviceID)
}

func (e *EventDispatcher) DispatchWithChannels(eventKey string, deviceID int, channelNumber int) error {
	return nil
}

func (e *EventDispatcher) dispatch(eventKey string, deviceId int) error {
	bodyString := fmt.Sprintf(`{"deviceId": %d}`, 1)
	body := []byte(bodyString)
	bodyReader := bytes.NewReader(body)

	_, err := http.Post(fmt.Sprintf("http://%s:%d/%s", e.host, e.port, eventKey), "application/json", bodyReader)
	return err
}

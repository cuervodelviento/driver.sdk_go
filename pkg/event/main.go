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
	bodyString := fmt.Sprintf(`{"deviceId": %d, }`, deviceID)
	body := []byte(bodyString)
	bodyReader := bytes.NewReader(body)

	_, err := http.Post(fmt.Sprintf("http://%s:%d/v1/topologia/misc/%s", e.host, e.port, eventKey), "application/json", bodyReader)
	return err
}

func (e *EventDispatcher) DispatchWithChannels(eventKey string, deviceID int, channelNumber int) error {
	bodyString := fmt.Sprintf(`{"deviceId": %d, "deviceByProps": [ "\"parentDevice\": %d", "\"channelNumber\": %d"]}`, deviceID, deviceID, channelNumber)
	body := []byte(bodyString)
	bodyReader := bytes.NewReader(body)
	_, err := http.Post(fmt.Sprintf("http://%s:%d/v1/topologia/misc/%s", e.host, e.port, eventKey), "application/json",
		bodyReader)
	// respString, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(respString))
	return err
}

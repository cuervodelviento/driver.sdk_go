package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type ConfigMessagePort interface {
	GetDeviceData() (interface{}, error)
}

type ConfigMessage struct {
	DeviceData *ConfigMessageDeviceData `json:"deviceData"`
	ConfigKey  NetsocsConfigKey         `json:"configKey"`
	Value      string                   `json:"value"`
	RequestID  string                   `json:"requestId"`
	rawMessage []byte
}

func (c *ConfigMessage) GetDeviceData() (interface{}, error) {
	return nil, nil
}

func (c *ConfigMessage) GetConfigKey() NetsocsConfigKey {
	return c.ConfigKey
}

func (c *ConfigMessage) GetRawMessage() []byte {
	return c.rawMessage
}

type ConfigMessageDeviceData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	IsSSL    bool   `json:"is_ssl"`
	SSLPort  int    `json:"ssl_port"`
}

var messages = make(chan *ConfigMessage)
var responses = make(chan *s_response)

type s_response struct {
	RequestId string `json:"requestId"`
	Data      string `json:"data"`
}

func ListenConfig(host string) {

	go func() {
		for {
			select {
			case message := <-messages:
				handler := handlersMap[message.ConfigKey]
				if handler != nil {
					response, err := handler(message.Value, message.DeviceData)
					if err == nil {
						responses <- &s_response{
							RequestId: message.RequestID,
							Data:      response,
						}
					} else {
						fmt.Println("Error in handler:", err)
					}
				}
			}
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: host, Path: "/ws/v1/config_communication"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), http.Header{
		"Authorization": []string{"sasas"},
	})
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			configMessage := &ConfigMessage{}
			err = json.Unmarshal(message, configMessage)
			if err != nil {
				log.Println("unmarshal:", err)
			} else {
				messages <- configMessage
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case response := <-responses:
			jsondata, err := json.Marshal(response)
			if err != nil {
				log.Println("marshal:", err)
			} else {
				err = c.WriteMessage(websocket.TextMessage, jsondata)
				if err != nil {
					log.Println("write:", err)
					return
				}
			}

		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			}
			return
		}
	}

}

/*
The event loader is a way for the client can send events available by the integration to DriverHub.
*/
package client

import (
	"encoding/json"
	"os"

	"github.com/Netsocs-Team/driver.sdk_go/pkg/config"
)

func loadsEventsFromFile(filename ...string) (config.GetEventsAvailableResponse, error) {
	events := config.GetEventsAvailableResponse{}
	jsonfilename := "events.json"
	if len(filename) > 0 {
		jsonfilename = filename[0]
	}
	data, err := os.ReadFile(jsonfilename)
	if err != nil {
		return events, err
	}

	err = json.Unmarshal(data, &events)
	if err != nil {
		return events, err
	}
	return events, nil
}

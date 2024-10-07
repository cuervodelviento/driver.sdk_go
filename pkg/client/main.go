package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Netsocs-Team/driver.sdk_go/pkg/config"
	"github.com/Netsocs-Team/driver.sdk_go/pkg/tools"
)

type NetsocsDriverClient struct {
	driverKey     string
	driverHubHost string
	isSSL         bool
}

func NewNetsocsDriverClient(driverKey string, driverHubHost string, isSSL bool) *NetsocsDriverClient {
	client := &NetsocsDriverClient{
		driverKey:     driverKey,
		driverHubHost: driverHubHost,
		isSSL:         isSSL,
	}

	// If the events.json file exists, add the handler for the actionListenEvents
	// for create a default behavior for the actionListenEvents
	events, err := loadsEventsFromFile()
	if err == nil {
		client.AddConfigHandler(config.GET_EVENTS_AVAILABLE, func(valueMessage config.HandlerValue) (interface{}, error) {
			return events, nil
		})
	}

	return client
}

func (d *NetsocsDriverClient) GetChildren(parentId int) ([]Device, error) {
	url := d.buildURL(fmt.Sprintf("get_childs/%d", parentId))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", d.driverKey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var devices []Device
	if err := json.NewDecoder(res.Body).Decode(&devices); err != nil {
		return nil, err
	}

	for i := range devices {
		if devices[i].Params != nil && devices[i].Params["child_id"] != nil {
			childrenId := devices[i].Params["child_id"].(string)
			devices[i].ChildID = childrenId
		}
	}
	return devices, nil
}

func (d *NetsocsDriverClient) UploadFileAndGetURL(file *os.File) (string, error) {
	return tools.UploadFileAndGetURL(d.driverHubHost, d.driverKey, file)
}

func (d *NetsocsDriverClient) ListenConfig() error {
	return config.ListenConfig(d.driverHubHost, d.driverKey)
}

func (d *NetsocsDriverClient) AddConfigHandler(configKey config.NetsocsConfigKey, configHandler config.FuncConfigHandler) error {
	return config.AddConfigHandler(configKey, configHandler)
}

func (d *NetsocsDriverClient) buildURL(uri string) string {
	if d.isSSL {
		return fmt.Sprintf("https://%s/api/v1/%s", d.driverHubHost, uri)
	}
	return fmt.Sprintf("http://%s/api/v1/%s", d.driverHubHost, uri)
}

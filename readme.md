# Netsocs Driver SDK

This is the official SDK for Netsocs Drivers

## Installation

```bash
go get github.com/Netsocs-Team/driver.sdk_go.git
```

## Usage


```go
package main

import "github.com/Netsocs-Team/driver.sdk_go/pkg/config"

func main() {
    handle := func(valueMessage config.HandlerValue) (interface{}, error) {
		// build connection for driver from device data
		connection := buildConnectionFromDeviceData(valueMessage.DeviceData)
        // this is logic of the driver
		channels := sunapi.Channels{
			Connection: connection,
		}
		// must return type of response type and error
		return channels.GetChannels() // return config.GetChannelResponse and error
	}

	config.AddConfigHandler(
		config.GET_CHANNELS,         // first param is key of config
		handle,                      // handle is a func that get a config.HandlerValue and return T and error; T is a type of response type
		struct{}{},                  // request type, when request type is struct{}{} it means that request type is empty
		config.GetChannelResponse{}, // response type, when response type is config.GetChannelResponse{} it means that response type is config.GetChannelResponse
	)
}

```
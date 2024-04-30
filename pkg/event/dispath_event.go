package event

import (
	"bytes"
	"fmt"
	"net/http"
)

type Dispatcher struct {
	DriverKey string
	Host      string
	Port      int

	// event fields
	imageURL      *string
	deviceId      *int
	originalValue *string
	channelNumber *int
	personId      *string
	videoURL      *string
}

func NewDispatcher(driverHubHost string, driverKey string, deviceId int) *Dispatcher {
	return &Dispatcher{
		DriverKey: driverKey,
		// Default values
		Host:     "netsocs.local",
		Port:     3070,
		deviceId: &deviceId,
	}
}

func (e *Dispatcher) SetImageURL(imageURL string) {
	e.imageURL = &imageURL
}

func (e *Dispatcher) SetOriginalValue(originalValue string) {
	e.originalValue = &originalValue
}

func (e *Dispatcher) SetChannelNumber(channelNumber int) {
	e.channelNumber = &channelNumber
}

func (e *Dispatcher) SetPersonId(personId string) {
	e.personId = &personId
}

func (e *Dispatcher) SetVideoURL(videoURL string) {
	e.videoURL = &videoURL
}

func (e *Dispatcher) DispatchAndResetFields(eventKey string) error {
	buff := "{"
	if e.imageURL != nil {
		buff += "\"eventPhotoUrl\": \"" + *e.imageURL + "\","
	}
	if e.deviceId != nil {
		buff += fmt.Sprintf("\"deviceId\": %d,", *e.deviceId)
	}
	if e.originalValue != nil {
		buff += "\"originalValue\": \"" + *e.originalValue + "\","
	}
	if e.channelNumber != nil {
		buff += fmt.Sprintf("\"deviceByProps\": [ \"\"parentDevice\": %d\", \"\"channelNumber\": %d\"],", *e.deviceId, *e.channelNumber)
	}
	if e.personId != nil {
		userType, userId, err := getUserTypeAndUserIdFromPersonID(*e.personId)
		if err != nil {
			return err
		}
		buff += fmt.Sprintf("\"userType\": %d, \"userId\": %d,", userType, userId)
	}
	if e.videoURL != nil {
		buff += "\"eventVideoUrl\": \"" + *e.videoURL + "\","
	}
	buff = buff[:len(buff)-1] + "}"
	body := []byte(buff)
	_, err := http.Post(fmt.Sprintf("http://%s:%d/v1/topologia/misc/%s", e.Host, e.Port, eventKey), "application/json", bytes.NewReader(body))

	// Reset fields
	e.imageURL = nil
	e.originalValue = nil
	e.channelNumber = nil
	e.personId = nil
	e.videoURL = nil

	return err

}

// Path: pkg/event/main.go

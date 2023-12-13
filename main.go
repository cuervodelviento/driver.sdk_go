package main

import (
	"fmt"

	"github.com/Netsocs-Team/driver.sdk_go/pkg/event"
)

func main() {
	dispatcher := event.NewEventDispatcher("24.126.103.213", "das")
	err := dispatcher.DispatchWithChannels("detectionAudioCamera", 184, 4)
	fmt.Println(err)
}

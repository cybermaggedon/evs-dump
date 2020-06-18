package main

import (
	"encoding/json"
	"fmt"
	"github.com/cybermaggedon/evs-golang-api"
	pb "github.com/cybermaggedon/evs-golang-api/protos"
	"log"
)

type DumpConfig struct {
	*evs.Config
}

func NewDumpConfig() *DumpConfig {
	return &DumpConfig{
		Config: evs.NewConfig("evs-dump", "cyberprobe"),
	}
}

// Dump analytic, just dumps out a kinda JSON equivalent of the event message.
type Dump struct {
	*DumpConfig
	*evs.EventSubscriber
	evs.Interruptible
}

func NewDump(dc *DumpConfig) *Dump {

	d := &Dump{DumpConfig: dc}

	var err error
	d.EventSubscriber, err = evs.NewEventSubscriber(d.Name, d.Input, d)
	if err != nil {
		log.Fatal(err)
	}

	d.RegisterStop(d)

	return d
}

// Event handler
func (d *Dump) Event(ev *pb.Event, properties map[string]string) error {

	// Marshal as JSON.  This doesn't work well, I couldn't get protojson
	// to work on these messages.

	j, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	fmt.Println(string(j))

	return nil

}

func main() {

	dc := NewDumpConfig()
	d := NewDump(dc)
	d.Run()
	log.Print("Shutdown.")

}

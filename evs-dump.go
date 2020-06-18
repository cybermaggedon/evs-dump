package main

import (
	"encoding/json"
	"fmt"
	evs "github.com/cybermaggedon/evs-golang-api"
	"log"
)

type DumpConfig struct {
	*evs.Config
}

func NewDumpConfig() *DumpConfig {
	return &DumpConfig{
		Config: evs.NewConfig("cyberprobe"),
	}
}

// Dump analytic, just dumps out a kinda JSON equivalent of the event message.
type Dump struct {
	*DumpConfig
	*evs.EventSubscriber
	*evs.EventProducer
	evs.Interruptible
}

func NewDump(dc *DumpConfig) *Dump {

	d := &Dump{ DumpConfig: dc }

	var err error
	d.EventSubscriber, err = evs.NewEventSubscriber("evs-dump", d.Input, d)
	if err != nil {
		log.Fatal(err)
	}

	d.EventProducer, err = evs.NewEventProducer(d.Outputs)
	if err != nil {
		log.Fatal(err)
	}

	d.RegisterStop(d)

	return d
}
	

// Event handler
func (d *Dump) Event(ev *evs.Event, properties map[string]string) error {

	// Marshal as JSON.  This doesn't work well, I couldn't get protojson
	// to work on these messages.

	j, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	fmt.Println(string(j))

	d.Output(ev, properties)

	return nil

}

func main() {

	dc := NewDumpConfig()

	d := NewDump(dc)
	
	d.Run()

	log.Print("Shutdown.")

}


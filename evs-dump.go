package main

import (
	"encoding/json"
	"fmt"
	evs "github.com/cybermaggedon/evs-golang-api"
	"os"
	"strings"
)

// Dump analytic, just dumps out a kinda JSON equivalent of the event message.
type Dump struct {
	evs.EventAnalytic
}

// Event handler
func (d *Dump) Event(ev *evs.Event, properties map[string]string) error {

	// Marshal as JSON.  This doesn't work well, I couldn't get protojson to work on
	// these messages.
	j, err := json.Marshal(ev)
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	d.OutputEvent(ev, properties)
	return nil
}

func main() {

	d := &Dump{}

	binding, ok := os.LookupEnv("INPUT")
	if !ok {
		binding = "cyberprobe"
	}

	out, ok := os.LookupEnv("OUTPUT")
	if !ok {
		d.Init(binding, []string{}, d)
	} else {
		outarray := strings.Split(out, ",")
		d.Init(binding, outarray, d)
	}

	d.Run()

}


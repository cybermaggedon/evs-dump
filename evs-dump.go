package main

import (
	"encoding/json"
	"fmt"
	evs "github.com/cybermaggedon/evs-golang-api"
)

type Dump struct {
	evs.EventAnalytic
}

func (d *Dump) Event(ev *evs.Event, mp map[string]string) error {
	j, err := json.Marshal(ev)
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	return nil
}

func main() {
	d := &Dump{}
	d.Init("cyberprobe", []string{}, d)
	d.Run()
}


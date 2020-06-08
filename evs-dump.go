package main

import (
	"log"
	"encoding/json"
	"fmt"
	evs "github.com/cybermaggedon/evs-golang-api"
)

type Dump struct {
	evs.EventAnalytic
}

func (d *Dump) Event(ev *evs.Event, mp map[string]string) {
	j, err := json.Marshal(ev)
	if err != nil {
		log.Fatalf("Error:", err)
	}
	fmt.Println(string(j))
}

func main() {
	d := &Dump{}
	d.Init("cyberprobe", []string{}, d)
	d.Run()
}


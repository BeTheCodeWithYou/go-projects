package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	http.HandleFunc("/events", sseHandler)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("unable to start server %s", err.Error())
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	memoryT := time.NewTicker(time.Second)
	defer memoryT.Stop()

	cupT := time.NewTicker(time.Second)
	defer cupT.Stop()

	clientDone := r.Context().Done()

	rc := http.NewResponseController(w)
	for {
		select {
		case <-clientDone:
			fmt.Println("client has disconnected")
		case <-memoryT.C:
			vm, err := mem.VirtualMemory()
			if err != nil {
				log.Printf("unable to get memory: %s", err.Error())
				return
			}
			if _, err := fmt.Fprintf(w, "event:memory\ndata:Total: %d, Used: %d, Percentage: %.2f%%\n\n", vm.Total, vm.Used, vm.UsedPercent); err != nil {
				log.Printf("unable to write %s", err.Error())
				return
			}
			rc.Flush()
		case <-cupT.C:
			cpu, err := cpu.Times(false)
			if err != nil {
				log.Printf("unable to get CPU data: %s", err.Error())
				return
			}
			if _, err := fmt.Fprintf(w, "event:cpu\ndata:User: %.2f, system: %.2f, Idle: %.2f\n\n", cpu[0].User, cpu[0].System, cpu[0].Idle); err != nil {
				log.Printf("unable to write %s", err.Error())
				return
			}
			rc.Flush()
		}
	}
}

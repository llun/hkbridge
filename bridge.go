package hkbridge

import (
	"log"
	"net"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	hclog "github.com/brutella/hc/log"
	"github.com/llun/hkbridge/accessories"
)

type Bridge struct {
	*accessory.Accessory

	Worker *accessories.Worker
}

func NewBridge(config accessories.Config) *Bridge {
	worker := accessories.NewWorker()
	info := accessory.Info{
		Name:         config.Name,
		Manufacturer: config.Manufacturer,
		SerialNumber: config.SerialNumber,
		Model:        config.Model,
	}
	acc := Bridge{
		Accessory: accessory.New(info, accessory.TypeBridge),
		Worker:    worker,
	}
	go worker.Run()
	return &acc
}

func Start() {
	config, err := accessories.ReadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var iface *net.Interface = nil
	if config.Interface != "" {

		interfaces, _ := net.Interfaces()
		for _, targetInterface := range interfaces {
			if targetInterface.Name == config.Interface {
				iface = &targetInterface
				log.Printf("Setting up accessories on %v", iface.Name)
				break
			}
		}
	}

	if config.Debug {
		hclog.Debug.Enable()
	}

	bridge := NewBridge(config)
	accessories := SetupAccessories(config, iface, bridge.Worker)
	t, err := hc.NewIPTransport(hc.Config{
		Pin:  config.Pin,
		Port: config.Port,
	}, bridge.Accessory, accessories...)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}

package hkbridge

import (
	"log"
	"net"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	hclog "github.com/brutella/hc/log"
)

type Bridge struct {
	*accessory.Accessory
}

func NewBridge(config Config) *Bridge {
	info := accessory.Info{
		Name:         config.Name,
		Manufacturer: config.Manufacturer,
		SerialNumber: config.SerialNumber,
		Model:        config.Model,
	}
	acc := Bridge{
		Accessory: accessory.New(info, accessory.TypeBridge),
	}
	return &acc
}

func Start() {
	config, err := ReadConfig("config.json")
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

	accessories := SetupAccessories(config, iface)
	bridge := NewBridge(config)
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

package hkbridge

import (
	"log"
	"net"

	"github.com/brutella/hc/accessory"
	. "github.com/deckarep/golang-set"

	"github.com/llun/hkbridge/accessories"
	"github.com/llun/hksensibo"
	"github.com/llun/hksoundtouch"
	"github.com/llun/hkwifioccupancy"
)

func SetupAccessories(config accessories.Config, iface *net.Interface, worker *accessories.Worker) []*accessory.Accessory {
	hkAccessories := make([]*accessory.Accessory, 0, 10)
	for _, accessory := range config.Accessories {
		switch accessory.Type {
		case "github.com/llun/hksoundtouch":
			hkAccessories = append(hkAccessories, soundtouch.AllAccessories(accessory, iface, worker)...)
		case "github.com/llun/hkwifioccupancy":
			sensorAccessory := setupWifiOccupancy(accessory, iface)
			if sensorAccessory != nil {
				hkAccessories = append(hkAccessories, sensorAccessory)
			}
		case "github.com/llun/hksensibo":
			hkAccessories = append(hkAccessories, hksensibo.AllAccessories(accessory, iface, worker)...)
		}

	}
	return hkAccessories
}

func setupWifiOccupancy(config accessories.AccessoryConfig, iface *net.Interface) *accessory.Accessory {
	option := config.Option

	presenceFile, ok := option["file"].(string)
	if !ok {
		log.Println("Cannot read presence file option")
		return nil
	}

	macAddresses, ok := option["addresses"].([]interface{})
	if !ok {
		log.Println("Cannot read addresses file option")
		return nil
	}
	sensor := wifioccupancy.NewSensor(presenceFile, NewSet(macAddresses...))
	return sensor.Accessory
}

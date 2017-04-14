package hkbridge

import (
	"log"
	"net"

	"github.com/brutella/hc/accessory"
	. "github.com/deckarep/golang-set"

	"github.com/llun/hksoundtouch"
	"github.com/llun/hkwifioccupancy"
)

func SetupAccessories(config Config, iface *net.Interface) []*accessory.Accessory {
	hkAccessories := make([]*accessory.Accessory, 0, 10)
	for _, accessory := range config.Accessories {
		switch accessory.Type {
		case "github.com/llun/hksoundtouch":
			hkAccessories = append(hkAccessories, setupSoundtouch(accessory, iface)...)
		case "github.com/llun/hkwifioccupancy":
			sensorAccessory := setupWifiOccupancy(accessory, iface)
			if sensorAccessory != nil {
				hkAccessories = append(hkAccessories, sensorAccessory)
			}
		}
	}
	return hkAccessories
}

func setupSoundtouch(config AccessoryConfig, iface *net.Interface) []*accessory.Accessory {
	speakers := soundtouch.Lookup(iface)
	soundtouchAccessories := make([]*accessory.Accessory, len(speakers))
	for idx, speaker := range speakers {
		soundtouchAccessories[idx] = speaker.Accessory
	}
	return soundtouchAccessories
}

func setupWifiOccupancy(config AccessoryConfig, iface *net.Interface) *accessory.Accessory {
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

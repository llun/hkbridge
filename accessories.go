package hkbridge

import (
	"log"
	"net"

	"github.com/brutella/hc/accessory"
	. "github.com/deckarep/golang-set"

	"github.com/llun/hksensibo"
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
		case "github.com/llun/hksensibo":
			hkAccessories = append(hkAccessories, setupSensibo(accessory, iface)...)
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
	log.Println("Soundtouchs, %v", soundtouchAccessories)
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

func setupSensibo(config AccessoryConfig, iface *net.Interface) []*accessory.Accessory {
	option := config.Option

	key, ok := option["key"].(string)
	if !ok {
		log.Println("Cannot read sensibo key")
		return nil
	}

	sensibos := hksensibo.Lookup(key)
	sensiboAccessories := make([]*accessory.Accessory, len(sensibos))
	for idx, sensibo := range sensibos {
		sensiboAccessories[idx] = sensibo.Accessory
	}
	return sensiboAccessories
}

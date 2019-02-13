package netxml

import (
	//"fmt"
	"github.com/jonas-p/go-shp"
	"log"
)

func WriteNetworkSHP(root *Root, file string) (count uint32) {

	count = 0
	shape, err := shp.Create(file, shp.POINT)
	if err != nil {
		log.Println(err)
		return
	}
	defer shape.Close()

	fields := []shp.Field{
		shp.StringField("BSSID", 32),
		shp.StringField("SSID", 32),
		shp.NumberField("Packets", 32),
		shp.StringField("WPS", 32),
	}
	// Can't write Packets
	shape.SetFields(fields)

	networks := root.WirelessNetworks

	for index, network := range networks {
		lat := network.GPS.Lat
		lon := network.GPS.Lon

		// Skip network if no GPS information
		if lat == 0 || lon == 0 {
			continue
		}

		point := shp.Point{lon, lat}
		shape.Write(&point)
		shape.WriteAttribute(index, 0, network.BSSID)
		shape.WriteAttribute(index, 1, network.SSID.Essid)
		shape.WriteAttribute(index, 2, network.SSID.Packets)
		shape.WriteAttribute(index, 3, network.SSID.Wps)
		count += 1
	}
	return
}

func WriteClientSHP(root *Root, file string) (count uint32) {

	count = 0
	shape, err := shp.Create(file, shp.POINT)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer shape.Close()
	fields := []shp.Field{
		shp.StringField("Mac", 32),
		shp.StringField("BSSID", 32),
		shp.StringField("SSID", 32),
	}

	networks := root.WirelessNetworks
	shape.SetFields(fields)
	for _, network := range networks {

		//clientCount := len(network.WirelessClients)

		for _, client := range network.WirelessClients {
			if client.GPS.Lat == 0 || client.GPS.Lon == 0 || client.Mac == network.BSSID {
				continue
			}
			point := shp.Point{client.GPS.Lon, client.GPS.Lat}
			shape.Write(&point)
			shape.WriteAttribute(int(count), 0, client.Mac)
			shape.WriteAttribute(int(count), 1, network.BSSID)
			shape.WriteAttribute(int(count), 2, network.SSID.Essid)
			count += 1

		}

	}
	return
}

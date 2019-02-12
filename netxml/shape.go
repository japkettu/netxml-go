package netxml

import (
	//"fmt"
	"github.com/jonas-p/go-shp"
	"log"
)

// Creates shapefile for wireless networks
//func nwWrite(network *WirelessNetwork, index int) {}
func WriteSHP(root *Root, file string) (count uint32) {

	count = 0

	shape, err := shp.Create(file, shp.POINT)
	if err != nil {
		log.Println(err)
		return
	}
	defer shape.Close()

	networks := root.WirelessNetworks

	for index, network := range networks {
		lat := network.GPS.Lat
		lon := network.GPS.Lon

		// Skip network if no GPS information
		if lat == 0 || lon == 0 {
			continue
		}

		fields := []shp.Field{
			shp.StringField("BSSID", 32),
			shp.StringField("SSID", 32),
			shp.NumberField("Packets", 32),
			shp.StringField("WPS", 32),
		}
		// Jostain syystä packets osio jää tyhjäksi
		// Johtuu varmaan määritetystä koosta
		shape.SetFields(fields)

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

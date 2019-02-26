package netxml

import (
	"github.com/jonas-p/go-shp"
	"log"
)

var (
	err   error
	count uint32
	shape *shp.Writer
)

func writeAttr(shape *shp.Writer, count *uint32, t []interface{}) {
	for i, el := range t {
		shape.WriteAttribute(int(*count), i, el)
	}
	*count += 1
}

func WriteNetworkSHP(root *Root, file string) (count uint32) {

	count = 0
	shape, err = shp.Create(file, shp.POINT)
	if err != nil {
		log.Println(err)
		return
	}
	defer shape.Close()

	fields := []shp.Field{
		shp.StringField("BSSID", 32),
		shp.StringField("SSID", 32),
		shp.NumberField("Packets", 10),
		shp.StringField("WPS", 32),
		shp.NumberField("Channel", 10),
		shp.StringField("Time", 24),
		shp.StringField("Encryption", 32),
	}

	shape.SetFields(fields)
	networks := root.WirelessNetworks

	for _, network := range networks {
		lat := network.GPS.Lat
		lon := network.GPS.Lon

		// Skip network if no GPS information
		if lat == 0 || lon == 0 {
			continue
		}

		point := shp.Point{lon, lat}
		shape.Write(&point)
		encryption := getEnc(&network)

		t := []interface{}{network.BSSID, network.SSID.Essid, network.SSID.Packets,
			network.SSID.Wps, network.Channel,
			network.SeenCard.Time, encryption}

		writeAttr(shape, &count, t)

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
		shp.NumberField("Packets", 10),
		shp.StringField("Time", 24),
	}

	networks := root.WirelessNetworks
	shape.SetFields(fields)
	for _, network := range networks {

		for _, client := range network.WirelessClients {
			if client.GPS.Lat == 0 || client.GPS.Lon == 0 || client.Mac == network.BSSID {
				continue
			}
			point := shp.Point{client.GPS.Lon, client.GPS.Lat}
			shape.Write(&point)

			t := []interface{}{client.Mac, network.BSSID, network.SSID.Essid,
				client.Packets.Total, client.SeenCard.Time}

			writeAttr(shape, &count, t)
		}
	}
	return
}

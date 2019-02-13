package netxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"log"
)

/*

ToDo:
- add encryption information
- add year to bsstimestamp
- add timestamp information
*/

func Parse(data []byte) *Root {

	var root Root
	reader := bytes.NewReader([]byte(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&root)
	if err != nil {
		log.Fatal(err)
	}

	return &root
}

func FileInfo(root *Root) {

	cs := root.CardSource
	nw := root.WirelessNetworks
	fmt.Printf("Interface: %s\nPackets: %d\nNetworks: %d\n",
		cs.Interface, cs.Packets, len(nw))
}

func printNetwork(network *WirelessNetwork) {

	BSSID := network.BSSID
	ESSID := network.SSID.Essid
	Wps := network.SSID.Wps
	bsstimestamp := network.Bsstimestamp
	lat := network.GPS.Lat
	lon := network.GPS.Lon
	time := network.SeenCard.Time

	fmt.Printf("%-20s %-35s [%f, %f] WPS:%-15s bssts: %s  Time: %s \n",
		BSSID, ESSID, lat, lon, Wps, bsstimestamp, time)
}
func printClients(clients []WirelessClient, totalClients *int) {

	nwClients := len(clients) - 1 // filter only clients connected to network
	*totalClients += nwClients

	if nwClients > 0 {
		for _, client := range clients[:nwClients] {
			fmt.Printf("\t%-20s [%d, %d] [Data packets:%d Total packets:%-3d] Time:%s\n",
				client.Mac, client.SNR.MaxDBM, client.SNR.MinDBM,
				client.Packets.Data, client.Packets.Total, client.SeenCard.Time)
		}
	}
}

// Prints struct. Useful for testing.
func Print(root *Root) {

	var totalClients int = 0
	networks := root.WirelessNetworks
	for _, network := range networks {

		printNetwork(&network)
		printClients(network.WirelessClients, &totalClients)

	}
	fmt.Printf("Networks:%d \nClients:%d\n", len(networks), totalClients)
}

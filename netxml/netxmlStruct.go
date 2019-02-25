package netxml

import (
	"encoding/xml"
)

type Root struct {
	XMLName          xml.Name          `xml:"detection-run"`
	CardSource       CardSource        `xml:"card-source"`
	WirelessNetworks []WirelessNetwork `xml:"wireless-network"`
}

type CardSource struct {
	XMLName   xml.Name `xml:"card-source"`
	UUID      string   `xml:"uuid, attr"`
	Interface string   `xml:"card-interface"`
	Type      string   `xml:"card-type"`
	Packets   int      `xml:"card-packets"`
	Channels  string   `xml:"card-channels"`
}

type SeenCard struct {
	XMLName xml.Name `xml:"seen-card"`
	Time    string   `xml:"seen-time"`
}

type WirelessNetwork struct {
	XMLName         xml.Name         `xml:"wireless-network"`
	BSSID           string           `xml:"BSSID"`
	Manuf           string           `xml:"manuf"`
	Channel         int              `xml:"channel"`
	Bsstimestamp    string           `xml:"bsstimestamp"`
	WirelessClients []WirelessClient `xml:"wireless-client"`
	SSID            SSID
	GPS             GPS
	SNR             SNR
	Packets         Packets
	SeenCard        SeenCard
}

type SNR struct {
	XMLName xml.Name `xml:"snr-info"`
	LastDBM int      `xml:"last_signal_dbm"`
	MinDBM  int      `xml:"min_signal_dbm"`
	MaxDBM  int      `xml:"max_signal_dbm"`
}

type GPS struct {
	XMLName xml.Name `xml:"gps-info"`
	Lat     float64  `xml:"avg-lat"`
	Lon     float64  `xml:"avg-lon"`
	Alt     float64  `xml:"avg-alt"`
}

type SSID struct {
	XMLName xml.Name `xml:"SSID"`
	//FirstTime   string   `xml:"first-time, attr"`
	Type        string   `xml:"type"`
	Packets     int      `xml:"packets"`
	Essid       string   `xml:"essid"`
	Wps         string   `xml:"wps"`
	Encryptions []string `xml:"encryption"`
}
type Packets struct {
	XMLName   xml.Name `xml:"packets"`
	LLC       int      `xml:"LLC"`
	Data      int      `xml:"data"`
	Crypt     int      `xml:"crypt"`
	Total     int      `xml:"total"`
	Fragments int      `xml:"fragments"`
	Retries   int      `xml:"retries"`
}

type WirelessClient struct {
	XMLName  xml.Name `xml:"wireless-client"`
	Mac      string   `xml:"client-mac"`
	GPS      GPS
	SNR      SNR
	Packets  Packets
	SeenCard SeenCard
}

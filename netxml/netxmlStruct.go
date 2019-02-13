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
	Channel         uint8            `xml:"channel"`
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
	LastDBM int8     `xml:"last_signal_dbm"`
	MinDBM  int8     `xml:"min_signal_dbm"`
	MaxDBM  int8     `xml:"max_signal_dbm"`
}
type GPS struct {
	XMLName xml.Name `xml:"gps-info"`
	Lat     float64  `xml:"avg-lat"`
	Lon     float64  `xml:"avg-lon"`
	Alt     float64  `xml:"avg-alt"`
}

type SSID struct {
	XMLName xml.Name `xml:"SSID"`
	Type    string   `xml:"type"`
	Packets uint32   `xml:"packets"`
	Essid   string   `xml:"essid"`
	Wps     string   `xml:"wps"`
}
type Packets struct {
	XMLName   xml.Name `xml:"packets"`
	LLC       uint32   `xml:"LLC"`
	Data      uint32   `xml:"data"`
	Crypt     uint32   `xml:"crypt"`
	Total     uint32   `xml:"total"`
	Fragments uint32   `xml:"fragments"`
	Retries   uint32   `xml:"retries"`
}
type WirelessClient struct {
	XMLName  xml.Name `xml:"wireless-client"`
	Mac      string   `xml:"client-mac"`
	GPS      GPS
	SNR      SNR
	Packets  Packets
	SeenCard SeenCard
}

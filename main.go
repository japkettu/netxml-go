package main

import (
	"flag"
	"fmt"
	"github.com/japkettu/netxml-go/netxml"
	"io/ioutil"
	"log"
)

func main() {

	file := flag.String("f", "", "File or path name")
	flag.Parse()

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	shapeout := "./output/networks.shp"
	fmt.Printf("File: %s\n", *file)
	root := netxml.Parse(data)
	netxml.Print(root)
	netxml.FileInfo(root)
	count := netxml.WriteNetworkSHP(root, shapeout)
	fmt.Printf("Shapefile: %d\n", count)

	clientFile := "./output/clients.shp"
	clientCount := netxml.WriteClientSHP(root, clientFile)
	fmt.Printf("Clients written to shapefile: %d\n", clientCount)

}

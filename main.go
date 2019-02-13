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
	clientSHP := flag.String("cs", "", "Client shapefile")
	nwSHP := flag.String("ns", "", "Network shapefile")
	flag.Parse()

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	root := netxml.Parse(data)
	netxml.Print(root)
	netxml.FileInfo(root)
	fmt.Printf("File: %s\n", *file)
	if *nwSHP != "" {
		count := netxml.WriteNetworkSHP(root, *nwSHP)
		fmt.Printf("Networks written to shapefile: %d\n", count)
	}
	if *clientSHP != "" {
		clientCount := netxml.WriteClientSHP(root, *clientSHP)
		fmt.Printf("Clients written to shapefile: %d\n", clientCount)
	}
}

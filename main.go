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
	fPrint := flag.Bool("p", false, "Print network information and clients")
	fInfo := flag.Bool("i", false, "Print netxml info")
	clientSHP := flag.String("cs", "", "Write clients to shapefile")
	nwSHP := flag.String("ns", "", "Write networks to shapefile")
	flag.Parse()

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	root := netxml.Parse(data)
	fmt.Printf("File: %s\n", *file)
	if *fPrint {
		netxml.Print(root)
	}
	if *fInfo {
		netxml.FileInfo(root)
	}
	if *nwSHP != "" {
		count := netxml.WriteNetworkSHP(root, *nwSHP)
		fmt.Printf("Networks written to shapefile: %d\n", count)
	}
	if *clientSHP != "" {
		clientCount := netxml.WriteClientSHP(root, *clientSHP)
		fmt.Printf("Clients written to shapefile: %d\n", clientCount)
	}
}

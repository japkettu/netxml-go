# netxml-go

A commandline tool and library for converting netxml files to other file formats. At this moment only ESRI Shapefile is supported. 

Available file formats: 

- ESRI Shapefile


## Usage:

```sh
$ ./netxml-go -h
Usage of ./main:
  -cs string
    	Write clients to shapefile
  -f string
    	File or path name
  -i	Print netxml info
  -ns string
    	Write networks to shapefile
  -p	Print network information and connected clients
```

## Example 1 

Create shapefiles for networks and clients. Print file information.

```sh
$ ./netxml-go -f path/to/Kismet.netxml -cs path/to/clients.shp -ns path/to/networks.shp -i
File: path/to/Kismet.netxml
Interface: wlan0mon
Packets: 117775
Networks: 2904
Networks written to shapefile: 2899
Clients written to shapefile: 665
```



## Installation:

```sh
$ git clone https://github.com/japkettu/netxml-go
$ go build main.go
```



## Dependencies:

- golang.org/x/net/html/charset
- github.com/jonas-p/go-shp



## TODO:

- add timestamp information
- fix missing year in bss timestamp
- add  encryption information
- convert multiple files at once




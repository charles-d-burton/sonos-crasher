package main

import (
	"log"

	sonos "github.com/ianr0bkny/go-sonos"
	"github.com/ianr0bkny/go-sonos/ssdp"
)

// This code identifies UPnP devices on the netork that support the
// MusicServices API.
func main() {
	log.Print("go-sonos example discovery\n")

	mgr := ssdp.MakeManager()

	// Discover()
	//  eth0 := Network device to query for UPnP devices
	// 11209 := Free local port for discovery replies
	// false := Do not subscribe for asynchronous updates
	mgr.Discover("en0", "11209", false)

	// SericeQueryTerms
	// A map of service keys to minimum required version
	qry := ssdp.ServiceQueryTerms{
		ssdp.ServiceKey("schemas-upnp-org-MusicServices"): -1,
	}

	// Look for the service keys in qry in the database of discovered devices
	result := mgr.QueryServices(qry)
	if devList, has := result["schemas-upnp-org-MusicServices"]; has {
		for _, dev := range devList {
			log.Printf("%s %s %s %s %s\n", dev.Product(), dev.ProductVersion(), dev.Name(), dev.Location(), dev.UUID())
			s := sonos.Connect(dev, nil, sonos.SVC_CONTENT_DIRECTORY|sonos.SVC_AV_TRANSPORT)

			if err := s.Stop(0); nil != err {
				panic(err)
			}
		}
	}
	mgr.Close()
}

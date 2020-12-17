package main

import "encoding/xml"

// Root was generated 2019-02-22 10:43:22 by charles.burton on Charless-MacBook-Pro.
type SonosDevice struct {
	XMLName     xml.Name `xml:"root"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	SpecVersion struct {
		Text  string `xml:",chardata"`
		Major string `xml:"major"` // 1
		Minor string `xml:"minor"` // 0
	} `xml:"specVersion"`
	Device struct {
		Text             string `xml:",chardata"`
		DeviceType       string `xml:"deviceType"`       // urn:schemas-upnp-org:devi...
		FriendlyName     string `xml:"friendlyName"`     // 192.168.5.92 - Sonos Play...
		Manufacturer     string `xml:"manufacturer"`     // Sonos, Inc.
		ManufacturerURL  string `xml:"manufacturerURL"`  // http://www.sonos.com
		ModelNumber      string `xml:"modelNumber"`      // S12
		ModelDescription string `xml:"modelDescription"` // Sonos Play:1
		ModelName        string `xml:"modelName"`        // Sonos Play:1
		ModelURL         string `xml:"modelURL"`         // http://www.sonos.com/prod...
		SoftwareVersion  string `xml:"softwareVersion"`  // 48.2-61220
		HardwareVersion  string `xml:"hardwareVersion"`  // 1.20.1.6-1
		SerialNum        string `xml:"serialNum"`        // 94-9F-3E-F8-70-0C:9
		UDN              string `xml:"UDN"`              // uuid:RINCON_949F3EF8700C0...
		IconList         struct {
			Text string `xml:",chardata"`
			Icon struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id"`       // 0
				Mimetype string `xml:"mimetype"` // image/png
				Width    string `xml:"width"`    // 48
				Height   string `xml:"height"`   // 48
				Depth    string `xml:"depth"`    // 24
				URL      string `xml:"url"`      // /img/icon-S12.png
			} `xml:"icon"`
		} `xml:"iconList"`
		MinCompatibleVersion    string `xml:"minCompatibleVersion"`    // 47.0-00000
		LegacyCompatibleVersion string `xml:"legacyCompatibleVersion"` // 36.0-00000
		ApiVersion              string `xml:"apiVersion"`              // 1.10.1
		MinApiVersion           string `xml:"minApiVersion"`           // 1.1.0
		DisplayVersion          string `xml:"displayVersion"`          // 10.0
		ExtraVersion            string `xml:"extraVersion"`
		RoomName                string `xml:"roomName"`            // Copy Room Left
		DisplayName             string `xml:"displayName"`         // Play:1
		ZoneType                string `xml:"zoneType"`            // 14
		Feature1                string `xml:"feature1"`            // 0x00000000
		Feature2                string `xml:"feature2"`            // 0x00403332
		Feature3                string `xml:"feature3"`            // 0x0001000e
		Seriesid                string `xml:"seriesid"`            // A200
		Variant                 string `xml:"variant"`             // 2
		InternalSpeakerSize     string `xml:"internalSpeakerSize"` // 5
		BassExtension           string `xml:"bassExtension"`       // 75.000
		SatGainOffset           string `xml:"satGainOffset"`       // 6.000
		Memory                  string `xml:"memory"`              // 256
		Flash                   string `xml:"flash"`               // 256
		AmpOnTime               string `xml:"ampOnTime"`           // 10
		RetailMode              string `xml:"retailMode"`          // 0
		ServiceList             struct {
			Text    string `xml:",chardata"`
			Service []struct {
				Text        string `xml:",chardata"`
				ServiceType string `xml:"serviceType"` // urn:schemas-upnp-org:serv...
				ServiceId   string `xml:"serviceId"`   // urn:upnp-org:serviceId:Al...
				ControlURL  string `xml:"controlURL"`  // /AlarmClock/Control, /Mus...
				EventSubURL string `xml:"eventSubURL"` // /AlarmClock/Event, /Music...
				SCPDURL     string `xml:"SCPDURL"`     // /xml/AlarmClock1.xml, /xm...
			} `xml:"service"`
		} `xml:"serviceList"`
		DeviceList struct {
			Text   string `xml:",chardata"`
			Device []struct {
				Text             string `xml:",chardata"`
				DeviceType       string `xml:"deviceType"`       // urn:schemas-upnp-org:devi...
				FriendlyName     string `xml:"friendlyName"`     // 192.168.5.92 - Sonos Play...
				Manufacturer     string `xml:"manufacturer"`     // Sonos, Inc., Sonos, Inc.
				ManufacturerURL  string `xml:"manufacturerURL"`  // http://www.sonos.com, htt...
				ModelNumber      string `xml:"modelNumber"`      // S12, S12
				ModelDescription string `xml:"modelDescription"` // Sonos Play:1 Media Server...
				ModelName        string `xml:"modelName"`        // Sonos Play:1, Sonos Play:...
				ModelURL         string `xml:"modelURL"`         // http://www.sonos.com/prod...
				UDN              string `xml:"UDN"`              // uuid:RINCON_949F3EF8700C0...
				ServiceList      struct {
					Text    string `xml:",chardata"`
					Service []struct {
						Text        string `xml:",chardata"`
						ServiceType string `xml:"serviceType"` // urn:schemas-upnp-org:serv...
						ServiceId   string `xml:"serviceId"`   // urn:upnp-org:serviceId:Co...
						ControlURL  string `xml:"controlURL"`  // /MediaServer/ContentDirec...
						EventSubURL string `xml:"eventSubURL"` // /MediaServer/ContentDirec...
						SCPDURL     string `xml:"SCPDURL"`     // /xml/ContentDirectory1.xm...
					} `xml:"service"`
				} `xml:"serviceList"`
				XRhapsodyExtension struct {
					Text               string `xml:",chardata"`
					Xmlns              string `xml:"xmlns,attr"`
					DeviceID           string `xml:"deviceID"` // urn:rhapsody-real-com:dev...
					DeviceCapabilities struct {
						Text               string `xml:",chardata"`
						InteractionPattern struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"interactionPattern"`
					} `xml:"deviceCapabilities"`
				} `xml:"X_Rhapsody-Extension"`
				XQPlaySoftwareCapability struct {
					Text string `xml:",chardata"` // QPlay:2
					Qq   string `xml:"qq,attr"`
				} `xml:"X_QPlay_SoftwareCapability"`
				IconList struct {
					Text string `xml:",chardata"`
					Icon struct {
						Text     string `xml:",chardata"`
						Mimetype string `xml:"mimetype"` // image/png
						Width    string `xml:"width"`    // 48
						Height   string `xml:"height"`   // 48
						Depth    string `xml:"depth"`    // 24
						URL      string `xml:"url"`      // /img/icon-S12.png
					} `xml:"icon"`
				} `xml:"iconList"`
			} `xml:"device"`
		} `xml:"deviceList"`
	} `xml:"device"`
}

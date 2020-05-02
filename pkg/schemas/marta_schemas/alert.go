package marta_schemas

import "encoding/xml"

type BusAlert struct {
	Text    string `json:"text,omitempty" xml:",chardata"`
	ID      string `xml:"id,attr" json:"id"`
	Title   string `xml:"title,attr" json:"title"`
	Desc    string `xml:"desc" json:"desc"`
	Expires string `xml:"expires" json:"expires"`
}

type RailAlert struct {
	Text    string `json:"text,omitempty" xml:",chardata"`
	ID      string `xml:"id,attr" json:"id"`
	Title   string `xml:"title,attr" json:"title"`
	Desc    string `xml:"desc" json:"desc"`
	Expires string `xml:"expires" json:"expires"`
}

type Alerts struct {
	XMLName xml.Name    `json:"-" xml:"Alerts"`
	Text    string      `json:"text,omitempty" xml:",chardata"`
	Bus     []BusAlert  `xml:"BUS" json:"Bus"`
	Rail    []RailAlert `xml:"RAIL" json:"Rail"`
}

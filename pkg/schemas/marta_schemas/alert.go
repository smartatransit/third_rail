package marta_schemas

import "encoding/xml"

type Alerts struct {
	XMLName xml.Name `json:"-" xml:"Alerts"`
	Text    string   `json:"text,omitempty" xml:",chardata"`
	Bus     []struct {
		Text    string `json:"text,omitempty" xml:",chardata"`
		ID      string `xml:"id,attr" json:"id"`
		Title   string `xml:"title,attr" json:"title"`
		Desc    string `xml:"desc" json:"desc"`
		Expires string `xml:"expires" json:"expires"`
	} `xml:"BUS"`
	Rail []struct {
		Text    string `json:"text,omitempty" xml:",chardata"`
		ID      string `xml:"id,attr" json:"id"`
		Title   string `xml:"title,attr" json:"title"`
		Desc    string `xml:"desc" json:"desc"`
		Expires string `xml:"expires" json:"expires"`
	} `xml:"RAIL" json:"rail"`
}

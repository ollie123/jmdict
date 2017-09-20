package jmdict

import "encoding/xml"

// Presence represents a bool determined by the presence of an XML element.
type Presence bool

// UnmarshalXML unmarshals an XML element into a Presence.
func (p *Presence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if err := d.Skip(); err != nil {
		return err
	}
	*p = true
	return nil
}

// YesNo represents a bool determined by a "y" or "n" value in an XML attribute.
type YesNo bool

// UnmarshalXMLAttr unmarshals an XML attribute into a YesNo.
func (y *YesNo) UnmarshalXMLAttr(attr xml.Attr) error {
	// Count only "y" as true.
	if attr.Value == "y" {
		*y = true
	} else {
		*y = false
	}
	return nil
}

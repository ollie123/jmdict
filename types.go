package jmdict

import (
	"encoding/xml"
	"errors"
)

// PresenceBool is a bool determined by the presence of an XML element.
type PresenceBool bool

// UnmarshalXML unmarshals an XML element into a Presence.
func (v *PresenceBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if err := d.Skip(); err != nil {
		return err
	}
	*v = true
	return nil
}

// YesNoBool is a bool determined by a 'y' or 'n' value in an XML attribute.
type YesNoBool bool

// UnmarshalXMLAttr unmarshals an XML attribute into a YesNoBool.
func (v *YesNoBool) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "N", "n":
		*v = false
	case "Y", "y":
		*v = true
	default:
		return errors.New("invalid value for YesNoBool")
	}
	return nil
}

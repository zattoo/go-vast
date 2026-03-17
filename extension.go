package vast

import "encoding/xml"

type SSAICreativeID struct {
	CreativeID string `xml:"creativeId,attr,omitempty" json:"creativeId,omitempty"`
	Data       string `xml:",cdata" json:"data,omitempty"`
}

type ExtensionParameter struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Value string `xml:",cdata" json:"value,omitempty"`
}

// Extension represent arbitrary XML provided by the platform to extend the
// VAST response or by custom trackers.
type Extension struct {
	Type           string     `xml:"type,attr,omitempty"`
	CustomTracking []Tracking `xml:"CustomTracking>Tracking,omitempty"  json:",omitempty"`
	// AdVerifications are IAB Open Measurement tags backported to VAST 2 and 3 as an extension
	AdVerifications *[]Verification      `xml:"AdVerifications>Verification,omitempty"  json:",omitempty"`
	SSAICreativeID  *SSAICreativeID      `xml:"SSAICreativeId,omitempty"  json:"ssaiCreativeId,omitempty"`
	Parameters      []ExtensionParameter `xml:"Parameter,omitempty"  json:"parameters,omitempty"`
}

// the extension type as a middleware in the encoding process.
type extension Extension

// MarshalXML implements xml.Marshaler interface.
func (e Extension) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	// create a temporary element from a wrapper Extension, copy what we need to
	// it and return it's encoding.
	var e2 interface{}

	e2 = extension{Type: e.Type, CustomTracking: e.CustomTracking, AdVerifications: e.AdVerifications}

	return enc.EncodeElement(e2, start)
}

// UnmarshalXML implements xml.Unmarshaler interface.
func (e *Extension) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	// decode the extension into a temporary element from a wrapper Extension,
	// copy what we need over.
	var e2 extension
	if err := dec.DecodeElement(&e2, &start); err != nil {
		return err
	}

	// copy the type, customTracking and adVerifications
	e.Type = e2.Type
	e.CustomTracking = e2.CustomTracking
	e.AdVerifications = e2.AdVerifications
	e.SSAICreativeID = e2.SSAICreativeID
	e.Parameters = e2.Parameters

	return nil
}

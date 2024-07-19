package vast

import "encoding/xml"

// Extension represent arbitrary XML provided by the platform to extend the
// VAST response or by custom trackers.
type Extension struct {
	Type           string     `xml:"type,attr,omitempty"`
	CustomTracking []Tracking `xml:"CustomTracking>Tracking,omitempty"  json:",omitempty"`
	// AdVerifications are IAB Open Measurement tags backported to VAST 2 and 3 as an extension
	AdVerifications *[]Verification `xml:"AdVerifications>Verification,omitempty"  json:",omitempty"`
	Data            string          `xml:",innerxml" json:",omitempty"`
}

// the extension type as a middleware in the encoding process.
type extension Extension

type extensionOnlyData struct {
	Type string `xml:"type,attr,omitempty"`
	Data string `xml:",innerxml" json:",omitempty"`
}

// MarshalXML implements xml.Marshaler interface.
func (e Extension) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	// create a temporary element from a wrapper Extension, copy what we need to
	// it and return it's encoding.
	var e2 interface{}
	// if we have custom trackers or ad verifications, we should ignore the data, if not, then we
	// should consider only the data
	if len(e.CustomTracking) == 0 && (e.AdVerifications == nil || len(*e.AdVerifications) == 0) {
		e2 = extensionOnlyData{Type: e.Type, Data: e.Data}
	} else {
		e2 = extension{Type: e.Type, CustomTracking: e.CustomTracking, AdVerifications: e.AdVerifications}
	}

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

	// copy the data only if customTracking and adVerifications are empty
	if len(e.CustomTracking) == 0 && (e.AdVerifications == nil || len(*e.AdVerifications) == 0) {
		e.Data = e2.Data
	}
	return nil
}

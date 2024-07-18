package vast

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	extensionCustomTracking = []byte(`<Extension type="testCustomTracking"><CustomTracking><Tracking event="event.1"><![CDATA[http://event.1]]></Tracking><Tracking event="event.2"><![CDATA[http://event.2]]></Tracking></CustomTracking></Extension>`)
	extensionAdVerification = []byte(`<Extension type="AdVerifications">
    <AdVerifications>
        <Verification vendor="doubleclickbygoogle.com-omid-video">
            <JavaScriptResource apiFramework="omid" browserOptional="true"><![CDATA[https://example.com/verify.js]]></JavaScriptResource>
            <VerificationParameters><![CDATA[example=1&param=2]]></VerificationParameters>
            <TrackingEvents>
                <Tracking event="verificationNotExecuted"><![CDATA[https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=active_view_verification_rejected&errorcode=%5BREASON%5D]]></Tracking>
            </TrackingEvents>
        </Verification>
    </AdVerifications>
</Extension>`)
	extensionData = []byte(`<Extension type="testCustomTracking"><SkippableAdType>Generic</SkippableAdType></Extension>`)
)

func TestExtensionCustomTrackingMarshal(t *testing.T) {
	e := Extension{
		Type: "testCustomTracking",
		CustomTracking: []Tracking{
			{
				Event: "event.1",
				URI:   "http://event.1",
			},
			{
				Event: "event.2",
				URI:   "http://event.2",
			},
		},
	}

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)

	// assert the resulting marshaled extension
	assert.Equal(t, string(extensionCustomTracking), string(xmlExtensionOutput))
}

func TestExtensionCustomTracking(t *testing.T) {
	// unmarshal the Extension
	var e Extension
	assert.NoError(t, xml.Unmarshal(extensionCustomTracking, &e))

	// assert the resulting extension
	assert.Equal(t, "testCustomTracking", e.Type)
	assert.Empty(t, string(e.Data))
	if assert.Len(t, e.CustomTracking, 2) {
		// first event
		assert.Equal(t, "event.1", e.CustomTracking[0].Event)
		assert.Equal(t, "http://event.1", e.CustomTracking[0].URI)
		// second event
		assert.Equal(t, "event.2", e.CustomTracking[1].Event)
		assert.Equal(t, "http://event.2", e.CustomTracking[1].URI)
	}

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)

	// assert the resulting marshaled extension
	assert.Equal(t, string(extensionCustomTracking), string(xmlExtensionOutput))
}

func TestExtensionCustomAdVerification(t *testing.T) {
	// unmarshal the Extension
	var e Extension
	assert.NoError(t, xml.Unmarshal(extensionAdVerification, &e))

	// assert the resulting extension
	assert.Equal(t, "AdVerifications", e.Type)
	assert.Empty(t, e.Data)
	if assert.Len(t, e.AdVerifications, 1) {
		assert.Equal(t, e.AdVerifications[0].Vendor, "doubleclickbygoogle.com-omid-video")
		assert.Equal(t, e.AdVerifications[0].JavaScriptResource[0], JavaScriptResource{
			ApiFramework:    "omid",
			BrowserOptional: true,
			URI:             "https://example.com/verify.js",
		})

		if assert.Len(t, e.AdVerifications[0].TrackingEvents, 1) {
			assert.Equal(t, e.AdVerifications[0].TrackingEvents[0], Tracking{
				Event: "verificationNotExecuted",
				URI:   "https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=active_view_verification_rejected&errorcode=%5BREASON%5D",
			})
		}
	}
}

func TestExtensionGeneric(t *testing.T) {
	// unmarshal the Extension
	var e Extension
	assert.NoError(t, xml.Unmarshal(extensionData, &e))

	// assert the resulting extension
	assert.Equal(t, "testCustomTracking", e.Type)
	assert.Equal(t, "<SkippableAdType>Generic</SkippableAdType>", string(e.Data))
	assert.Empty(t, e.CustomTracking)

	// marshal the extension
	xmlExtensionOutput, err := xml.Marshal(e)
	assert.NoError(t, err)

	// assert the resulting marshaled extension
	assert.Equal(t, string(extensionData), string(xmlExtensionOutput))
}

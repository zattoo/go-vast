package vast

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const vast4StringAttrsXML = `<VAST version="4.2">
    <Ad id="1">
        <InLine>
            <AdSystem>GDFP</AdSystem>
            <AdTitle><![CDATA[t]]></AdTitle>
            <AdVerifications>
                <Verification vendor="v">
                    <ExecutableResource apiFramework="omid" type="application/x-executable"><![CDATA[https://example.com/e.bin]]></ExecutableResource>
                </Verification>
            </AdVerifications>
            <Creatives>
                <Creative>
                    <Linear>
                        <MediaFiles>
                            <InteractiveCreativeFile type="text/html" apiFramework="SIMID" variableDuration="true"><![CDATA[https://example.com/i.html]]></InteractiveCreativeFile>
                            <ClosedCaptionFiles>
                                <ClosedCaptionFile type="text/srt" language="en"><![CDATA[https://example.com/a.srt]]></ClosedCaptionFile>
                            </ClosedCaptionFiles>
                        </MediaFiles>
                    </Linear>
                </Creative>
            </Creatives>
        </InLine>
    </Ad>
</VAST>`

func TestVast4StringAttributes(t *testing.T) {
	var v VAST
	require.NoError(t, xml.Unmarshal([]byte(vast4StringAttrsXML), &v))

	inline := v.Ads[0].InLine
	require.NotNil(t, inline)

	verifications := *inline.AdVerifications
	assert.Equal(t, "application/x-executable", verifications[0].ExecutableResource[0].Type)

	mf := inline.Creatives[0].Linear.MediaFiles
	icf := (*mf.InteractiveCreativeFile)[0]
	assert.Equal(t, "text/html", icf.Type)
	assert.True(t, icf.VariableDuration)

	cc := (*mf.ClosedCaptionFiles)[0]
	assert.Equal(t, "text/srt", cc.Type)
	assert.Equal(t, "en", cc.Language)
	assert.Equal(t, "https://example.com/a.srt", cc.URI)

	// Round trip.
	out, err := xml.Marshal(v)
	require.NoError(t, err)

	var v2 VAST
	require.NoError(t, xml.Unmarshal(out, &v2))
	assert.Equal(t, v, v2)
}

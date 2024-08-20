package vast

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func GetTestdataFileURI(file string) string {
	return filepath.Join("testdata", file)
}

func ReadTestdataFile(file string) ([]byte, error) {
	return os.ReadFile(GetTestdataFileURI(file))
}

func ReadTestdataXML(file string, output interface{}) error {
	data, err := ReadTestdataFile(file)
	if err != nil {
		return fmt.Errorf("failed to read testdata file: %w", err)
	}

	err = xml.Unmarshal(data, &output)
	if err != nil {
		return err
	}

	return nil
}

func TestVastParser(t *testing.T) {
	for _, tc := range []struct {
		name     string
		file     string
		expected VAST
	}{
		{
			name: "inline with verification",
			file: "vast_verification_inline.xml",
			expected: VAST{
				XMLName: xml.Name{
					Local: "VAST",
				},
				Version: "4.0",
				Ads: []Ad{
					{
						ID: "5925032079",
						InLine: &InLine{
							AdSystem: &AdSystem{
								Name: "GDFP",
							},
							AdTitle: PlainString{
								CDATA: "External - Base Asset for Linear Redirect",
							},
							Description: &CDATAString{
								CDATA: "External - Base Asset for Linear Redirect ad",
							},
							Errors: []CDATAString{
								{
									CDATA: "\n                https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=videoplayfailed[ERRORCODE]",
								},
							},
							Impressions: []Impression{
								{
									URI: "\n                https://pagead2.googlesyndication.com/pcs/view?xai=AKAOjsshftjNvjbKP9QAzJDwK7vT8ckUjR8j3ipt-UVk1r98ZRAlQPSxUoq5mti16wWwDqZ5H2t4Ba6Q9RflSU9yUonWsIEWW_0ALB0TRfjV97NW-ALY_Vd29aOHRNfzQGA_QqNH5Je4BgA893seJLACl5q1KDuGnQiOfP1tWvC6KDEaHIGj79J_rJ_YPFRV9khoNLuQTmnW0TI1zr1dHhiR_1NQSoV1DxXn7dRQbffQHYXpLq9MDGgB6pyr9QjQhSlh1EVOqZ33lmXgm-W5a01mEXPD8wlZPoDpos0qMVilhP8lgcSTOs7jQASEL_GrgJ0u6vGBVZPOHrKP0Zxn0ouXA93xtokxQGKvuh-uaTHVjMOYbQgxxbY6KqU49yn2H1AaUHRMiKZ0oZkN-30&sig=Cg0ArKJSzBhGwFhIktYJEAE&uach_m=%5BUACH%5D&adurl=",
								},
							},
							AdVerifications: &[]Verification{
								{
									Vendor: "doubleclickbygoogle.com-omid-video",
									JavaScriptResource: []JavaScriptResource{{
										ApiFramework:    "omid",
										BrowserOptional: true,
										URI:             "\n                        https://www.googletagservices.com/activeview/js/current/rx_omid_video.js",
									}},
									TrackingEvents: []Tracking{{
										Event: "verificationNotExecuted",
										URI:   "\n                            https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=active_view_verification_rejected&errorcode=%5BREASON%5D",
									}},
									VerificationParameters: `{"param":"val"}`,
								},
							},
							Creatives: []Creative{
								{
									ID:         "138381630420",
									AdID:       "IdhwY3cRXqA",
									LegacyAdID: "IdhwY3cRXqA",
									Sequence:   1,
									Linear: &Linear{
										Duration: Duration(10 * time.Second),
										TrackingEvents: []Tracking{
											{
												Event: EventTypeStart,
												URI:   "\n                                https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=part2viewed&ad_mt=%5BAD_MT%5D",
											},
											{
												Event: EventTypeFirstQuartile,
												URI:   "\n                                https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=videoplaytime25&ad_mt=%5BAD_MT%5D",
											},
											{
												Event: EventTypeMidpoint,
												URI:   "\n                                https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=videoplaytime50&ad_mt=%5BAD_MT%5D",
											},
											{
												Event: EventTypeThirdQuartile,
												URI:   "\n                                https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=videoplaytime75&ad_mt=%5BAD_MT%5D",
											},
											{
												Event: EventTypeComplete,
												URI:   "\n                                https://pagead2.googlesyndication.com/pagead/interaction/?ai=Bt7src9CCZofvMqChiM0Pi8qQkAPFnbOVRgAAABABII64hW84AVjUt8DBgwRglfrwgYwHsgETZ29vZ2xlYWRzLmdpdGh1Yi5pb7oBCjcyOHg5MF94bWzIAQXaATRodHRwczovL2dvb2dsZWFkcy5naXRodWIuaW8vZ29vZ2xlYWRzLWltYS1odG1sNS92c2kvwAIC4AIA6gIlLzIxNzc1NzQ0OTIzL2V4dGVybmFsL3V0aWxpdHlfc2FtcGxlc_gC8NEegAMBkAPIBpgD4AOoAwHgBAHSBQYQj6GjiRagBiOoB7i-sQKoB5oGqAfz0RuoB5bYG6gHqpuxAqgHg62xAqgH4L2xAqgH_56xAqgH35-xAqgH-MKxAqgH-8KxAtgHAdIIMQiR4YBwEAEYHTIH64uA7r-AAToPgNCAgICAhAiAgICAgJQuSL39wTpY1cHtiZmGhwPYCAKACgWYCwGqDQJERdAVAfgWAYAXAQ&sigh=UTbooye19j8&label=videoplaytime100&ad_mt=%5BAD_MT%5D",
											},
										},
										VideoClicks: &VideoClicks{
											ClickThroughs: []VideoClick{
												{
													ID:  "GDFP",
													URI: "\n                                https://pagead2.googlesyndication.com/pcs/click?xai=AKAOjstChlSO_p9X4aSy7VmEGpOEAeahfWxkn8PbLWBdBU1AFpofuO5Bbfjh3pBNR9w1r3PSpBvlgqHipXPOrRucZFtCXAYFHLUCU26AVUfrlpRVeMo1ke3nmBnN3HoqiD0srT0NiqYtjNg8XJFTkcaNXF_ElVN54KsdqXWcDPWzvDYM38p_46Z2w1Kc1r5nmWCliNLUeQovIhVGr9oAiVGZslpfQeRD3Q6LZTfucsp9hC40yjvbJsH7GZXpa1T6W6zejGYtZpvwFRPIGHyU4cG0bcr44TcHSxRAnQ01fYyMDSlHZkoAp93uDBVUty584Mb5ZVyBY4qOpoUzUsDk1r_jQ4IXtowH11elYve26Pzw4audSCPZoU1tLf1Z0IpO93ioqjrAYPWrfAlX5uTEIw&sig=Cg0ArKJSzK1NuvSNbMR-&fbs_aeid=%5Bgw_fbsaeid%5D&adurl=https://googleads.github.io/googleads-ima-html5/vsi/",
												},
											},
										},
										Icons: &Icons{
											XMLName: xml.Name{
												Local: "Icons",
											},
											Icon: &[]Icon{
												{
													Program:   "GoogleWhyThisAd",
													Width:     18,
													Height:    18,
													XPosition: "right",
													YPosition: "bottom",
													AltText:   "Google Why This Ad",
													IconClickThrough: &CDATAString{
														CDATA: "\n                                        https://adssettings.google.com/whythisad?source=display&reasons=AUpiJ5ST8AwQvijjZwhBo_iRZiUfNmVO27Su0rhb53hC2QwuZ5tL8bpClRkhj_5dA98tlhsBqmkPgKhlD-LOTz0DtKzOZpEf-M8BK9M-qXbdUU3HA3nABVM_cGI3jupYJgFFytzkxXYCUcpzO8mcgmjAcbTbKc6uzdlDM9eyb1nV-ujw2mOCUn4QBQWqK4KNgXr6B07bghr5GnLYsuxuh30bwH8vlbE4qjcdGxJaCoVK8y8lrKmQ6xxwv61cxMs9N6NV_PZ68NZxqID7H48hlm13UF7woE5-NEcjWY4ndDC8IYt9nyC9kXdbhKup3Re5h28f_8TfMTjU7LOSEgwdc1lcqnuRLvfi7IPgcflf5ECIXclpDEU8IIP9uQYPsEcNLceG5R96TgolZ99ggXgGCrc2ldWwlQHTPyN3Jwswa1QfcgTPPYVPqj5f_t5hnf7wCGsRjVSNkRqbzhqVaeRBZjyRzWYUU2nSpv8-Hp3ZqLUkkJvjaxj0koFGITZ50yS-2SHyv0ODrSxezuPOwPRIR97q0grSbLepKOWeyQEFBJmVfqpdKxUw8F-YkDe_Dy73za5sqVmWZUfY7N5WecAVazgEMTah0ynh2JRYPFXQIRm9BELAZQhi10OAphxisiGMT3RDlZu3EGpIZYg1fUg67_gA7w204t2O-LfpzZTkc4D77I412L5UKMA9iMWdvyimSqbLdxu-DgWJv6S6gnxbvDOECBKvlhtVVo3OOH__F-DNqVgOMQIjEq3f7JV2F_5ayUtspFDgT2U81CHtcFg9bbu_SGQ9bdYHkSGjUzLQWU-EU5tqMpGsy1jI1_HN3zzk18Lh9dL86_TblvP0NLyBb93JaX-udI1wBrUGxA4S9LmsNmrAvpfuKQ59Tz627K5CrO_OK86Iu_PriiERbQhjgIVYp8FONaaqz4Xw4UfXFWUpKUzQblLxpAh0VpHlkmJf6iMPWcR_MYackHmm6iJZTecVVa__iuKlKlVGflO_&opi=122715837",
													},
													StaticResource: &StaticResource{
														CreativeType: "image/png",
														URI:          "\n                                    https://imasdk.googleapis.com/formats/wta/help_outline_white_24dp_with_3px_trbl_padding.png?wp=ca-pub-9939518381636264",
													},
												},
											},
										},
										MediaFiles: &MediaFiles{
											MediaFile: &[]MediaFile{
												{
													ID:                  "GDFP",
													Delivery:            "progressive",
													Width:               640,
													Height:              360,
													Type:                "video/mp4",
													Bitrate:             165,
													Scalable:            true,
													MaintainAspectRatio: true,
													URI:                 "\n                                https://redirector.gvt1.com/videoplayback/id/9c7099ddf54e030c/itag/18/source/dclk_video_ads/requiressl/yes/acao/yes/mime/video%2Fmp4/ctier/L/ip/0.0.0.0/ipbits/0/expire/1719870675/sparams/ip,ipbits,expire,id,itag,source,requiressl,acao,mime,ctier/signature/9BE516AE003D468816584E1DC2FD0FBF2576B83D.210B410A9EDA02E0C8FAECA367A7D49C280E1779/key/ck2/file/file.mp4",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var vast VAST
			err := ReadTestdataXML(tc.file, &vast)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, vast)
		})
	}
}

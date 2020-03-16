# natural-catastrophe-cli
A Golang based CLI too for determining natural catastrophe near you, or a location specified.

Golang Framework:https://github.com/spf13/cobra

Source: https://eonet.sci.gsfc.nasa.gov/docs/v2.1

UI: https://worldview.earthdata.nasa.gov/

URL to FORM:  https://www.google.com/maps/@?api=1&map_action=map&center=-37.29356,-71.95059&zoom=12&basemap=terrain
**NOTE** The coordinates received from the events payload need to be revered in the URL:
https://eonet.sci.gsfc.nasa.gov/api/v2.1/events?api_key=XXXXXXXXXXXXXXXX
```json
{
	"title": "EONET Events",
	"description": "Natural events from EONET.",
	"link": "https://eonet.sci.gsfc.nasa.gov/api/v2.1/events",
	"events": [
		{
			"id": "EONET_4591",
			"title": "Wildfire - Tucapel, Chile",
			"description": "",
			"link": "https://eonet.sci.gsfc.nasa.gov/api/v2.1/events/EONET_4591",
			"categories": [
				{
					"id": 8,
					"title": "Wildfires"
				}
			],
			"sources": [
				{
					"id": "PDC",
					"url": "http://emops.pdc.org/emops/?hazard_id=103651"
				}
			]
    }
  ]
}
```

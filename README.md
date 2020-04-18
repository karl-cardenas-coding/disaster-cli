# natural-catastrophe-cli
A Golang based CLI too for determining natural catastrophe near you, or a location specified.

Golang Framework:https://github.com/spf13/cobra

Go-pretty: https://github.com/jedib0t/go-pretty

Source: https://eonet.sci.gsfc.nasa.gov/docs/v3


UI: https://worldview.earthdata.nasa.gov/

URL to FORM:  https://www.google.com/maps/@?api=1&map_action=map&center=-37.29356,-71.95059&zoom=12&basemap=terrain
**NOTE** The coordinates received from the events payload need to be revered in the URL:
https://eonet.sci.gsfc.nasa.gov/api/v3/events
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


## API Web Service Rate Limits
Limits are placed on the number of API requests you may make using your API key. Rate limits may vary by service, but the defaults are:

Hourly Limit: 1,000 requests per hour
For each API key, these limits are applied across all api.nasa.gov API requests. Exceeding these limits will lead to your API key being temporarily blocked from making further requests. The block will automatically be lifted by waiting an hour. If you need higher rate limits, contact us.

For more info visit https://api.nasa.gov/

**Note**: To generate an API key visit https://api.nasa.gov/

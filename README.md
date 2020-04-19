![Go](https://github.com/karl-cardenas-coding/disaster-cli/workflows/Go/badge.svg?branch=master)
![Go version](https://img.shields.io/github/go-mod/go-version/karl-cardenas-coding/disaster-cli)
# Natural Catastrophe CLI
A Golang based CLI too for determining natural catastrophe near you, or a location specified.

Golang Framework:https://github.com/spf13/cobra

Go-pretty: https://github.com/jedib0t/go-pretty

Source: https://eonet.sci.gsfc.nasa.gov/docs/v3


UI: https://worldview.earthdata.nasa.gov/


## API Web Service Rate Limits
Limits are placed on the number of API requests you may make using your API key. Rate limits may vary by service, but the defaults are:

Hourly Limit: 1,000 requests per hour
For each API key, these limits are applied across all api.nasa.gov API requests. Exceeding these limits will lead to your API key being temporarily blocked from making further requests. The block will automatically be lifted by waiting an hour. If you need higher rate limits, contact us.

For more info visit https://api.nasa.gov/

**Note**: To generate an API key visit https://api.nasa.gov/

### Synopsis

A Golang based CLI too for determining natural catastrophe near you, or a location specified. Visit https://github.com/karl-cardenas-coding/disaster-cli for more information.

```
disaster-cli [flags]
disaster-cli version [flags]
disaster-cli events [flags]

```

### Usage

* [disaster-cli](disaster-cli.md)	 - A CLI too for determining natural catastrophe near you, or a location specified
* [events](disaster-cli_events.md)	 - Returns all events occurring in the world at this point in time.
* [version](disaster-cli_version.md)	 - Print the version number of disaster-cli

### Maintenance
https://github.com/marketplace/actions/github-tag


| Commit message                                                                                                                                                                                   | Release type  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------- |
| `fix(pencil): stop graphite breaking when too much pressure applied`                                                                                                                             | Patch Release |
| `feat(pencil): add 'graphiteWidth' option`                                                                                                                                                       | Minor Release |
| `perf(pencil): remove graphiteWidth option`<br><br>`BREAKING CHANGE: The graphiteWidth option has been removed.`<br>`The default graphite width of 10mm is always used for performance reasons.` | Major Release |

If no commit message contains any information, then **default_bump** will be used.

[![Actions Status](https://github.com/karl-cardenas-coding/disaster-cli/workflows/Go/badge.svg?branch=master)](https://github.com/karl-cardenas-coding/disaster-cli/actions?branch=master)
[![Go version](https://img.shields.io/github/go-mod/go-version/karl-cardenas-coding/disaster-cli)](https://golang.org/dl/)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/karl-cardenas-coding/disaster-cli.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/karl-cardenas-coding/disaster-cli/alerts/)
[![dependa-bot](https://badgen.net/dependabot/karl-cardenas-coding/disaster-cli/247598538?icon=dependabot)](https://badgen.net/dependabot/thepracticaldev/dev.to?icon=dependabot)

<p align="center">
  <img src="/static/img/disaster-gopher.png" alt="drawing" width="400"/>
</p>

A Golang based CLI tool for determining natural catastrophe near you, or a location specified. [Earth Observatory Natural Event Tracker (EONET)](https://eonet.sci.gsfc.nasa.gov/what-is-eonet) is the source for all Data.


## Installation
Disaster-CLI is distributed as a single binary. [Download](https://github.com/karl-cardenas-coding/disaster-cli/releases) the binary and install Disaster-CLI by unzipping it and moving it to a directory included in your system's [PATH](https://superuser.com/questions/284342/what-are-path-and-other-environment-variables-and-how-can-i-set-or-use-them). `~/bin` is the recommended path for UNIX/LINUX environments. 


## Usage

```
disaster [flags]
disaster version [flags]
disaster events [flags]
disaster categories [flags]
disaster update [flags]

```

* [disaster](/documentation/disaster.md)	 - A CLI tool for determining natural catastrophe near you, or a location specified
* [events](/documentation/disaster_events.md)	 - Returns all events occurring in the world at this point in time.
* [categories](/documentation/disaster_categories.md) - Prints all the unique categories of all the events.
* [update](/documentation/disaster_update.md)	 - Updates the local version disaster-cli
* [version](/documentation/disaster_version.md)	 - Print the version number of disaster-cli

## Contributing to Disaster-CLI

For a complete guide to contributing to disaster-cli , see the [Contribution Guide](CONTRIBUTING.md).

Contributions to disaster-cli of any kind including documentation, organization, tutorials, blog posts, bug reports, issues, feature requests, feature implementations, pull requests, answering questions on the forum, helping to manage issues, etc.

## API Web Service Rate Limits
Limits are placed on the number of API requests you may make using your API key. Rate limits may vary by service, but the defaults are:

Hourly Limit: 1,000 requests per hour
For each API key, these limits are applied across all api.nasa.gov API requests. Exceeding these limits will lead to your API key being temporarily blocked from making further requests. The block will automatically be lifted by waiting an hour. If you need higher rate limits, contact us.

For more info visit https://api.nasa.gov/

**Note**: To generate an API key visit https://api.nasa.gov/. Use the `--a` flag to pass in your API key.

## Helpful Links

API Documentation:  https://eonet.sci.gsfc.nasa.gov/docs/v3

Golang Cobra CLI Framework:https://github.com/spf13/cobra

Go-pretty: https://github.com/jedib0t/go-pretty

Source: https://eonet.sci.gsfc.nasa.gov/docs/v3

UI: https://worldview.earthdata.nasa.gov/

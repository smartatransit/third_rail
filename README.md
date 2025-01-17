![smarta](https://user-images.githubusercontent.com/8289478/57379460-f873e280-7174-11e9-9c32-b737bc49650c.png)
<img src="https://user-images.githubusercontent.com/8289478/56633099-d6357d00-662a-11e9-9592-0c58dab8ca55.png" width="300" height="72" />

The Third Rail API is part of the SMARTA project - a collection of tools and services built around
[MARTA APIs](http://www.itsmarta.com/app-developer-resources.aspx) supplemented
with analysis of historic patterns, static schedule data, and external sources like Twitter. 

## Continuous Integration Status

[![Build Status](https://travis-ci.com/smartatransit/third_rail.svg?branch=master)](https://travis-ci.com/smartatransit/third_rail)
[![codecov](https://codecov.io/gh/smartatransit/third_rail/branch/master/graph/badge.svg)](https://codecov.io/gh/smartatransit/third_rail)

## Project Goals

Goals? Oh we've got goals - check 'em out in the
[overview document](https://github.com/smartatransit/infohub/blob/master/vision/overview.md).

### Generating swagger docs

When making changes to API functionality and _especially_ when changing the signature of API calls, you should update the corresponding Swagger comments as well, and then regenerate the Swagger docs.

To do this, first install the Swaggo CLI:

```
go get -u github.com/swaggo/swag/cmd/swag
```

Then run:

```
go generate .
```

### TODO

- [x] Find rail schedule by line
- [x] Find rail schedule by station
- [ ] Find bus schedule by stop
- [ ] Find bus stop by route
- [ ] Find routes by stop
- [x] Find rail stations by location
- [x] Find nearest stations
- [x] Parking status updates
- [ ] Emergency notification updates
- [ ] Add projected arrival/departure time based on historical trends

## Project Maturity

SMARTA is _very_ young. Young, scrappy, and hungry. 😎

## Prerequisites

You will need a [MARTA API key](https://www.itsmarta.com/developer-reg-rtt.aspx)
to fetch the live results from MARTA's base API. For Twitter interactions you will
need a [Twitter Developer account](https://developer.twitter.com/en/apply-for-access) 
and an API client and secret. 

To build and run Third Rail as a container you will need Docker.  

[leiningen]: https://github.com/technomancy/leiningen

## Building

To build the application, run:

    make

## Tests

To run tests for the application, run:

    make test

## License

Copyright© 2020 SMARTA Transit

Distributed under the GNU Public License version 3

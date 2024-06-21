## Overview

A sample repo for learning Go.

`main` contains a the start of a Go version of the caching [prototype](https://github.com/CDCgov/vocabulizer) built for accessing the PHIN VADS FHIR API.

Use the `basic-gin` or `basic-no-gin` branch as a starting point for building out a simple REST API.

## Getting Started

### Prerequisites

Download and install [Go](https://go.dev/)

**Recommended**

<a name="air"></a>[Air](https://github.com/air-verse/air) for live reloading.

### Run

- Download the project file and navigate to the root directory
- `go run main.go` will spin up the project on localhost:8000. If you make any changes, you'll need to restart the server for them to take effect.
- Alternatively, follow the setup instructions for [air](#air) for live reloading

### Endpoints

There are currently four endpoints, a cached and a non-cached version of the following:

**Ping**

Non-cached: `http://localhost:8000/ping`  
Cached: `http://localhost:8000/cache_ping`

Both endpoints will return a JSON response with a timestamp to demonstrate the cached vs. non-cached response.

**Value Set by ID**

Non-cached: `http://localhost:8000/phinvads/ValueSet/:id`  
Cached: `http://localhost:8000/phinvads/cache/ValueSet/:id`

Both will return a single ValueSet resource from PHIN VADS. An ID parameter is required (the bundled `/ValueSet` endpoint has not been built out).

You can add optional query parameters to filter your request.

An invalid request should return a plaintext error.

#### Example Requests

| Request                                    | Response                                                        |
| ------------------------------------------ | ----------------------------------------------------------------- |
| `/phinvads/cache/ValueSet/PH`              | returns the ValueSet named **PHVS_CountrySubdivision_ISO_3166-2** |
| `/phinvads/cache/ValueSet/bananas`         | "ValueSet Name:bananas not found "                                |
| `/phinvads/cache/ValueSet/PH?version=1`    | returns the ValueSet named **PHVS_CountrySubdivision_ISO_3166-2** |
| `/phinvads/cache/ValueSet/PH?version=2`    | "No valueset version #2 found for PH "                            |
| `/phinvads/cache/ValueSet/PH?type=bananas` | returns the ValueSet named **PHVS_CountrySubdivision_ISO_3166-2** |

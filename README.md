## Go module for Open-Meteo Weather API access
[![GoDoc](https://godoc.org/github.com/tdrn-org/go-open-meteo?status.svg)](https://godoc.org/github.com/tdrn-org/go-open-meteo)
[![Build](https://github.com/tdrn-org/go-open-meteo/actions/workflows/build.yml/badge.svg)](https://github.com/tdrn-org/go-open-meteo/actions/workflows/build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=tdrn-org_go-open-meteo&metric=coverage)](https://sonarcloud.io/summary/new_code?id=tdrn-org_go-open-meteo)

### Open-Meteo Weather API
The [Open-Meteo Weather API](https://open-meteo.com/) provides a [REST API](https://open-meteo.com/en/docs) to query weather forecast data. This library provides the Go wrapper for this API.

## License & Attribution
This project is subject to the the MIT License (see the [LICENSE](./LICENSE) file for details).

#### Third-Party Notices
This project uses the official OpenAPI specification from **[Open-Meteo](https://github.com/open-meteo)** as a source for code generation.

* **Component:** Open-Meteo Weather API OpenAPI Specification (`api/forecast.yml`)
* **Copyright:** See https://github.com/open-meteo/open-meteo
* **License:** [GNU Affero General Public License Version 3 (AGPLv3)](https://github.com/open-meteo/open-meteo/blob/main/LICENSE)

The OpenAPI specification file remains under the terms of the AGPLv3. This module's Go code generator itself and any newly generated output code are subject to the MIT License.

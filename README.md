# New Relic Synthetics Runner

### Build and Validate New Relic Synthetics Monitor Scripts on a Local Machine using Docker

---

[New Relic Synthetics](https://newrelic.com/products/synthetics) is a suite of automated, scriptable tools to monitor websites, critical business transactions, and API endpoints. New Relic Synthetics ensures the website is not only available but fully functional.

New Relic Synthetics Scripted Browsers emulate and run complex test cases against a website with real, Selenium-powered Google Chrome browsers to ensure critical processes like checkout and login are always running smoothly.

Depending on the complexity of monitors, building scripts using New Relic Synthetics IDE-style script editor can take longer than expected to productionize. Because, for each validation, New Relic has to provision and run the monitor job in one of their public minion locations.

This repository enable Synthetics managers to build and validate New Relic Synthetics Monitor Scripts on a local machine using Docker and ship faster the well-validated script to production in New Relic.

---

## Prerequisites:
* Docker

## Run:
* Clone or download this repository
* Change current working directory to the cloned location
  * Example `cd newrelic-synthetics-runner`
* As required, edit the `DOCKER_IMAGE_NAME` and `DOCKER_IMAGE_TAG` values in the `go` script file
* Run `./go setup` to pull `newrelic/synthetics-minion-runner` docker image
* As required, edit the appropriate monitor script source file
  * For Simple Browser monitors, edit `src-simple-browser.js`
  * For Scripted Browser monitors, edit `src-script-browser.js`
  * For API tests, edit `src-script-api.js`
* For advanced configurations like Secure credentials, Timeout, Verbose logging, etc., edit the appropriate environment variables file
  * For Simple Browser monitors, edit `env-simple-browser`
  * For Scripted Browser monitors, edit `env-script-browser`
  * For API tests, edit `env-script-api`
* Run `./go <MONITOR-TYPE>` to validate the monitor script
  * For Simple Browser monitors, run `./go simple-browser`
  * For Scripted Browser monitors, run `./go script-browser`
  * For API tests, run `./go script-api`
  * Alternatively, run `./go all` to validate all the monitor types at once
* View Script log, Browser log, Screenshot, and other artifacts in the output directory
  * For Simple Browser monitors, view `output-simple-browser` directory
  * For Scripted Browser monitors, view `output-script-browser` directory
  * For API tests, view `output-script-api` directory
* To delete all the artifacts, run `./go clean`
* Run `./go help` to view the usage of this script

#### For bugs, enhancements, or other requests create an issue in this repository

---

*Note: This repository is purely aimed to augment the velocity of building and validating the New Relic Synthetics Advanced Monitor Scripts under [New Relic License](https://docs.newrelic.com/docs/licenses) and [Acceptable Use Policy](https://docs.newrelic.com/docs/licenses/license-information/general-usage-licenses/acceptable-use-policy)*

---

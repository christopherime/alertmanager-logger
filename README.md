# Alertmanager Logger

[![Build Status](https://github.com/christopherime/alertmanager-logger/workflows/CI/badge.svg)](https://github.com/christopherime/alertmanager-logger/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/christopherime/alertmanager-logger)](https://goreportcard.com/report/github.com/christopherime/alertmanager-logger)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The Alertmanager Logger is a tool that outputs alerts received by Alertmanager to a log file. This can be useful for debugging and troubleshooting alerting rules, as well as for keeping a historical record of alerts.

## Features

* Outputs alerts to a log file in JSON format.
* Supports multiple Alertmanager instances.
* Rotates log files daily.
* Docker image available.

## Installation

### From source

```bash
 go build
```

### From Docker

```bash
docker pull ghcr.io/christopherime/alertmanager-logger:latest docker run -d --name alertmanager-logger -v /path/to/log/directory/:/var/log/amlogger/ -p 9095:9095 ghcr.io/christopherime/alertmanager-logger:latest
```
### Docker Compose

```bash
docker-compose up -d
```

## Configuration

The Alertmanager Logger can be configured using environment variables.

* `ALERTMANAGER_URL`: The URL of the Alertmanager instance.
* `LOG_FILE`: The path to the log file.
* `LOG_LEVEL`: The log level (debug, info, warn, error).
* `ROTATE_LOG_FILES`: Whether to rotate log files daily (true or false).

## Usage

To use the Alertmanager Logger, simply create a new receiver in your Alertmanager configuration and route alerts to it as a webhook.

For example:

```yaml
route:
  routes:
    - receiver: 'alertmanager-logger'
  receivers:
    name: 'alertmanager-logger'
  webhook_configs:
    url: 'http://alertmanager-logger:9095/logger'
```

# alertmanager-logger

## Description

Tools that output in a log file any alert received by AlertManager.

The log file is stored by default in `/var/log/amlogger/alerts.log`

I strongly suggest running this tool in a docker container.

Then bind your directory that you want to `/var/log/amlogger/` in the container.


## Installation

In your alertmanager, create a new receiver and route that will send alerts to this tool as a webhook.

```yaml
route:
    routes:
        receiver: 'alertmanager-logger'
receivers:
  - name: 'alertmanager-logger'
    webhook_configs:
      - url: 'http://alertmanager-logger:9095/logger'
```

### From source

```bash
$ git clone
$ cd alertmanager-logger
$ go build
$ ./amlogger
```

### From docker

Docker run:
```bash
$ docker pull ghcr.io/christopherime/alertmanager-logger:latest
$ docker run -d --name alertmanager-logger -v /path/to/log/directory/:/var/log/amlogger/ -p 9095:9095 ghcr.io/christopherime/alertmanager-logger:latest
```

Docker compose:
```bash
$ docker-compose up -d
```

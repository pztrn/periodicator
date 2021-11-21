# Periodicator

A thing that do some things periodically.

Initially created for periodic issues (or tasks) creation in Gitlab.

## Installing

There are several ways to install Periodicator.

### Using binary release

Head over releases page, grab your binary and configure your system to start binary using cron.

### Using Docker image

Compose a configuration file (read below) and add this to your cron:

```shell
docker run --rm -v ./config.yaml:/periodicator.yaml pztrn/periodicator:latest
```

## Configuring

See config.example.yaml file in repository's root with configuration file structure and parameters description.

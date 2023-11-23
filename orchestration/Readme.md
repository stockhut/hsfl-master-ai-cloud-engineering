# Orchestration

This orchestration tool allows to automatically create multiple docker containers based on a configuration file.

## Features

- automatically pulls images
- set environment variables inside spawned containers
- mount host directories inside containers

## Running

```shell
go run ./cmd/
```
⚠️ If you are running rootless Docker, you won't be able to ping or otherwise connect to your containers. [^rootlessdocker]
[^rootlessdocker]: https://stackoverflow.com/questions/73952959/is-it-possible-to-use-bridge-network-when-running-docker-in-rootless-mode

## Configuration

See `config.yml` for examples.

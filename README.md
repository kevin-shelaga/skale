# Skale - WIP

Skale is a CLI written in go to automatically scale up and down deployments in your cluster

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kevin-shelaga/skale)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevin-shelaga/cronjob-cleaner)](https://goreportcard.com/report/github.com/kevin-shelaga/skale)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/Apache)
[![codecov](https://codecov.io/gh/kevin-shelaga/skale/branch/main/graph/badge.svg?token=SY9DP3GBB8)](https://codecov.io/gh/kevin-shelaga/skale)
![build](https://github.com/kevin-shelaga/skale/workflows/build/badge.svg)

## Install

```sh
go get github.com/kevin-shelaga/skale
```

## Usuage

```sh
Usage:
  skale [command]

Available Commands:
Skale was built to dynamically scale all deployments up and down
in your cluster. This cli can be used as a cost saving measure to force cluster 
auto scaling. For example:

skale up
skale down

  down        Dynamically scale all deployments down
  help        Help about any command
  up          Dynamically scale all deployments up

Flags:
  -A, --all-namespaces string   all namespaces
  -d, --dry-run                 dry run, no changes will be made to the cluster
  -h, --help                    help for skale
  -n, --namespace string        namespace to scale (default "default")
  -v, --verbose                 verbose logging
      --version                 version for skale

Use "skale [command] --help" for more information about a command.
```

## Whats left

### TODO

- [ ] More/better tests
- [ ] Policies around contributions

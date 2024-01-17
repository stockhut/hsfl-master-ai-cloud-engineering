

# Taste Bud Tribe
[![](https://codecov.io/gh/stockhut/hsfl-master-ai-cloud-engineering/graph/badge.svg?token=JD50FBWWAJ)](https://codecov.io/gh/stockhut/hsfl-master-ai-cloud-engineering) 
Note: Depending on how recently a workflow has run, codecov might only display the coverage of that workflow/service (e.g. only coverage for recipe). Click on the bagde to see what coverage is currently shown.

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/stockhut/hsfl-master-ai-cloud-engineering?label=Version)

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/stockhut/hsfl-master-ai-cloud-engineering?filter=frontend-*&label=Version)

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/stockhut/hsfl-master-ai-cloud-engineering?filter=recipe-*&label=Version)
[![Test recipe service](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/recipe-test.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/recipe-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe)

![GitHub tag (with filter)](https://img.shields.io/github/v/tag/stockhut/hsfl-master-ai-cloud-engineering?filter=auth-*&label=Version)
[![Test auth service](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/auth-test.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/auth-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication)

[![Test common package](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-common.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-common.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/common)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/common)

[![Test reverse-proxy](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-reverse-proxy.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-reverse-proxy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy)

[![Test load-balancer](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-load-balancer.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-load-balancer.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer)

A Food-focused social media platform for like-minded people, centered around tags and communities (tribes). Share your favorite recipies, follow your favorite Foodies on the platform and create meal plans to get you through the week.

## Requirements

To build all programs without docker, you will need the following tools installed:

- Go 1.21
- sqlc
- protobuf
- protoc-gen-go
- protoc-gen-go-grpc

To develop and test GitHub actions locally, we recommend [act](https://github.com/nektos/act).

If you want to analyse the profiling data from the Recipe Service, you need
- pprof
- graphviz

## API Client

There is a collection for the API client [Bruno](https://www.usebruno.com/) in the `bruno` directory.

## Versioning

We use [cocogitto](https://docs.cocogitto.io/) and their [GitHub action](https://github.com/cocogitto/cocogitto-action) to automatically create new semantic version releases based on conventional commits.
The [Create new Version](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/release.yml) workflow is executed manually and will create a new commit and new version number for each project prefixed with the name, e.g. `auth-1.2.3`.
After the new commit is tagged and pushed, the action will check whether the `authentication` or `recipe` container images need a rebuild and trigger the respective workflow using the `gh` CLI client (workaround because workflows are not triggered by push events from other workflows [source](https://stackoverflow.com/questions/72110432/github-workflow-is-not-triggered-after-pushing-tags)).
The container Image will be tagged using information from the git tags (see `auth-build-push-image.yml` for details)

## Quickstart

The easiest way to start the app is via Kubernetes/minikube. Please see `deployment` for how to do that!

## Authors
Fabian Wesemann\
fabian.wesemann@stud.hs-flensburg.de\
Hochschule Flensburg

Nele Kirchner\
nele.kirchner@stud.hs-flensburg.de\
Hochschule Flensburg

Alexander Brandt\
alexander.brandt@stud.hs-flensburg.de\
Hochschule Flensburg

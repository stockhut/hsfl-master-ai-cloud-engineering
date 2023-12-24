

# Taste Bud Tribe

[![](https://codecov.io/gh/stockhut/hsfl-master-ai-cloud-engineering/graph/badge.svg?token=JD50FBWWAJ)](https://codecov.io/gh/stockhut/hsfl-master-ai-cloud-engineering)

[![Test recipe service](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-recipe.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-recipe.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe)

[![Test auth service](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-auth.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-auth.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication)

[![Test common package](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-common.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-common.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/common)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/common)

[![Test reverse-proxy](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-reverse-proxy.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-reverse-proxy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/reverse-proxy)

[![Test load-balancer](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-load-balancer.yml/badge.svg)](https://github.com/stockhut/hsfl-master-ai-cloud-engineering/actions/workflows/test-load-balancer.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer)](https://goreportcard.com/report/github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer)

A Food-focused social media platform for like-minded people, centered around tags and communities (tribes). Share your favorite recipies, follow your favorite Foodies on the platform and create meal plans to get you through the week.

## API Client

There is a collection for the API client [Bruno](https://www.usebruno.com/) in the `bruno` directory.

## Versioning

We use [cocogitto](https://docs.cocogitto.io/) and their [GitHub action](https://github.com/cocogitto/cocogitto-action) to automatically create new semantic version releases based on conventional commits.
The `release.yml` workflow is executed for every change on the `main` branch and will create a new commit and new version number for each project prefixed with the name, e.g. `auth-1.2.3`.
After the new commit is tagged and pushed, the action will check whether the `authentication` or `recipe` container images need a rebuild and trigger the respective workflow using the `gh` CLI client (workaround because workflows are not triggered by push events from other workflows [source](https://stackoverflow.com/questions/72110432/github-workflow-is-not-triggered-after-pushing-tags)).
The container Image will be tagged using information from the git tags (see `auth-build-push-image.yml` for details)


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

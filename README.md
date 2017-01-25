# assimilator
WIP. An attempt to port minimum valuable subset of Sentry from Python to the Golang.

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/diyan/assimilator/blob/master/LICENSE)
[![Travis Build](https://travis-ci.org/diyan/assimilator.svg?branch=master)](https://travis-ci.org/diyan/assimilator)
[![Go Report Card](https://goreportcard.com/badge/diyan/assimilator)](http://goreportcard.com/report/diyan/assimilator)

### Setup development enviroment on Arch Linux

```bash
$ sudo pacman -S git vim go go-tools glide
$ vim ~/.zlogin
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
$ mkdir -p $GOPATH/src/github.com/diyan && cd $_
$ git clone git@github.com:diyan/assimilator.git
$ make get-go-tools
$ make get-go-deps
$ make watch-go
```

### Run Sentry in Docker containers
```bash
$ sentry/run_sentry_containers.sh
$ docker exec -ti acme_sentry_web raven test http://763a78a695424ed687cf8b7dc26d3161:763a78a695424ed687cf8b7dc26d3161@localhost:9000/2
```

### Connect to the Sentry's Postgres database
```
$ docker run --rm -ti --name=pgcli \
  -e PGPASSWORD=RucLUS8A \
  --link=acme_sentry_db \
  diyan/pgcli \
  --host=acme_sentry_db \
  --dbname=sentry \
  --user=sentry
```

### TODOs
- Improve error handling, use errors.Wrap
- Helper method to get projectID for current HTTP request that contains orgSlug, projectSlug
- Test framework
- Continuous testing
- DB test fixtures
- Implement Prefix / SetPrefix in Echo/Logrus logger
- Hell of a lot things to do

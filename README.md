# assimilator
WIP. An attempt to port minimum valuable subset of Sentry from Python to the Golang.

### Setup development enviroment on Arch Linux

```bash
$ sudo pacman -S git vim go go-tools glide
$ vim ~/.zlogin
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
$ mkdir ~/go/assimilator && cd $_
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

### TODOs
- Test framework
- Continuous testing
- DB test fixtures
- Implement Prefix / SetPrefix in Echo/Logrus logger
- Hell of a lot things to do
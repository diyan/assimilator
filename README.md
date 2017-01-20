# assimilator
WIP. An attempt to port minimum valuable subset of Sentry from Python to the Golang.

### Development enviroment setup on Arch Linux

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

### TODOs
- Test framework
- Continuous testing
- DB test fixtures
- Implement Prefix / SetPrefix in Echo/Logrus logger
- Hell of a lot things to do
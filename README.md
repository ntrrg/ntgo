**QA:**
[![goreport](https://goreportcard.com/badge/github.com/ntrrg/ntgo)](https://goreportcard.com/report/github.com/ntrrg/ntgo) 
[![BCH compliance](https://bettercodehub.com/edge/badge/ntrrg/ntgo?branch=master)](https://bettercodehub.com/results/ntrrg/ntgo)
[![GoDoc](https://godoc.org/github.com/ntrrg/ntgo?status.svg)](https://godoc.org/github.com/ntrrg/ntgo)

**CI:**
[![Travis build btatus](https://travis-ci.com/ntrrg/ntgo.svg?branch=master)](https://travis-ci.com/ntrrg/ntgo)
[![Gitlab build btatus](https://gitlab.com/ntrrg/ntgo/badges/master/pipeline.svg)](https://gitlab.com/ntrrg/ntgo/commits/master)
[![Drone build status](https://ci.nt.web.ve/api/badges/ntrrg/ntgo/status.svg)](https://ci.nt.web.ve/ntrrg/ntgo)

**Docker:**
[![Docker Build Status](https://img.shields.io/docker/build/ntrrg/ntgo.svg)](https://store.docker.com/community/images/ntrrg/ntgo/)
[![MicroBadger Size](https://img.shields.io/microbadger/image-size/ntrrg/ntgo.svg)](https://microbadger.com/images/ntrrg/ntgo)

**Packages:**

`math`:

&nbsp;&nbsp;`arithmetic`:
![API status](https://img.shields.io/badge/status-stable-brightgreen.svg)
[![gocover](http://gocover.io/_badge/github.com/ntrrg/math/arithmetic)](https://gocover.io/github.com/ntrrg/ntgo/math/arithmetic)
[![GoDoc](https://godoc.org/github.com/ntrrg/ntgo/math/arithmetic?status.svg)](https://godoc.org/github.com/ntrrg/ntgo/math/arithmetic)

`net`:

&nbsp;&nbsp;`http`:
![API status](https://img.shields.io/badge/status-unstable-red.svg)
[![gocover](http://gocover.io/_badge/github.com/ntrrg/net/http)](https://gocover.io/github.com/ntrrg/ntgo/net/http)
[![GoDoc](https://godoc.org/github.com/ntrrg/ntgo/net/http?status.svg)](https://godoc.org/github.com/ntrrg/ntgo/net/http)

## Install

**All the packages:**

```shell-session
$ go get -u github.com/ntrrg/ntgo/...
```

**Specific package:**

```shell-session
$ go get -u github.com/ntrrg/ntgo/math/arithmetic
```

## Features

See <https://godoc.org/github.com/ntrrg/ntgo>.

## Uninstall

**All the packages:**

```shell-session
$ go clean -i github.com/ntrrg/ntgo/...
```

```shell-session
$ rm -rf $GOPATH/github.com/ntrrg/ntgo
```

```shell-session
$ rm -rf $GOPATH/pkg/$(go env GOHOSTOS)_$(go env GOHOSTARCH)/github.com/ntrrg/ntgo
```

**Specific package:**

```shell-session
$ go clean -i github.com/ntrrg/ntgo/math/arithmetic
```

```shell-session
$ # Si solo se había instalado el paquete a eliminar
$ rm -rf $GOPATH/github.com/ntrrg/ntgo
```

## Acknowledgment

Working on this project I use/used:

* [Debian](https://www.debian.org/)

* [XFCE](https://xfce.org/)

* [st](https://st.suckless.org/)

* [Zsh](http://www.zsh.org/)

* [GNU Screen](https://www.gnu.org/software/screen)

* [Git](https://git-scm.com/)

* [EditorConfig](http://editorconfig.org/)

* [Vim](https://www.vim.org/)

* [GNU make](https://www.gnu.org/software/make/)

* [Chrome](https://www.google.com/chrome/browser/desktop/index.html)

* [Gogs](https://gogs.io/)

* [Github](https://github.com)

* [Gitlab](https://gitlab.com/)

* [Docker](https://docker.com)

* [Drone](https://drone.io/)

* [Travis CI](https://travis-ci.org)

* [Go Report Card](https://goreportcard.com)

* [GoCover](http://gocover.io)

* [Better Code Hub](https://bettercodehub.com)


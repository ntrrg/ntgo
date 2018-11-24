[![Travis build btatus](https://travis-ci.com/ntrrg/ntgo.svg?branch=master)](https://travis-ci.com/ntrrg/ntgo)
[![codecov](https://codecov.io/gh/ntrrg/ntgo/branch/master/graph/badge.svg)](https://codecov.io/gh/ntrrg/ntgo)
[![goreport](https://goreportcard.com/badge/github.com/ntrrg/ntgo)](https://goreportcard.com/report/github.com/ntrrg/ntgo) 

**ntgo** is a collection of packages that provides HTTP utilities, data
structures implementations and tools commonly used by other projects.

| Package | Status |
|:-|:-|
| `container` | |
|`container/arithmetic` | ![API status](https://img.shields.io/badge/status-stable-brightgreen.svg) [![GoDoc](https://godoc.org/github.com/ntrrg/ntgo/container/arithmetic?status.svg)](https://godoc.org/github.com/ntrrg/ntgo/container/arithmetic) |
| `net` | |
| `net/http` | ![API status](https://img.shields.io/badge/status-unstable-red.svg) [![GoDoc](https://godoc.org/github.com/ntrrg/ntgo/net/http?status.svg)](https://godoc.org/github.com/ntrrg/ntgo/net/http) |
| `net/http/middleware` | ![API status](https://img.shields.io/badge/status-testing-yellow.svg) [![GoDoc](https://godoc.org/github.com/ntrrg/ntgo/net/http/middleware?status.svg)](https://godoc.org/github.com/ntrrg/ntgo/net/http/middleware) |

---

**Warning:** since this project is personal and experimental, it doesn't
provide any kind of guarantee or backward compatibility. My recommendation is
to vendor it or just copy the needed pieces of code.

---

## Install

```shell-session
$ go get -u github.com/ntrrg/ntgo/...
```

**Specific package:**

```shell-session
$ go get -u github.com/ntrrg/ntgo/container/arithmetic
```

## Uninstall

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
$ go clean -i github.com/ntrrg/ntgo/container/arithmetic
```

## Contributing

See the [contribution guide](CONTRIBUTING.md) for more information.

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


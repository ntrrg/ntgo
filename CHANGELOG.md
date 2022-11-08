# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

[Unreleased]: https://github.com/ntrrg/ntgo/compare/v0.8.0...master
## [Unreleased][]

### Added

* `io/fs`: New package for file system operations
* `os/bin`: New package with UNIX like utilities
* `errors`: New package for error handling

[0.8.0]: https://github.com/ntrrg/ntgo/compare/v0.7.0...v0.8.0
## [0.8.0][]

### Added

* `net/http`: `ListenAndServeUDS` and `ShutdownServerOn` functions
* `os`: `GetenvOr` function

### Removed

* `net/http`: Remove `Server` since it provides an extra layer of abstraction
  without enough advantages
* `bytes`: Remove package

[0.7.0]: https://github.com/ntrrg/ntgo/compare/v0.6.0...v0.7.0
## [0.7.0][]

### Added

* `net/http/middleware`: `ResponseWriteAdapter` interface
* `net/http/middleware`: `AdaptResponseWriter` and `IsAdaptedResponseWriter`
  functions
* `bytes`: Bytes slice pool implementation (`Pool`)
* `runtime/memrep`: New package for retrieving low-level memory presentations
* `runtime`: `IsBigEndian` and `IsLittleEndian` functions
* `bytes`: `BufferPool.AddWait` and `BufferPool.GetWait` methods
* `os`: `Copy`, `CopyFile` and `CopyDir` functions

### Changed

* `net/http`: Use a single `ListenAndServe` method for `Server`
* Improve project structure

[0.6.0]: https://github.com/ntrrg/ntgo/compare/v0.5.0...v0.6.0
## [0.6.0][]

### Changed

* `reflect/arithmetic`: Renamed `GetVal` to `Val`
* Improve project structure

### Fixed

* `reflect/arithmetic`: Panics with zero operanders

### Removed

* `os`: `Cp` helper

[0.5.0]: https://github.com/ntrrg/ntgo/compare/v0.4.1...v0.5.0
## [0.5.0][]

### Changed

* Migrate module domain
* Improve project documentation and structure

[0.4.1]: https://github.com/ntrrg/ntgo/compare/v0.4.0...v0.4.1
## [0.4.1][]

### Changed

* Improve project documentation

[0.4.0]: https://github.com/ntrrg/ntgo/compare/v0.3.1...v0.4.0
## [0.4.0][]

### Added

* `generics/arithmetic`: `Ne` function

### Changed

* `generics/arithmetic`: Rename to `reflect/arithmetic`
* `generics/arithmetic`: Improve package structure

### Fixed

* `bytes`: BufferPool overflowing

[0.3.1]: https://github.com/ntrrg/ntgo/compare/v0.3.0...v0.3.1
## [0.3.1][]

### Fixed

* `bytes`: Bad buffer initialization ([#4](https://github.com/ntrrg/ntgo/issues/4))

[0.3.0]: https://github.com/ntrrg/ntgo/compare/v0.2.1...v0.3.0
## [0.3.0][]

### Added

* `bytes`: Buffer pool implementation (`BufferPool`)

[0.2.1]: https://github.com/ntrrg/ntgo/compare/v0.2.0...v0.2.1
## [0.2.1][]

### Fixed

* `os`: Can't overwrite subdirectories ([#2](https://github.com/ntrrg/ntgo/issues/2))

[0.2.0]: https://github.com/ntrrg/ntgo/compare/v0.1.0...v0.2.0
## [0.2.0][]

### Added

* `os`: `Cp` helper

### Changed

* `container/arithmetic`: Renamed to `generics/arithmetic`


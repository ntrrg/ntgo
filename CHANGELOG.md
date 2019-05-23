# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

[Unreleased]: https://github.com/ntrrg/ntgo/compare/v0.3.1...master
## [Unreleased][]

### Add

* `generics/arithmetic`: Add Ne function

### Changed

* `generics/arithmetic`: Renamed to `reflect/arithmetic`
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

* `os`: Cp helper

### Changed

* `container/arithmetic`: Renamed to `generics/arithmetic`


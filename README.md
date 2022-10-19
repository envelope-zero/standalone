# Envelope Zero Standalone version

[![Release](https://img.shields.io/github/release/envelope-zero/standalone.svg?style=flat-square)](https://github.com/envelope-zero/standalone/releases/latest) [![Go Reference](https://pkg.go.dev/badge/github.com/envelope-zero/standalone.svg)](https://pkg.go.dev/github.com/envelope-zero/standalone) [![Go Report Card](https://goreportcard.com/badge/github.com/envelope-zero/standalone)](https://goreportcard.com/report/github.com/envelope-zero/standalone)

## Usage

### Quick start

Download the latest release and start the executable. This will do two things:

- A terminal will open, showing the log output you can ignore this
- A tray item “Envelope Zero“ will be created with options to open Envelope Zero (which will open it in your browser) and to quit the app.

If the tray does not work for some reason, you can always open http://localhost:3200 directly in your browser.

## Backing up of data

Envelope Zero will create an `envelope-zero` directory in the standard application data directory for your operating system. This is the following:

- Windows: `%APPDATA%/envelope-zero`
- macOS: `~/Library/Application Support/envelope-zero`
- Other Unix based systems: `~/.local/share/envelope-zero`

If you back up this directory, all Envelope Zero data is backed up.

## Supported Versions

This software is constantly developed, therefore only the latest version is supported. If you encounter an issue, please update to the latest version and verify that it still exists in that version.

Please check the [releases page](https://github.com/envelope-zero/standalone/releases) for the latest release.

## Versioning

This project uses [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html). This means that the version has three parts: `major.minor.patch`.
Releases are fully automated and happen on every _feature_ and _bug fix_ that is merged into the `main` branch.

The versions increase as follows:

- The **major** version increases when there are breaking changes, meaning that the behavior of the software changes compared to an earlier version (unless that behavior was wrong, then it's a **patch** version increase). Please check the [upgrading documentation](docs/upgrading.md) before updating major versions!
- The **minor** version increases when there are new features
- The **patch** version increases when bugs are fixed.
- If the Envelope Zero backend or frontend are updated, the version bump is the same as for the dependency that is updated
- If a release with only dependency updates is made, it bumps the `PATCH` version.

Whenever a version increases, all numbers to the right of it are reset to 0.

The following things are looked at for versioning (called “public API” in Semantic Versioning):

- Location of the data on your computer
- Behavior of the application, e.g. how budget values are calculated, including new features

## Contributing

Please see [the contribution guidelines](CONTRIBUTING.md).

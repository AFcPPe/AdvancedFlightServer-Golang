# AdvancedFlightServer-Go

[中文版](Readme-CN.md)

## Introduction

AdvancedFlightServer-Go is a flight simulation multiplayer server written in Go. It supports Swift, Echo, Euroscope or other custom clients.

## Features

- Faster speed
- Uses FSD Version 3.000 Draft 9 protocol
- Outputs data in JSON format
- Stores accounts in PostgreSQL database
- Supports high concurrency
- Cross-platform support (developed and tested on Windows 11 and CentOS 8, theoretically supports all platforms supported by Golang)
- Resolves the issue of random crashes when running FSD V3.000 Draft 9 on Linux systems
- Built-in METAR data.

## Usage Guide

1. Compile the code using `go build AdvancedFlightServer/main`.
2. Run the compiled executable.
3. The program will generate a `config.ini` file on the first run. Close the program.
4. Fill in the entries in the `config.ini` file.
5. Reopen the program.

## Open Source License

This software is licensed under the AGPL v3.0.

## Contribution

Contributions to this codebase are welcome via pull requests.

## Commercial Version

In addition to the open-source version, we also offer a commercial version with the following features:

- Uses Vatsim v3.4 protocol
- Supports modification of flight plans by controllers
- Supports Euroscope built-in ATIS
- Supports INFO line and Logoff time for controllers
- Supports VisualPosition refresh (refreshes position information every 0.2 seconds when the aircraft is on the ground)
- Supports locking of flight plans (automatically locks modified flight plans until the controller goes offline)
- Supports operation record of labels, and more

Please note that the commercial version may have different details and licensing requirements. If you are interested in the commercial version, please contact us for more information.
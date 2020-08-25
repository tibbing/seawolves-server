# Seawolves-server

## Description

This is the serverless backend of a turn-based strategy game, inspired by [Herrscher der Meere ("Ruler of the Seas")](https://www.mobygames.com/game/herrscher-der-meere) published by attic Entertainment Software GmbH in 1997.

## Prerequisites

This repository is intended to be opened using VS Code Dev Containers.
This means that all the tooling you need should be available already,
as long as you have Docker installed.

To get started, install [the extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
if you need to, and then VS Code should prompt you to open the workspace
in remote container instead. If not, choose
"Remote Containers: Open Folder in Container..." from the command palette.

## Usage  

### Build

To build all lambda functions:  
`sh build.sh`

To build a specific lambda function:  
`sh build.sh [function name]`

### Test

`go test lib/... -v && go test app/... -v`

### Deploy

This application is built for AWS Lambda. Running the build script will output one zip file per lambda function, ready to deploy. A CI/CD pipeline will be used in the future to automate this process.

## License

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

See [LICENSE](LICENSE) for full details.

# Assetto Corsa Competizione Server Web Interface

[![Discord Shield](https://discordapp.com/api/guilds/913752018588422174/widget.png?style=shield)](https://discord.gg/AVWdF56t6c)
[![CircleCI](https://circleci.com/gh/assetto-corsa-web/accweb.svg?style=svg)](https://circleci.com/gh/assetto-corsa-web/accweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/assetto-corsa-web/accweb)](https://goreportcard.com/report/github.com/assetto-corsa-web/accweb)

The successor of [acweb](https://github.com/assetto-corsa-web/acweb)! accweb lets you manage your Assetto Corsa Competizione servers via a nice and simple web interface. You can start, stop and configure server instances and monitor their status.

## Table of contents

1. [Features](#features)
2. [Changelog](#changelog)
3. [Installation](#installation)
4. ~~[Docker](#docker)~~
5. [Backup](#backup)
6. [Contribute and support](#support)
7. [Build release](#release)
8. [Links](#links)
9. [License](#license)
10. [ACCWeb Discord Server](#discord)

## Features
<a name="features" />

* create and manage as many server instances as you like
* configure your instances in browser
* start/stop instances and monitor their status
* view server logs
* copy server configurations
* import/export server configuration files
* delete server configurations
* three different permissions: admin, mod and read only (using three different passwords)
* status page for non logged in users
* easy setup
    * no database required
    * simple configuration using environment variables
    
## Changelog
<a name="changelog" />

See [CHANGELOG.md](CHANGELOG.md).

## Installation and configuration
<a name="installation" />

accweb is installed by extracting the zip on your server, modifing the YAML configuration file to your needs and starting it in a terminal.

### Manuall installation

1. download the latest release from the release section on GitHub
2. extract the zip file on your server
3. edit the `config.yml` to match your needs
4. open a terminal
5. change directory to the accweb installation location
6. start accweb using `./accweb` on Linux and `accweb.exe` on Windows
8. leave the terminal open (or start in background using screens on Linux for example)
9. visit the server IP/domain and port you've configured, for example: http://example.com:8080

I recommend to setup an SSL certificate, but that's out of scope for this instructions. You can enable a certificate inside the `config.yml`.

**Note that you have to install [wine](https://www.winehq.org/) if you're on Linux.**

## Docker
<a name="docker" />

**Docker is not supported at the moment, you can try using the image, but I highly recommend waiting for v2 to be finished (which will have a new image).**

Keep in mind that this docker image doesn't include accserver.exe file, you will have to copy it manually to the /accserver directory inside the container after it started. Check the volumes section for details.

### Docker Image

To build the image manually:

* git clone https://github.com/assetto-corsa-web/accweb
* `cd accweb`
* `docker build --pull --rm -f "Dockerfile" -t accweb:latest "."`

An official image is availabe at: https://hub.docker.com/r/kugel/accweb. It can be pulled using: `docker pull kugel/accweb` and started by running `docker run -it accweb`.

By default this image exposes port 8080 in order to have the web UI working. If you want to run the game server you will need to open the ports manually. For example, if I want to run two server instances I need to open two differents port:

```
docker run -it accweb -p 8080:8080 -p 2600:2600 -p 2601:2601
```

This will allow you to run a server on 2600 and 2601 port.

### Docker Compose File

The docker-compose file at the root of the git repository can be used to deploy accweb. It comes with pre-defined env variables and volumes. **Keep in mind that this compose file is not secured. Please follow the documentation and check the recommandations.**

### Volumes

Here is the list of docker volumes and their purpose:
* `accweb`: Contains the accweb project
* `accserver`: Must contain the `accserver.exe`. It has to be put there manually as we can't bundle it inside docker
* `sslcerts`: This volume is dedicated to SSL certificates

### Environment Variables List

| Variable name | Description | Default value |
|---|---|---|
| ACCWEB_HOST | ACC web server host URL | 0.0.0.0:8080 (not secure) |
| ACCWEB_ENABLE_TLS | Enable or disable SSL | false (not secure) |
| ACCWEB_CERT_FILE | Certificate file location | /sslcerts/certificate.crt |
| ACCWEB_PRIV_FILE | Certificate key location | /sslcerts/private.key |
| ACCWEB_ADMIN_PASSWORD | Admin password | weakadminpassword |
| ACCWEB_MOD_PASSWORD | Moderator password | weakmodpassword |
| ACCWEB_RO_PASSWORD | Read only password | weakropassword |
| ACCWEB_LOGLEVEL | App log level | info |
| ACCWEB_CORS | Default CORS origin | "*" (accept all sources) |

## Backup
<a name="backup" />

To backup your files, copy and save the `config` directory as well as the `config.yml`. The `config` directory can later be placed inside the new accweb version directory and you can adjust the new `config.yml` based on your old configuration (don't overwrite it, there meight be breaking changes).

## Contribute and support
<a name="support" />

If you like to contribute, have questions or suggestions you can open tickets and pull requests on GitHub.

All Go code must have been run through go fmt. The frontend and backend changes must be (manually) tested on your system. If you have issues running it locally open a ticket.

To run the accweb locally is really simple, make sure that the attribute `dev` is set to true in your `config.yml` file.

### Frontend development environment

Our current frontend was built using Vue.js and can be found inside `public` directory.

To run the watcher use the following command.

```shell
make run-dev-frontend
```
Then when you edit any js file, the watcher will detect and rebuild the js package.

### Backend development environment

ACCweb backend is running over golang and can be found inside `internal` directory.

Use the following command to run the backend on your terminal.

```shell
make run-dev-backend
```
Keep in mind that you need to restart the command for see the changes that you made in the code working (or not :zany_face:) 

### Visual Studio Code - Remote container

There is a pre-built development environment setup for ACCWeb for Visual Studio Code and Remote Containers. Please, check here how to setup and use: https://code.visualstudio.com/docs/remote/containers

## Build release
<a name="release" />

To build a release, execute the `build_release.sh` script (on Linux) or follow the steps inside the script. You need to pass the build version as the first parameter. Example:
To build a release, execute the `build_release.sh` script (on Linux) or follow the steps inside the script. You need to pass the build version as the first parameter. Example:

```shell
./build/build_release.sh 1.2.3
```

This will create a directory `releases/accweb_1.2.3` containing the release build of accweb. This directory can be zipped, uploaded to GitHub and deployed on a server.

## Links
<a name="links" />

* [Docker Hub](https://cloud.docker.com/repository/docker/kugel/accweb/general)
* [Assetto Corsa Forums](https://www.assettocorsa.net/forum/index.php?threads/release-accweb-assetto-corsa-competizione-server-management-tool-via-web-interface.57572/)

## License
<a name="license" />

MIT

## ACCWeb Discord Server
<a name="discord" />

Join us on our Discord server to get and provide support.

[![ACCWeb Discord](https://discordapp.com/api/guilds/913752018588422174/widget.png?style=banner4)](https://discord.gg/AVWdF56t6c)

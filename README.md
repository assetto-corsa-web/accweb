# Assetto Corsa Competizione Server Web Interface

[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/assetto-corsa-web/accweb) 
[![CircleCI](https://circleci.com/gh/assetto-corsa-web/accweb.svg?style=svg)](https://circleci.com/gh/assetto-corsa-web/accweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/assetto-corsa-web/accweb)](https://goreportcard.com/report/github.com/assetto-corsa-web/accweb)

The successor of [acweb](https://github.com/assetto-corsa-web/acweb)! accweb lets you manage your Assetto Corsa Competizione servers via a nice and simple web interface. You can start, stop and configure server instances and monitor their status.

## Table of contents

1. [Features](#features)
2. [Changelog](#changelog)
3. [Installation](#installation)
4. [Backup](#backup)
5. [Contribute and support](#support)
6. [Build release](#release)
7. [Links](#links)
8. [License](#license)
9. [Screenshots](#screenshots)

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

### Version 1.9.1

* better default configuration
* fixed exporting configuration if server name contains invalid characters for filenames

### Version 1.9.0

* added automatic generation of private/public token files
* switched to yaml configuration instead of environment variables
* new design

### Version 1.8.0

* minor changes to the global view
* corrections of values for "formationLapType"
* added parameter "simracerWeatherConditions" in event.json
* added parameter "isFixedConditionQualification" in event.json
* added "bop.json"
* added "assistRules"

IMPORTANT: You will have to delete the servers already created in order to create new ones!

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

accweb should generate the key files for authentication on its own, but in case that doesn't work you can do it manually.
To generate the RSA key pair, you can use the `gen_rsa_keys.sh` on Linux or install one of the tools available for Windows and run `gen_rsa_keys.cmd`.
You can also use an online service which generates RSA key pairs (search for "generate rsa key pair online").

**Note that you have to install [wine](https://www.winehq.org/) if you're on Linux.**

## Backup
<a name="backup" />

To backup your files, copy and save the `config` directory as well as the `config.yml`. The `config` directory can later be placed inside the new accweb version directory and you can adjust the new `config.yml` based on your old configuration (don't overwrite it, there meight be breaking changes).

## Contribute and support
<a name="support" />

If you like to contribute, have questions or suggestions you can open tickets and pull requests on GitHub.

All Go code must have been run through go fmt. The frontend and backend changes must be (manually) tested on your system. If you have issues running it locally open a ticket. You can use the `dev.sh` and `gen_rsa_keys.sh` scripts to start accweb on your computer (on Linux).

## Build release
<a name="release" />

To build a release, execute the `build_release.sh` script (on Linux) or follow the steps inside the script. You need to pass the build version as the first parameter. Example:

```
./build_release.sh 1.2.3
```

This will create a directory `accweb_1.2.3` containing the release build of accweb. This directory can be zipped, uploaded to GitHub and deployed on a server.

## Links
<a name="links" />

* [Docker Hub](https://cloud.docker.com/repository/docker/kugel/accweb/general)
* [Assetto Corsa Forums](https://www.assettocorsa.net/forum/index.php?threads/release-accweb-assetto-corsa-competizione-server-management-tool-via-web-interface.57572/)

## License
<a name="license" />

MIT

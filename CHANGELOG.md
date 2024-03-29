# Changelog

## 1.25.0
* html adjustments on server password [#273](https://github.com/assetto-corsa-web/accweb/pull/273)
* New track nurburgring_24h
* Remove circleci integration

## 1.24.0
* GT2 pack update [#269](https://github.com/assetto-corsa-web/accweb/pull/269)

## 1.23.0
* Feature: token timeout in configuration file [#262](https://github.com/assetto-corsa-web/accweb/issues/262)
* Fix setTyreSetCount default value to 50. [#227](https://github.com/assetto-corsa-web/accweb/issues/227)
* Fix LFM files encoding [#265](https://github.com/assetto-corsa-web/accweb/issues/265) [#266](https://github.com/assetto-corsa-web/accweb/issues/266)

## 1.22.2
* bump some go libs
* first test with wine killing processes issue

## 1.22.1
* Full support for non numeric instances id

## 1.22.0
* Fix arm64 node dependencies
* Fix live view for non numeric instance ids
* Add official docker support

## 1.21.0
* Add Morocoo nationality #245
* Update go to 1.21.1
* Update go dependencies
* Fix Authorization issues
* Fix css from read only server list #247
* Fix README #228

## 1.20.1
* Add event metadata field #223
* Hide /admin on chat command #233

## 1.20.0
* Changes for DLC 1.19 #238
* Update golang dep versions
* Update node dep versions

## 1.19.0
* Fix some typo errors. #210 #212
* Fix float point handling for rain and cloud. #211
* Introducing the advance windows features as Firewall management, custom core affinity and cpu priority. #213
* Add accserver auto update before start instance. 
 
## 1.18.0
* Add sorting by number of players.
* Add session remaining time in the servers list.
* Move server name to the beggining of the create/update server settings.
* Add live sessions chat to the view.
* Add config `skip_wine` to skip wine usage even on linux (good for development on windows).
* Update js libraries versions.
* [Add devcontainer to development environment (vscode)](https://code.visualstudio.com/docs/remote/containers).
* Add America DLC Tracks

## 1.17.0
* Link logo to front page.
* Add title attribute to "log out"-button. 
* Group small input fields together to avoid long lists of wide input fields.
* Reorder fields to group similar functionality together and group them visually.
* Add PublicIP field to the configuration json file.
* Add new Challengers Pack DLC content
* Add nationality field to the drivers entry list.


## 1.16.2
* Fix server password handler.
* Fix session timeout, redirecting to login screen.
* Add confirmation before stop all acc servers.
* Fix live laps delta calculation.
* Add live gaps column during races.
* Add live current splits.
* Fix live driver handshake.
* Removing the necessity of external scripts to generate secret keys.

## 1.16.1
* Fix acc dedicated server create instance. ( #185 )
* Building releases assets automatically.
* Generating Swagger docummention of accweb api.

## 1.16.0
* Adding instance live view.
* Fix bug on UI to prevent unnecessary backend requests.
* Add utility buttons to server configuration for event session to add Q/R and P/Q/R automagically.

## 1.15.1
* Sorting option on server list screen.
* Fix release mode message when it's in production mode.
* Update js libraries versions

## 1.15.0
* Reorganization of the codebase
* added server instance auto start ( #74 and #119 )
* added stop all instances
* embedding web client public files
* simplified file import ( #84 )

## 1.14.2
* Update js libraries versions
* Fix build script to force linux OS build

## 1.14.1
* fixed import servers with old encoding charset

## 1.14.0
* fixed server names with character "/"
* fixed OpenSSL download url
* ACC v1.8 update

## 1.13.4

* fixed ballast -> ballastKg in bop.json
* updated dependencies

## 1.13.3

* added trackTemp support
* fixed openssl download link on Windows

## 1.13.2

* changed default for trackMedalsRequirement from -1 to 0
* added ignorePrematureDisconnects with a default of 1 (use 0 on Linux)
* updated dependencies

## 1.13.1

* added missing tracks for 2020

## 1.13.0

* added tracks from the British GT DLC

## 1.12.3

* added missing tracks
* updated dependencies

## 1.12.2

* add track and cars for 2020 DLC to BOP

## 1.12.1

* fixed tracks in event configuration

## 1.12.0

* added 2020 cars and tracks
* fixed error starting accweb due to incompatible clib on Linux

## 1.11.2

* converted bunch of number fields to check boxes, where applicable
* fixed broken selection fields for entrylist
* fixed defaultGridPosition being in the wrong location
* fixed some inconsistent code formatting

## 1.11.1

* added selection dropdown view for DriverCategory and ForcedCarModel

## 1.11.0

**When upgrading from a previous version, make sure you replace all true/false values with 1/0 for simracerWeatherConditions and isFixedConditionQualification in event JSON.**

* fixed type of simracerWeatherConditions and isFixedConditionQualification
* added tyreSetCount to eventRules.json

## 1.10.0

* added car groups for GT4 pack DLC

## 1.9.2

* fixed saving/loading assist rules
* fixed assist rules import
* fixed name of ballast to ballastKg in entrylist

## 1.9.1

* better default configuration
* fixed exporting configuration if server name contains invalid characters for filenames

## 1.9.0

* added automatic generation of private/public token files
* switched to yaml configuration instead of environment variables
* new design

## 1.8.0

* minor changes to the global view
* corrections of values for "formationLapType"
* added parameter "simracerWeatherConditions" in event.json
* added parameter "isFixedConditionQualification" in event.json
* added "bop.json"
* added "assistRules"

IMPORTANT: You will have to delete the servers already created in order to create new ones!

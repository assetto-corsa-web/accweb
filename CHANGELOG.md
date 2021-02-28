# Changelog

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

# Project Flogo Implementation for CatalystML

## Overview of Specification
CatalystML is a language agnostic specification aimed at facilitating data transformations.  The initiating use-case being transforming incoming data into a form for a 1-to-1 input into a Machine Learning (ML) model.  Refer to the [CatalystML specification](https://github.com/TIBCOSoftware/catalystml) for specifics around the specification, supported operations, etc.

## Flogo/Golang implementation
This repository containeds an implementation of CatalystML for [Flogo](flogo.io) written in [Golang](https://golang.org/).  With relatively minor context switching it can also be used as a Golang implementation.  The documents included within this repository detail the choices that were made when creating this implementation as well as any deviations from the specification itself (primarily things yet to be implemented). 

Within Flogo this specification is implementated as an action.  A flogo action is an engine within Flogo that runs a specific type of event manipulation function (like a flow, stream, rules engine, etc.).  A more detailed discussion is included in the documentation.

## Use of this implementation

As discussed above this implementation is written with Golang within the Flogo ecosystem.  As such CatalystML can be used with the flogo command line interface (with a flogo.json) or the Golang Flogo API (library).  Two examples for each of the CLI or the API are discussed below.

### Flogo Command Line Interface

Within a flogo.json CatalystML

#### Flogo Action

Running the CatalystML action alone on top of a trigger.

#### Flogo Activity

Running the CatalystML action as an activity (read function) within another action such as flows or streams.

### Golang Flogo API (library)

#### Within Flogo app (with Trigger)

#### Raw Golang code (without Trigger)



# Project Flogo Implementation for CatalystML

## Overview of Specification
CatalystML is a language agnostic specification aimed at facilitating data transformations.  The initiating use-case being transforming incoming data into a form for a 1-to-1 input into a Machine Learning (ML) model.  Refer to the [CatalystML specification](https://github.com/TIBCOSoftware/catalystml) for specifics around the specification, supported operations, etc.

## Flogo/Golang implementation
This repository containeds an implementation of CatalystML for [Flogo](flogo.io) written in [Golang](https://golang.org/).  With relatively minor context switching it can also be used as a Golang implementation.  The documents included within this repository detail the choices that were made when creating this implementation as well as any deviations from the specification itself (primarily things yet to be implemented). 

Within Flogo this specification is implementated as an action.  A flogo action is an engine within Flogo that runs a specific type of event manipulation function (like a flow, stream, rules engine, etc.).  A more detailed discussion is included in the documentation.

## Use of this implementation

As discussed above this implementation is written with Golang within the Flogo ecosystem.  As such CatalystML can be used with the flogo command line interface (with a flogo.json) or the Golang Flogo API (library).  Two examples for each of the CLI or the API are discussed below.

### Flogo Command Line Interface

[Flogo's Command line interface](https://tibcosoftware.github.io/flogo/flogo-cli/flogo-cli/) is built around a json object that represents the structure of a Flogo application.  Compiling this json ([here is an example of compiling a flogo.json](https://github.com/project-flogo/ml/blob/master/examples/flowsOutlier/README.md)) with the Flogo CLI then creates an executable binary.  Within a flogo.json CatalystML

There are multiple ways to embed a CatalystML structure within flogo:
1) As a flogo action that responds to a trigger.  In this case a trigger responds to input data, while the CatalystML action transforms that data.  An example flogo.json of CatalystML as a flogo action is located [here](https://github.com/project-flogo/catalystml-flogo/tree/master/examples/flogoAction).
2) As a Flogo activity within a flogo flow or stream.  Flows and streams are flogo actions that allow you to chain predefined functions called activities.  In this case CatalystML is simply one step in a chain of functions.  This is how CatalystML-flogo would be used to interact with a machine learning model with the model executed within another [activity](https://github.com/project-flogo/ml/tree/master/activity/inference).  An example flogo.json of CatalystML as a flogo activity within a stream is located [here](https://github.com/project-flogo/catalystml-flogo/tree/master/examples/flogoActivity).

### Golang Flogo API (library)

Project-Flogo allows for the functinality of flogo to be integrated with custom Golang code by using the Flogo Golang API as a Golang library.  This can be done by either following a template that includes triggers and flows/streams or by just calling the CML action as a function.

1) An example of Golang code that includes CatalystML within a Flogo template that includes triggers and flows is located [here](https://github.com/project-flogo/catalystml-flogo/tree/master/examples/golangWithTrigger).

2) [Here](https://github.com/project-flogo/catalystml-flogo/tree/master/examples/golangWOTrigger) is an example of using CatalystML in Golang using the CatalystML action as a function.



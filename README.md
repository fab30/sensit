Sensit
=================

`sensit` command allows to lauch a small server that receives callbacks from the sensit server and stores the measures in an [InfluxDB](https://influxdb.com/) timeseries database.

## Building

### Prerequesites
To build the command you must have installed :
* [Go](https://golang.org/) 1.5+
* [`godep`](https://github.com/tools/godep)



### Checkout the code
You can checkout the code with a `git clone` or with a

    go get github.com/joelvim/sensit

### Build the executable
To build the executable, enter

    cd $GOPATH/github.com/joelvim/sensit
    godep go build

You should get the `sensit` command in the folder.

If you want to install the sensit command in your path, run :

    cd $GOPATH/github.com/joelvim/sensit
    godep go install

This command will compile the executable and copy it in the `$GOPATH/bin` directory. Just add this directory in your `$PATH` env variable, and play with it.

## Usage

The command is documented, just type `sensit -h` to get the help.

## HTTP API

The server provides 2 endpoints

    /ping #healthcheck your app
    /api/v1/temperature #receive the metrics.

## What's missing

* __SSL Support__ : the application does not support SSL but is targeted to be behind a proxy like nginx or Apache HTTP.
* __Support for other metrics__ than temperature
* __Support for periodic callback types__ : agregation is made by the TSDB, we don't need preagregated data.

## Coming soon (or later)

* Command to import the history by extracting it from the [Sensit API](https://api.sensit.io).

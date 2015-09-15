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
    go build

You should get the `sensit` command in the folder.

If you want to install the sensit command in your path, run :

    cd $GOPATH/github.com/joelvim/sensit
    go install

This command will compile the executable and copy it in the `$GOPATH/bin` directory. Just add this directory in your `$PATH` env variable, and play with it.

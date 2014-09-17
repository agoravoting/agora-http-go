# Install

## Install dependencies

Configure your go paths. To do so, first install mercurial and git. This can
be done for example in opensuse this way (in other distros will vary):

    sudo zypper install git-core mercurial

Then configure your paths and install godep:

    echo 'export GOPATH=$HOME/go' >> ~/.bashrc
    echo 'export PATH=$GOPATH/bin:$PATH' >> ~/.bashrc
    echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
    source ~/.bashrc
    go get github.com/tools/godep

Once that's done, you can install dependencies, compile and install agora-http-go:

    cd path/to/agora-http-go
    godep restore
    godep go install

## Execute unit tests

In order to execute unit tests you must set up a test database for the demo api to work against.

The configuration of the database is setup in two places. One used for
goose, a go database migration tool placed in `agora-http-go/db/dbconf.yml`, and the
other for the general configuration of agora-http-go, which can be placed anywhere,
but there's a provided example `agora-http-go/config.json`.

Here we will use the default configuration for the database for convenience,
but you should obviously do not use it (especially for the passwords) in a
production environment.

Prerequisites: install postgresql database server in your system. Then, create
the database (typically, with the postgres system user):

    su - postgres
    createuser -P agora_http_go
    createdb -O agora_http_go agora_http_go

You must also have goose installed, if it is not use

    godep go install bitbucket.org/liamstask/goose/cmd/goose

Then create the tables using the goose migration system:

    cd path/to/agora-http-go
    goose up

Once you have set up the database, you can run tests with

    ./demoapi/test.sh

or

    demoapi\test.cmd
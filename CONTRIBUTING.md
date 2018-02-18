# Contributing

By participating to this project, you agree to abide our [code of
conduct](/CODE_OF_CONDUCT.md).

## Setup your machine

`fortress` is written in [Go](https://golang.org/).

Prerequisites are:

* Build:
  * `make`
  * [Go 1.9+](http://golang.org/doc/install)

Clone `fortress` from source into `$GOPATH`:

```sh
$ git clone git@github.com:leandro-lugaresi/fortress.git $GOPATH/src/github.com/leandro-lugaresi/fortress
$ cd $GOPATH/src/github.com/leandro-lugaresi/fortress
```

If you created a fork clone your fork and after add the my repository as upstream remote:

```sh
$ git clone git@github.com:{your-name}/fortress.git $GOPATH/src/github.com/leandro-lugaresi/fortress
$ cd $GOPATH/src/github.com/leandro-lugaresi/fortress
$ git remote add upstream git@github.com:leandro-lugaresi/fortress.git
```

### Install

Install the build and lint dependencies:

``` sh
$ make setup
```

A good way of making sure everything is all right is running the test suite:

``` sh
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

``` sh
$ go build
```

When you are satisfied with the changes, we suggest you run:

``` sh
$ make ci
```

Which runs all the linters and tests.

## Submit a pull request

Push your branch to your `example` fork and open a pull request against the
master branch.

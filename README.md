# Jrnl

Jrnl is a Go based journal for the command line. It's inspired by the fantastic python tool of the same name; [jrnl](http://jrnl.sh/). It stores the entries in a database using BoltDB in the user home directory, and has no external dependencies. 

Simply put the binary on your path and away you go!

## Usage

``` bash
$ jrnl
jrnl is a CLI journal manager.

Usage:
  jrnl [command]

Available Commands:
  add         Add a new entry to your journal
  remove      Remove an entry from your journal
  list        List all entries in your journal

Use "jrnl [command] --help" for more information about a command.
```

## Credit

* Jrnl is inspired by the fantastic python tool of the same name [jrnl](http://jrnl.sh/).
* Jon Calhoun follwing his [exericse](https://gophercises.com/exercises/) building CLI's

## TODO

* [ ] implement date tags to `add` command
* [ ] improve domain model to include time
* [ ] test coverage

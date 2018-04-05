# Jrnl

Jrnl is a Go based journal for the command line. It's inspired by the fantastic python tool of the same name; [jrnl](http://jrnl.sh/). It stores the entries in a database using BoltDB in the user home directory, and has no external dependencies.

Simply put the binary on your path and away you go!

## Installing

Using Jrnl is easy. First, use go get to install the latest version of the application. This command will install the executable along with the library and its dependencies:

`$ go get -u github.com/willis7/jrnl`

## Usage

```
$ jrnl
jrnl is a CLI journal manager

Usage:
  jrnl [command]

Available Commands:
  add         Adds an entry in your journal,
  export      Export a formatted journal
  help        Help about any command
  list        Lists all of your entries.
  remove      Removes a journal entry.

Flags:
  -h, --help   help for jrnl

Use "jrnl [command] --help" for more information about a command.
```

## Examples

`list` returns the full list of entries
```
$ jrnl list
Here's your jrnl entries:
1. 2018-04-01 Easter
```

`add today` is a shorthand way of adding an entry for today. `yesterday` works too.
```
$ jrnl add today
Enter text: Easter Day
created; 1. 2018-04-01 Easter Day
```

`add <date>` takes a date layout as yyyy-mm-dd.
```
$ jrnl add 2018-04-01
Enter text: Easter Day
created; 1. 2018-04-01 Easter Day
```

## Credit

* Jrnl is inspired by the fantastic python tool of the same name [jrnl](http://jrnl.sh/).
* Jon Calhoun follwing his [exericse](https://gophercises.com/exercises/) building CLI's

## TODO

* [x] implement date keywords to `add` command (such as `today`, `yesterday`).
* [x] improve domain model to include time
* [ ] test coverage

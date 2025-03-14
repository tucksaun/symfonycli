# Contributing

This guide is meant to help you start contributing to the Symfony CLI by
providing some key hints and explaining specifics related to this project.

## Go choice

First-time contributors could be surprised by the fact that this project is
written in Go whereas it is highly related to the Symfony Framework which is
written in PHP.

Golang has been picked because it is well suited for system development and
has close-to-zero runtime dependencies which make releasing quite easy. This
is ideal for a tool that might be used on a wide range of platform and
potentially on systems where the requirements to run Symfony are not meant.
Golang is also usually quite easy to apprehend for PHP developers having some
similiarities in the mindset.

## Setup Go

Contributing to the CLI, implies that one must first setup Go locally on their
machine. Instructions are available on the [https://golang.org/dl](Go website).
Just pick the latest version available: Go will automatically download the
version currently in use in the project and dependencies as required.

## Local setup

First fork this repository and clone it in some place of your liking.

> **FIXME:** is it required to create the `go.work` ? 

Next, try to build and run the project:

````bash
$ go build .
```

If any error happen you must fix them before going on. If no error happen, this
should produce a binary in the project directory. By default, this binary is
named after the project directory named and suffixed by `.exe` on Windows.

You should be able to run it right away:

````bash
$ ./symfony-cli version
```

This binary is self-contained: you can copy it as-is to another system and
execute it without any installation process.

> *Tip:* This binary can be executed from anywhere by using it's absolute path.
> This is pretty handy during development when you need to run it in a project
> directory and you don't want to overwrite your system-installed Symfony CLI.

Before and after changing code you should ensure tests are passing:

```bash
$ go test ./...
```

## Coding style

The CLI follows the Go standard code formatting. To fix the code formatting one
can use the following command:

```bash
$ go fmt ./...
```

One can also uses the `go vet` command in order to fix common mistakes:

```bash
$ go vet ./...
```

## Cross compilation

By its purpose, the CLI has to be multiplatform which means that as some point
you might need to compile the code for another platform.

One can easily do it thanks to Go crossplatform compiling capabilities. For
example the following command will compile the CLI for Windows:
```bash
$ GOOS=windows go build .
```

`GOOS` and `GOARCH` environment variables are used to target another and/or
another platform.

During development, please take into consideration (particularily in the
processes management sections) that we currently support the following
platforms/architectures combinations:
- Linux / 386
- Linux / amd64
- Linux / arm
- Linux / arm64
- Darwin / amd64
- Darwin / arm64
- Windows / 386
- Windows / amd64

## Cod generation

Part of the code is generated automatically. One should not need to regenerate
the code themselves as an action in the CI is responsible for it. In the
eventuality one would need to debug it code generation can be run as follows:

```bash
$ go generate ./...
```

If you add a new code generation command, please also update the GitHub
workflow in `.github/workflows/go_generate_update.yml`.

## Additional repositories

Contrary to the Symfony PHP Framework which is a mono repository, the CLI
tool is developped in multiple repositories. `symfony-cli/symfony-cli` is the
main repository where lies most of the logic and is the only project producing
a binary.

Every other repositories are mostly independant and it is highly probable that
you don't need to have a look at them. Hoewever, in the eventuality where you
would have to, here is the description of each repositories scope:
- `symfony-cli/symfony-cli` is the main repository where lies most of the logic
  lies
- `symfony-cli/phpstore` is a independant library in charge of the PHP
  installations discovery and the logic to match a specific version to a given
  version constraint.
- `symfony-cli/console` is a independant library created to ease the process
  of Go command-line application. This library has been created with the goal
  of mimicing the look and feel of the Symfony Console for the end-user.
- `symfony-cli/terminal` is a wrapper around the Input and Ouput in a command
  line context. It provides helpers around styling (output formatters and
  styling - à la Symfony) and interactivity (spinners and questions helpers)
- `symfony-cli/dumper` is a library similar to Symfony VarDumper component
  providing a `Dump` function usefull to introspect variables content and
  particularily useful in the strictly typed context of Go.

If you ever have to work on those package, you can setup your local working
copy of the CLI to work with a local copy by creating a `go.work` file at the
`symfony-cli` root with the following content:

````
# go.work

use .

use ../path-to/cert

use ../path-to/console

use ../path-to/terminal

use ../path-to/dumper

use ../path-to/phpstore
```
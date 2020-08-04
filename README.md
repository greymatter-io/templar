# Templar

[![CircleCI](https://circleci.com/gh/greymatter-io/templar.svg?style=svg&circle-token=7c00fc19d2b4b5df2cc804381fd95059c4666950)](https://circleci.com/gh/greymatter-io/templar)
[![Maintainability](https://api.codeclimate.com/v1/badges/581e8ca47befad079798/maintainability)](https://codeclimate.com/github/greymatter-io/templar/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/581e8ca47befad079798/test_coverage)](https://codeclimate.com/github/greymatter-io/templar/test_coverage)

A command line utility for rendering Go template files.

## Overview

Templar allows you to quickly render one or more Go template files to disk injecting environment variables or command line arguments.

## Installation

At present, templar only supports Linux and macOS due to the fact that we allow file system permission manipulation in an operating system specific way.

### Binaries

Release binaries are available for templar from the release page [here](https://github.com/greymatter-io/templar/releases).  Download the appropriate binary for your operating system and either invoke it explicitly (e.g., `./templar.linux`) or copy it to your `PATH`.

Note, it's likely that you'll want to rename it `templar` for ease of use and the following instructions make this assumption.

### Docker

Docker images of templar are available as well and can be found in [DockerHub](https://hub.docker.com/repository/docker/greymatterio/templar).

## Usage

Templar uses [Cobra](https://github.com/spf13/cobra) so all commands support a help option (e.g., `templar -h` or `templar --help`). As a result, the following usage instructions only cover the top level use cases and do not attempt to provide descriptions of all options available.

### Version

To print the current version and commit for the templar binary run the following command:

    templar version 

### Render

Rendering a template to disk is performed by running the following syntax:

    templar render TEMPLATE_PATH:RENDERED_PATH[:RENDERED_MODE]

Where:

- `TEMPLATE_PATH` is the path of the template file
- `RENDERED_PATH` is the path of to which the template file is rendered
- `RENDERED_MODE` is the Unix permission in numeric notation (i.e., 0700) of the rendered file

For example, the following command will render the template at `/etc/example/example.json.tpl` to `/etc/example/example.json` and ensure that it only can only be read by the owner.

    templar render /etc/example/example.json.tpl:/etc/example/example.json:0400

Also note that it is also possible to render multiple templates in a single command:

    templar render /etc/example/example.json.tpl:/etc/example/example.json /etc/example/example.yaml.tpl:/etc/example/example.yaml

#### User Variables

Templar supports injecting variables into rendered templates with command line arguments using the `-v` or `-variable` option. The following command sets a variable named `AUTHOR` to `Joan Wilder`.

    templar render -v "AUTHOR=Joan Wilder" /var/lib/rts.json.tpl:/var/lib/rts.json

Using a command line variable within the template is done with the `variable` function exposed to the template:

    {{ variable "AUTHOR" }}

#### Environment Variables

Templar supports injecting environment variables into rendered templates with command line arguments using the `-e` or `--environment` option.

    templar render -e AUTHOR /var/lib/rts.json.tpl:/var/lib/rts.json

Using an environment variable within the template is done with the `variable` function exposed to the template:

    {{ variable "AUTHOR" }}

Note that only environment variables explicitly passed via the `-e` or `--environment` option will be availble to prevent accidentally leaking sensitive material.

#### Available Functions

Templar also supports a number of additional functions that can be used to inject and manipulate common values in templates.

| Function   | Example            | Description                                                                                     |
| ---------- | ------------------ | ----------------------------------------------------------------------------------------------- |
| `variable` | `{{ var "NAME" }}` | Returns the value of a command line or environment variable if set; otherwise, an empty string. |

Note that all other [functions](https://golang.org/pkg/text/template/#hdr-Functions) supported within Go templates are also available.

## Building

### Dependencies

In order to build templar the following dependencies are required.

- Go (1.13.7)
- Make (3.81 or greater)

### Build

To build a binary for the current operating system run the following command:

    make build

Additionally, build targets are provided for macOS, Linux and Docker (i.e., `make build.docker`, `make build.linux`, and `make build.macos`).

### Testing

To test templar run the following command:

    make test

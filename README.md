## Description

This is a Heroku buildpack for Go web apps that have their entire Go workspace
under version control.

Only the `src` directory needs to be under version control in the Go workspace.
An additional `.go` file needs to be created to specify certain information
that the buildpack needs. The main package needs to be on the first line of
this file, while the Go version to use to compile the app may be on the
second line:

    <main-package>
    [go-version]

## Build Process

The buildpack will use the latest version if the second line in the `.go` file
is omitted.  Currently, that is 1.2. It will download Go to the cache of the
Heroku app and use it to compile the main package into a binary, which Heroku
will then execute as a web process.

Specifying the buildpack can be done in the following ways:

    # During Heroku app creation
    heroku create --buildpack https://github.com/perriv/go-workspace-heroku-buildpack.git
    # Modifying an existing Heroku app
    heroku config:set BUILDPACK_URL=https://github.com/perriv/go-workspace-heroku-buildpack.git

## Additional Information

An example repository structure can be found in `webapp`. Please notice these
three things:

1. The contents of the `.go` file.
2. The relative path of the template.
3. The `PORT` environment variable.

From the root of the repository, this should successfully compile and run the
web app:

    GOPATH="$(pwd)" go install "$(head -n 1 .go)"
    PORT=8080 "./bin/$(head -n 1 .go)"

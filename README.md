# Palworld API Wrapper
Easy to use wrapper for the Palworld REST API


[Set up a Palworld Server](https://tech.palworldgame.com/)

[Read the API Docs](https://tech.palworldgame.com/category/rest-api)

## Installation

You can install the package by using `go get` in your Go Project.

```
go get -u github.com/Ju0x/palworldapi
```
This will pull the latest release.

## Usage

First you need to initialize a new PalworldAPI instance with credentials.

```go
pal := palworldapi.New("http://localhost:8212", "admin", "admin password")
```

Typically the username is `admin`.

The `AdminPassword` is set in the Server Configuration. ([How to configure the Palworld Server](https://tech.palworldgame.com/settings-and-operation/configuration))


## Examples

You can find some examples [here](https://github.com/Ju0x/palworldapi/tree/main/examples)
